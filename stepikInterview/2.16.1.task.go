package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Write the functions produce and main.

Produce function:
- receives the pipe channel as input
- writes integers to the pipe endlessly, starting from 0
- on a signal from main should terminate work

When finished, it should fall asleep for 3 seconds, then print “produce finished”

Main function:

- should create a pipe channel
- run produceCount functions produce and start reading from the pipe, printing each number
- when receiving a number, produceStop from pipe should stop reading new numbers from the pipe and should send
- a signal in the produce function that completes their work
should wait for all "produce finished" messages and print "main finished"

To implement the requirements, you can add additional arguments to the produce function
*/

const (
	produceCount = 3
	produceStop  = 10
)

func produce(ctx context.Context, wg *sync.WaitGroup, pipe chan<- int) { // допускается добавить доп. аргументы
	defer wg.Done()
	defer func() {
		time.Sleep(3 * time.Second)
		fmt.Println("produce finished")
	}()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		case pipe <- i:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	pipeCh := make(chan int)
	for i := 0; i < produceCount; i++ {
		wg.Add(1)
		go produce(ctx, &wg, pipeCh)
	}

	for i := range pipeCh {
		fmt.Println(i)

		if i == produceStop {
			cancel()
			break
		}
	}

	wg.Wait()
	fmt.Println("main finished")
	close(pipeCh)
}
