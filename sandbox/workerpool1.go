package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := 5
	jobsChan := make(chan int, jobs)
	results := make(chan int, jobs)

	totalWorker := 3

	for i := 1; i <= totalWorker; i++ {
		go worker(i, results, jobsChan)
	}

	for i := 1; i <= jobs; i++ {
		jobsChan <- i
	}
	close(jobsChan)

    for a := 1; a <= jobs; a++ {
		fmt.Println("result:  ", <-results)
	}
	

}

func worker(i int, results chan int, jobsChan chan int) {
	for j := range jobsChan {
		fmt.Println("worker", i, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", i, "finished job", j)
		results <- j * 2
	}
}
