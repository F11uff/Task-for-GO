package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Write a download function.
Download function:
the input receives download addresses - urls
competitively downloads information from each url (to download, use the fakeDownload function)

if calls to fakeDownload return errors, then you need to return them all (see errors.Join)
*/

// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
// timeoutLimit = 100 - ошибок не будет;
// timeoutLimit = 0 - всегда будет возвращаться ошибка.
const timeoutLimit = 100

type Result struct {
	msg string
	err error
}

func main() {
	msgs, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(msgs)
}

func download(urls []string) ([]string, error) {
	wq := sync.WaitGroup{}
	downlReq := make([]string, 0, len(urls))
	ch := make(chan Result, len(urls))
	var err error

	for _, url := range urls {
		wq.Add(1)
		go func(url string) {
			defer wq.Done()
			ch <- fakeDownload(url)
		}(url)
	}

	go func() {
		wq.Wait()
		close(ch)
	}()

	for result := range ch {
		if result.err != nil {
			err = errors.Join(err, result.err)
		}

		downlReq = append(downlReq, result.msg)
	}

	return downlReq, err
}

func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)

	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}
	return Result{
		msg: fmt.Sprintf("downloaded data from %s\n", url),
	}
}
