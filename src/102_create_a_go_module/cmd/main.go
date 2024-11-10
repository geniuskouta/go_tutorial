package main

import (
	"fmt"
	"log"
	"tutorial_102/pkg/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // print the command name ("greetings: ")
	log.SetFlags(0)              //  at the start of its log messages, without a time stamp or source file information.

	names := []string{"Kouta", "Gill", "Kat"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err) // print the error and stop the program
	}

	fmt.Println(messages)
}
