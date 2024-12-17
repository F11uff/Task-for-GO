package main

import "fmt"

//ch1 <- 10, 20, 22, 44
//[10,20] or [13,20]  = example (ch1, batchSize)

/*
Реализовать функцию выполняющую батчинг значенийиз канала ch.
Для возвратабатчей использовать выходной канал chan []int *c - входной канал значений.
Из этих значений должны формироваться батчи batchSize - размер батчей
*/

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 10
		ch1 <- 20
		ch1 <- 22
		ch1 <- 44
	}()

	for i := range example(ch1, 2) {
		fmt.Println(i)
	}

}

func example(ch1 <-chan int, batchSize int) chan []int {

	chResult := make(chan []int)

	go func() {
		defer close(chResult)
		slice := make([]int, 0, batchSize)
		for i := range ch1 {
			slice = append(slice, i)
			if len(slice) == batchSize {
				chResult <- slice
				slice = make([]int, 0, batchSize)
			}
		}

		if len(slice) > 0 {
			chResult <- slice
		}
	}()

	return chResult
}
