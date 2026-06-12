/*
Errors

Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}

(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
*/
package interfaces

import (
	"fmt"
	"time"
)

type MyError_11 struct {
	When time.Time
	What string
}

func (e *MyError_11) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run_11() error {
	return &MyError_11{
		time.Now(),
		"it didn't work",
	}
}

func I11() {
	if err := run_11(); err != nil {
		fmt.Println(err)
	}
}
