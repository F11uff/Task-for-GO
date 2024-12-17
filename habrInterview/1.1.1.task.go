package main

import (
	"fmt"
	"math/rand"
)

/*
	На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
*/

func main() {
	slice1 := make([]int, 0, rand.Int()%20)
	slice2 := make([]int, 0, rand.Int()%10)
	var interSlice []int

	for i := 0; i < cap(slice1); i++ {
		slice1 = append(slice1, rand.Int()%20)
	}

	for i := 0; i < cap(slice2); i++ {
		slice2 = append(slice2, rand.Int()%20)
	}

	interSlice = intersectionSlice(slice1, slice2)
	fmt.Println(interSlice)
}

func intersectionSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 || len(slice2) == 0 {
		return []int{}
	}

	var interSlice []int
	mp := make(map[int]int)

	for _, value := range slice1 {
		mp[value]++
	}

	for _, value := range slice2 {
		if countEl, ok := mp[value]; ok && countEl > 0 {
			interSlice = append(interSlice, value)
			mp[value]--
		}
	}

	return interSlice
}
