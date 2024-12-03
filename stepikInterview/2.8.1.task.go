package main

/*
There is a function executeTask that can hang for an indefinite period of time during execution.
Implement a wrapper function executeTaskWithTimeout that:
executes executeTask
takes a context as an argument
terminates either as a result of executing executeTask or as a result of canceling the context. In the latter case,
return a context error.
*/

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	ch := make(chan interface{})

	go func() {
		executeTask()
		ch <- interface{}(nil)
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		return nil
	}
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
