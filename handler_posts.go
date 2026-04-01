package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beerded/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) > 0 {
		i, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Could not convert limit to an integer: %w", err)
		}
		limit = int32(i)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:		user.ID,
		Limit:		limit,
	})
	if err != nil {
		return fmt.Errorf("Error Getting posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: 			%s\n", post.Title)
		fmt.Printf("URL:			%s\n", post.Url)
		fmt.Println("====================================")
	}
	return nil
}
