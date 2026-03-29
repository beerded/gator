package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/beerded/gator/internal/database"
)

func handlerCreateFeedFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("URL is required to follow feed")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error looking up feed: %w", err)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID: 	user.ID,
		FeedID:		feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error Creating Feed Follow: %w", err)
	}

	fmt.Println("Created Followed Feed:")
	printFeedFollow(feedFollow)
	return nil
}

func handlerGetFeedFollowsForUser(s *state, cmd command) error {
	user := s.cfg.CurrentUserName

	feedFollowsForUser, err := s.db.GetFeedFollowsForUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Error Getting Feed Follows for user '%s': %w", user, err)
	}
	fmt.Printf("Feed follows for '%s':\n", user)
	for _, feedFollow := range feedFollowsForUser {
		fmt.Printf("	%s\n", feedFollow.FeedName)
	}
	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("	ID:		%v\n", feedFollow.ID)
	fmt.Printf("	Created At:	%v\n", feedFollow.CreatedAt)
	fmt.Printf("	Updated At:	%v\n", feedFollow.UpdatedAt)
	fmt.Printf("	UserID:		%v\n", feedFollow.UserID)
	fmt.Printf("	User Name:	%v\n", feedFollow.UserName)
	fmt.Printf("	FeedID:		%v\n", feedFollow.FeedID)
	fmt.Printf("	Feed Name:	%v\n", feedFollow.FeedName)
	fmt.Printf("==============================================\n")
}
