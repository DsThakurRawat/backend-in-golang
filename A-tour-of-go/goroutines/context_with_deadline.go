// context_with_deadline.go
// ---------------------------------------------------------
// 3. WithDeadline: Automatic cancellation at a specific time
// ---------------------------------------------------------
// Exactly the same behavior as WithTimeout, but takes an 
// absolute time.Time instead of a time.Duration.
//
// Scenario: We want an operation to abort at exactly a specific
// wall-clock time in the future.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// We want the context to cancel at exactly [current time + 300ms]
	deadline := time.Now().Add(300 * time.Millisecond)
	
	// Create the context with the absolute deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	
	// Always defer cancel
	defer cancel()

	fmt.Println("Main: Waiting for deadline (300ms from now)...")

	// Block until the context is done.
	// ctx.Done() closes exactly when the deadline is reached.
	<-ctx.Done()

	// ctx.Err() will show 'context deadline exceeded'
	fmt.Println("Main: Deadline reached! Context error:", ctx.Err())
}
