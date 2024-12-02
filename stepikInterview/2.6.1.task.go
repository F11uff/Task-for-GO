package main

/*
Write a mergeSorted function that takes two "sorted" channels as input and returns the resulting
sorted channel.
*/

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func mergeSorted(a, b <-chan int) <-chan int {
	ch := make(chan int, len(a)+len(b))
	wg := sync.WaitGroup{}
	slice := make([]int, 0, len(a)+len(b))

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range a {
			slice = append(slice, i)
		}

	}()

	go func() {
		time.Sleep(3 * time.Second)
		defer wg.Done()
		for i := range b {
			slice = append(slice, i)
		}
	}()

	wg.Wait()

	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range slice {
			ch <- slice[i]
		}
		close(ch)
	}()

	return ch
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}
func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}
func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)
	for val := range c {
		fmt.Printf("%d ", val)
	}
}
