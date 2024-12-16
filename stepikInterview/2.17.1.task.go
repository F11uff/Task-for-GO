package main

import (
	"fmt"
	"sync"
)

/*
Implement the once structure, the new function, and the thread-safe do method.
The implementation of once and new should use pipes, do not use the sync package.
The new function returns a pointer to a once structure
do method:
- receives function f as input
- executes f only if do is called for the first time on this once instance. Other wise does nothing


The main function must print the call to the console exactly once.
*/

const goroutinesNumber = 10

type once struct {
	ch chan struct{}
}

func new() *once {
	ch := make(chan struct{}, 1)
	defer close(ch)
	ch <- struct{}{}
	return &once{ch: ch}
}

func (o *once) do(f func()) {
	if _, ok := <-o.ch; ok {
		f()
	}
}

func funcToCall() {
	fmt.Printf("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()
	wg.Add(goroutinesNumber)
	for i := 0; i < goroutinesNumber; i++ {
		go func(f func()) {
			defer wg.Done()
			so.do(f)
		}(funcToCall)
	}
	wg.Wait()
}
