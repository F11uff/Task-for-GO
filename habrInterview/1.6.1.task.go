package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Напишите программу, которая использует горутины для параллельного вычисления суммы элементов массива.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}

	runtime.GOMAXPROCS(3)

	wg.Add(2)
	go sum(arr[:len(arr)/2], ch, wg)
	go sum(arr[len(arr)/2:], ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	x, y := <-ch, <-ch

	fmt.Println(x + y)
}

func sum(numbers []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0

	for i := range numbers {
		sum += numbers[i]
	}

	ch <- sum
}
