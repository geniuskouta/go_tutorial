package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // print the command name ("greetings: ")
    log.SetFlags(0) //  at the start of its log messages, without a time stamp or source file information.
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err) // print the error and stop the program
	}

	fmt.Println(message)
}