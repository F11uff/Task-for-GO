package main

/*
Write an orDone function that routes data from the in channel to the return channel as long as the in channel
is open and the context is not canceled.
*/

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	var res []interface{}
	for v := range orDone(context.Background(), ch) {
		res = append(res, v)
	}

	fmt.Println(res)

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
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
