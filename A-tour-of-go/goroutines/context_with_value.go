// context_with_value.go
// ---------------------------------------------------------
// 4. WithValue: Passing request-scoped data
// ---------------------------------------------------------
// Used to pass metadata through the call stack (e.g., Trace IDs, 
// Auth Tokens, User IDs) without altering function signatures.
//
// Scenario: An HTTP middleware extracts a User ID from a token and 
// puts it into the context. A deeply nested database or service
// function extracts that User ID to log or process data.

package main

import (
	"context"
	"fmt"
)

// Define a custom type for context keys. 
// This is a CRITICAL best practice to prevent key collisions between 
// different packages. If everyone used strings, two packages might both
// use the key "userID" and overwrite each other's data.
type contextKey string

const userIDKey contextKey = "userID"

func main() {
	// 1. Inject data into the context (usually done in middleware)
	// context.WithValue creates a new context containing the key-value pair.
	ctx := context.WithValue(context.Background(), userIDKey, "user_998877")

	fmt.Println("Main: Injected User ID into context.")
	
	// Pass the context down the call stack
	processRequest(ctx)
}

// processRequest simulates a downstream function (like a service handler)
// It only takes the context as a parameter, not the specific user data.
func processRequest(ctx context.Context) {
	fmt.Println("Service: Extracting data from context...")
	
	// 2. Extract the data deep in the call stack
	val := ctx.Value(userIDKey)

	// Type assert safely since ctx.Value returns 'any' (interface{})
	if userID, ok := val.(string); ok {
		fmt.Printf("Service: Processing data for User ID: %s\n", userID)
	} else {
		fmt.Println("Service: No user ID found in context, or wrong type")
	}
}
