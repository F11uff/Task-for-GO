package main

import (
	"context"
	"reflect"
)

/*
Write a function bridge that reads the next channel from the channel ins and forwards the data from the read channel
to the returned channel as long as the read channel is open and the context is not canceled.
The function continues to run as long as the channel ins is open and the context is not canceled.
*/

func main() {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}
	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
	}

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}

func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		for i := range ins {

			if i == nil {
				continue
			}

			select {
			case <-ctx.Done():
				return
			default:
			loop:
				for {
					select {
					case <-ctx.Done():
						return
					case val, ok := <-i:
						if !ok {
							break loop
						}

						ch <- val
					}
				}
			}
		}
	}()

	return ch
}
