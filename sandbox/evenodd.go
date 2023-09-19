package main

import (
	"fmt"
	"time"
)

func main(){

	evenChan := make(chan int)
	oddChan := make(chan int, 1)
	done := make(chan bool)

	oddChan <- 1

	go printOdd(oddChan , evenChan)
	go printEven(oddChan, evenChan , done)


	<- done
	close(evenChan)
	close(oddChan)
}


func printOdd(oddChan chan int , evenChan chan int){
	for odd := range oddChan {
		fmt.Println("odd :", odd)
		even := odd+1
		evenChan <- even
	}
}


func printEven(oddChan chan int , evenChan chan int, done chan bool){
	for even := range evenChan {
		time.Sleep(time.Second)
		fmt.Println("even :", even)
		if even == 10 {
			done <- true
			break
		}
		odd := even+1
		oddChan <- odd
	}
}


