package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kouta")
	fmt.Println(message)
}