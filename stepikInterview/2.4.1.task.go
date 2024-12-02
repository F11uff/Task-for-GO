package main

import (
	"context"
	"fmt"
	"math/rand"
)

/*
Write repeatFn and take functions.
repeatFn calls fn infinitely and writes its result to the return channel.
Terminates early if the context is canceled.

Take reads at most num from in while in is open and writes the value to the return channel.
Terminates early if the context is canceled.
*/

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()

	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	ch := make(chan interface{}, num)

	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				ch <- val
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} { return rand.Int() }
	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, rand), 3) {
		res = append(res, num)
		fmt.Println(res)
	}
	if len(res) != 3 {
		panic("wrong code")
	}
}
