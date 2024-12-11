package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	ch := make(chan int, size)
	return &ringBuffer{ch}
}

func (b *ringBuffer) write(v int) {
	select {
	case b.c <- v:
	default:
		<-b.c //Считать с канала, чтобы перезаписать другое число (буфер канал
		// не дает возможности переписать значения в кольцевой очери канала, из-за sendx)
		b.write(v)
	}
}

func (b *ringBuffer) close() {
	close(b.c)
}

func (b *ringBuffer) read() (v int, ok bool) {
	v, ok = <-b.c
	return v, ok
}

func main() {
	buff := newRingBuffer(3)
	for i := 1; i <= 6; i++ {
		buff.write(i)
	}
	buff.close()
	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}
	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
