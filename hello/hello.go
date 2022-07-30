package main

import (
	"fmt"
	"log"

	greetings "example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("zamonia")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
