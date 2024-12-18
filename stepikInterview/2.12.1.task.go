package main

import (
	"errors"
	"sync"
)

/*
Write a Run function that starts concurrent execution of fs functions and waits for them to finish. If one or
several fs functions failed, Run returns any of them.
*/

type fn func() error

func main() {
	expErr := errors.New("error")

	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}
func Run(fs ...fn) error {
	errorCh := make(chan error, len(fs))
	wg := sync.WaitGroup{}

	for _, f := range fs {
		wg.Add(1)
		go func(f fn) {
			defer wg.Done()
			errorCh <- f()
		}(f)
	}

	go func() {
		wg.Wait()
		close(errorCh)
	}()

	for e := range errorCh {
		if e != nil {
			return e
		}
	}

	return nil
}
