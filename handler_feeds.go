package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/beerded/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed requires name and url as arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
		Url: 		url,
		UserID:		user.ID,
	})
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

	fmt.Println("Feed:")
	fmt.Printf("	ID: 			%v\n", feed.ID)
	fmt.Printf("	CreatedAt:		%v\n", feed.CreatedAt)
	fmt.Printf("	UpdatedAt:		%v\n", feed.UpdatedAt)
	fmt.Printf("	Name:			%v\n", feed.Name)
	fmt.Printf("	URL:			%v\n", feed.Url)
	fmt.Printf("	UserID:			%v\n", feed.UserID)
	fmt.Println("=====================================================")
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
	fmt.Printf("	Feed Name:			%v\n", name)
	fmt.Printf("	Feed URL:			%v\n", url)
	fmt.Printf("	Added By:			%v\n", addedBy)
	fmt.Println("====================================================")
}

