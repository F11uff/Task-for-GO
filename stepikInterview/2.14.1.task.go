package main

func main() {
	first := make(chan int)
	last := make(<-chan int)
	n := 10

	last = inc(first)
	for i := 0; i < n-1; i++ {
		last = inc(last)
	}

	first <- 0

	close(first)
	if n != <-last {
		panic("wrong code")
	}
}
func inc(in <-chan int) <-chan int {
	ch := make(chan int, len(in))

	go func() {
		for i := range in {
			ch <- i + 1
		}

		close(ch)
	}()

	return ch
}
