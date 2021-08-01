package greetings

import (
	"errors"
	"fmt"
)

/* the first capital letter func is an exported name */
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name") //any Go function can return multiple values
	}


	/*
	* initialize message variable in one line instead of the following:
	* var message string
	* message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil //nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded
}