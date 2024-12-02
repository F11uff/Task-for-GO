package main

import (
	"context"
	"fmt"
)

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
