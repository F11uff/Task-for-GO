package main

import (
	"fmt"
	"sync"
	"time"
)

/*

You need to write a worker pool: you need to execute numJobs of jobs in parallel, using numWorkers goroutines,
which are launched once during the execution of the program.
To do this, write the worker and main functions.
The worker function: receives as input a function for executing f, a channel for receiving arguments jobs and a channel
for writing results reads from jobs and writes the result of executing f(job) to results.
Main function: runs the worker function in numWorkers goroutines;
the worker uses the multiplier function as the first argument; writes numbers from 1 to numJobs to the jobs channel;
reads and displays the received values from the results channel, in parallel with the work of workers
*/

const (
	numJobs    = 5
	numWorkers = 3
)

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}

	ch2 := make(chan int)
	go func() {
		for r := range results {
			fmt.Println(r)
		}
		close(ch2)
	}()

	for i := 0; i < numJobs; i++ {
		time.Sleep(2 * time.Second)
		jobs <- i
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	<-ch2
}

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}
}
