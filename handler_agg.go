package main

import (
	"context"
	"fmt"
)

const testUrl = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	fmt.Printf("Feed:\n%+v\n", rssFeed)

	return nil
}
