/*
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
*/

package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User Not Fount: %v", e.Username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	e1 := &UserNotFound{Username: "mike"}
	e2 := &UserNotFound{Username: "mike"}
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
		if err == e1 {

		}

		if err == e2 {

		}
	}
}
