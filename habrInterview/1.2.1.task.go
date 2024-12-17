package main

import "fmt"

/*
	Сделать кастомную waitGroup на семафоре
*/

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := NewSemaphore(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)
}

func (s Semaphore) Inc(count int) {
	for i := 0; i < count; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Dec(count int) {
	for i := 0; i < count; i++ {
		<-s
	}
}
