package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/beerded/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed requires name and url as arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("Unable to get info for user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
		Url: 		url,
		UserID:		user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Unable to create feed: %w", err)
	}
	// Follow the feed
	followParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("Unable to follow own feed: %w", err)
	}

	fmt.Printf("Feed:\n%+v\n", feed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	data, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error listing feeds: %w", err)
	}

	fmt.Println("FEEDS:")
	for _, row := range data {
		printFeed(row.Name, row.Url, row.Addedby)
	}
	return nil
}

func printFeed(name, url, addedBy string) {
	fmt.Printf("Feed Name:			%s\n", name)
	fmt.Printf("Feed URL:			%s\n", url)
	fmt.Printf("Added By:			%s\n", addedBy)
	fmt.Println("===========================================")
}

