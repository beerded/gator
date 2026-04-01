package main

import (
	"context"
	"fmt"
	"database/sql"
	"time"

	"github.com/beerded/gator/internal/database"
)

// const testUrl = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: agg <time_between_reqs>\n")
	}
	timeStr := cmd.args[0]
	duration, err := time.ParseDuration(timeStr)
	if err != nil {
		return fmt.Errorf("Error setting parse duration: %w\n", err)
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

	params := database.MarkFeedFetchedParams{
		ID: 			nextFeed.ID,
		LastFetchedAt:	sql.NullTime{Time: time.Now(), Valid: true,},
	}
	feed, err := s.db.MarkFeedFetched(ctx, params)
	if err != nil {
		return fmt.Errorf("Error Marking feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* '%s'\n", item.Title)
	}
	return nil
}
