package main

import (
	"context"
	"fmt"
)

/*
Напишите функции generator и squarer.

Функция generator принимает на вход контекст и слайс целых чисел, элементы которого последовательно записываются в
возвращаемый канал.

Функция squarer принимает на вход контекст и канал целых чисел.
Функция последовательно читает из канал числа,
возводит их в квадрат и пишет в возвращаемый канал.

Обе функции должны уметь завершаться по отмене контекста.
*/

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int, len(in))

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		for _, val := range in {
			select {
			case <-ctx.Done():
				cancel()
				return
			case ch <- val:
			}
		}

		close(ch)
	}()

	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	ch := make(chan int)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer close(ch)
		for val := range in {
			select {
			case <-ctx.Done():
				cancel()
				return
			case ch <- val * val:
			}
		}
	}()

	return ch
}
