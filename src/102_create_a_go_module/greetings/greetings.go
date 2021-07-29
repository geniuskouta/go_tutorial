package greetings

import "fmt"

/* the first capital letter func is an exported name */
func Hello(name string) string {
	/*
	* initialize message variable in one line instead of the following:
	* var message string
	* message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}