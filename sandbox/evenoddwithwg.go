package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	evenChan := make(chan int)
	oddChan := make(chan int, 1)
	//done := make(chan bool)

	oddChan <- 1
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go printOdd(oddChan , evenChan)
	go printEven(oddChan, evenChan , wg)


	wg.Wait()

	// close(evenChan)
	// close(oddChan)
}


func printOdd(oddChan chan int , evenChan chan int){
	for odd := range oddChan {
		fmt.Println("odd :", odd)
		even := odd+1
		evenChan <- even
	}
}


func printEven(oddChan chan int , evenChan chan int, wg *sync.WaitGroup){
	for even := range evenChan {
		time.Sleep(time.Second)
		fmt.Println("even :", even)
		if even == 10 {
			wg.Done()
			close(evenChan)
			close(oddChan)
			break
		}
		odd := even+1
		oddChan <- odd
	}
}


