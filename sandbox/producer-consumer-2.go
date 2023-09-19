package main

import (
	"fmt"
)

var messages = []string{"Test", "the", "producer", "consumer", "problem", "in", "golang", "today", "is", "awesome", "day"}
var consumers = 3;
// put only channel
func producer(link chan<- string) {
	for _, msg := range messages {
		fmt.Println("Producing msg::", msg)
		link <- msg
	}
	close(link)
}

// get only channel
func consumer(worker int,link <-chan string, done chan<- bool) {
	for msg := range link {
		fmt.Printf("Message %v is consumed by worker %v.\n", msg, worker)
	}
	done <- true
}

func main() {
	 
	link := make(chan string)
	done := make(chan bool)

	go producer(link)
	for i :=0 ; i< consumers; i++{
		go consumer(i, link, done)
	}

	<-done

}
