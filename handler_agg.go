package main

import (
	"context"
	"fmt"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/beerded/gator/internal/database"
)

// const testUrl = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <time_between_reqs>", cmd.name)
	}
	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid parse duration: %w\n", err)
	}
	fmt.Printf("Collecting feeds every %s\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Unable to fetch next feed: %w", err)
	}

	feed, err := s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID:				nextFeed.ID,
		LastFetchedAt: 	sql.NullTime{Time: time.Now(), Valid: true,},
	})
	if err != nil {
		return fmt.Errorf("Error Marking feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:	t,
				Valid: 	true,
			}
		} else {
			fmt.Printf("Couldn't parse timestamp %v so just publishing as NULL\n", item.PubDate)
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now(),
			UpdatedAt:		time.Now(),
			Title:			item.Title,
			Url:			item.Link,
			Description:	item.Description,
			PublishedAt:	publishedAt,
			FeedID:			feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			return fmt.Errorf("Error saving post: %w", err)
		}
	}

	fmt.Printf("Collected %v posts from feed '%v'\n", len(rssFeed.Channel.Item), feed.Name)
	return nil
}
