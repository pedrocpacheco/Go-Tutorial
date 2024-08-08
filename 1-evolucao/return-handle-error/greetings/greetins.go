package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name
	message := fmt.Sprintf("Hello %s", name)
	return message, nil
}

// ! Chamando errors
// Preciso importar a biblioteca "errors"
// Usei o errors.New para criar um novo error
