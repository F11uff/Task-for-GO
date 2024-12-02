package main

import (
	"context"
	"fmt"
)

/*
Write functions generator and squarer.

The generator function takes as input a context and a slice of integers, the elements of which are sequentially written into
return channel.

The squarer function takes as input a context and an integer channel. The function sequentially reads numbers from the channel,
squares them and writes to the returned channel.

Both functions must be able to terminate when the context is canceled.
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

	go func() {
		for _, val := range in {
			select {
			case <-ctx.Done():
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

	go func() {
		defer close(ch)
		for val := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- val * val:
			}
		}
	}()

	return ch
}
