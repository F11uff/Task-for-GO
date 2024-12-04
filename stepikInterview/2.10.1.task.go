package main

/*
Write a function tee that routes data from the in channel to both return channels simultaneously
(i.e., they receive the same data) while the in channel is open and the context is not canceled.
Hint: use the orDone from the previous problem to simplify the code.
*/

import (
	"context"
	"reflect"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	i := 0
	inc := func() interface{} {
		i++
		return i
	}
	out1, out2 := tee(ctx, take2(ctx, repeatFn2(ctx, inc), 3))
	var res1, res2 []interface{}
	for val1 := range out1 {
		res1 = append(res1, val1)
		res2 = append(res2, <-out2)
	}
	exp := []interface{}{1, 2, 3}

	if !reflect.DeepEqual(res1, exp) || !reflect.DeepEqual(res2, exp) {
		panic("wrong code")
	}
}
func tee(ctx context.Context, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1, out2 := make(chan interface{}), make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)
		for i := range orDone2(ctx, in) {
			out1 <- i
			out2 <- i
		}
	}()

	return out1, out2
}

func repeatFn2(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}()

	return out
}

func take2(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func orDone2(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
				case out <- val:
				}
			}
		}
	}()

	return out
}
