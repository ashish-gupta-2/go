package main

import (
	"fmt"
	// "example.com/greetings"
)

var messages = []string{"Test", "the", "producer", "consumer", "problem", "in", "golang", "today", "is", "awesome", "day"}

// put only channel
func producer(link chan<- string) {
	for _, msg := range messages {
		fmt.Println("Producing msg::", msg)
		link <- msg
	}
	close(link)
}

// get only channel
func consumer(link <-chan string, done chan<- bool) {
	for msg := range link {
		fmt.Println("Consuming msg ::", msg)
	}
	done <- true
}

func main() {
	// Get a greeting message and print it.
	// message := greetings.Hello("Gladys")
	// fmt.Println(message)

	link := make(chan string)
	done := make(chan bool)

	go producer(link)
	go consumer(link, done)

	<-done

}
