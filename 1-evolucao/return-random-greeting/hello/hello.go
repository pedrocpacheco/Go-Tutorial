package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// ! Configura as mensagens no terminal
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Pedro")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
