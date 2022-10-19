package main

import (
	"fmt"

	"github.com/mayalaat/go-exercises/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Alejandro")
	fmt.Println(message)
}
