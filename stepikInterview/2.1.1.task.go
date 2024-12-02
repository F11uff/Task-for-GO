package main

/*
Write merge and fillChan functions.

fillChan function:
- receives an integer n as input;
- returns channel;
- writes n numbers from 0 to n-1 to this channel.

merge function:
- receives an array of cs channels as input;
- returns channel;
- reads in parallel from each channel from cs and writes the received value to the returned channel.
*/

import (
	"fmt"
	"sync"
)

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)

	d := merge(a, b, c)

	fmt.Println()
	for i := range d {
		fmt.Print(i)
	}

}

func fillChan(n int) <-chan int {
	ch := make(chan int, n)
	defer close(ch)

	for i := 0; i < n; i++ {
		ch <- i
	}

	return ch
}

func merge(ch ...<-chan int) <-chan int {
	chA := make(chan int)

	wq := sync.WaitGroup{}

	for _, c := range ch {
		wq.Add(1)
		go func(ch <-chan int) {
			defer wq.Done()
			for v := range ch {
				chA <- v
			}
		}(c)
	}

	go func() {
		wq.Wait()
		close(chA)
	}()

	return chA
}
