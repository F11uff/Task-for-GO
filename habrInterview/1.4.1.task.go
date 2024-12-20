package main

import (
	"fmt"
)

//Реализовать алгоритм бинарного поиска(для отсортированного массива).

func main() {
	slice := []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}
	if anw, ok := binarySearch(0, slice); ok {
		fmt.Println(anw)
	} else {
		fmt.Println("Not found")
	}
}

func binarySearch(value int, slice []int) (int, bool) {
	start, end := 0, len(slice)

	for start <= end {
		middle := ((end - start) / 2) + start

		if value == slice[middle] {
			return slice[middle], true
		} else if value > slice[middle] {
			start = middle + 1
		} else if value < slice[middle] {
			end = middle - 1
		}
	}

	return -1, false
}
