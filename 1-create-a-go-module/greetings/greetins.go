package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())

	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome",
		"Great to see you, %v!",
		"Hello %v. How u doing?",
	}
	return formats[rand.Intn(len(formats))]
}

// Aprendi a usar MAP
