package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	fmt.Println("--->starting worker:", workerId)
	for job := range jobs {
		fmt.Printf("job %d is processed by worker %d \n", job, workerId)
		time.Sleep(time.Second * 2)
		fmt.Println("value in job:", job)
	}
}

func main() {
	fmt.Println("Hello World")
	totalWorkers := 5
	totalJobs := 15
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(jobs, &wg, i)
	}

	for i := 0; i < totalJobs; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Println("All jobs are sent to workers, closing the channel")

	wg.Wait()

}
