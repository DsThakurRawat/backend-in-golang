// context_with_timeout.go
// ---------------------------------------------------------
// 2. WithTimeout: Automatic cancellation after a duration
// ---------------------------------------------------------
// Used everywhere in network requests, database queries, and 
// external API calls. It cancels automatically after a set time.
//
// Scenario: We simulate a database query that takes 600ms.
// However, our context has a strict timeout of 400ms.
// The context will automatically cancel the DB query midway.

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// Context will automatically cancel after 400ms
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	
	// Still defer! Frees resources if work finishes BEFORE the timeout happens.
	defer cancel()

	fmt.Println("Main: Starting DB query (Timeout set to 400ms)...")
	
	// DB takes 600ms, but timeout is 400ms. We expect an error.
	err := simulateSlowDatabase(ctx, 600*time.Millisecond) 

	if err != nil {
		// context.DeadlineExceeded is the specific error returned when a timeout occurs
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Main: Error - DB query took too long (Timeout reached)")
		} else {
			fmt.Printf("Main: Error - %v\n", err)
		}
	} else {
		fmt.Println("Main: Query succeeded!")
	}
}

// simulateSlowDatabase acts like a database query that listens for context cancellation.
func simulateSlowDatabase(ctx context.Context, duration time.Duration) error {
	// We use select to wait on either the work finishing OR the context canceling
	select {
	case <-time.After(duration):
		// This simulates the work taking exactly 'duration' to complete
		fmt.Println("DB: Query finished successfully")
		return nil
	case <-ctx.Done():
		// This triggers if the context timeout is reached before the work finishes
		fmt.Println("DB: Query aborted midway because context was canceled")
		return ctx.Err()
	}
}
