// context_with_cancel.go
// ---------------------------------------------------------
// 1. WithCancel: Manual Cancellation
// ---------------------------------------------------------
// Use this when you want to explicitly say "stop working"
// to all downstream goroutines.
//
// Scenario: A worker goroutine is doing some background task.
// When the main goroutine decides it's time to shut down, it
// calls `cancel()`. This sends a signal to `ctx.Done()`, which
// the worker is listening for.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context.Background() is the root context.
	// WithCancel returns a copy of the parent context with a new Done channel.
	ctx, cancel := context.WithCancel(context.Background())
	
	// ALWAYS defer the cancel function.
	// This ensures that all resources associated with the context are freed
	// when this function exits, even if we never explicitly call cancel().
	defer cancel()

	fmt.Println("Main: Starting worker goroutine...")

	// Start a worker goroutine, passing the context to it.
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				// c.Done() returns a channel that's closed when cancel() is called.
				// This case executes when the cancellation signal is received.
				fmt.Println("Worker: Received cancel signal. Shutting down gracefully.")
				return // ALWAYS return to prevent goroutine leaks
			default:
				// If c.Done() is not ready, do the actual work.
				fmt.Println("Worker: Doing work...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	// Main function waits a bit, letting the worker do its thing
	time.Sleep(500 * time.Millisecond)
	
	// Now, manually trigger the cancellation
	fmt.Println("Main: Calling cancel() now...")
	cancel()
	
	// Wait a moment to let the worker print its shutdown message before main exits
	time.Sleep(100 * time.Millisecond) 
	fmt.Println("Main: Exiting.")
}
