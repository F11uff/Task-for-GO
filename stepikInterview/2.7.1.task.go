package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	msg string
	err error
}

type searh func() *result
type replicas []searh

/*
Напишите функцию getFirstResult, которая принимает контекст и запускает конкурентый поиск, возвращая первый
доступный результат из replicas. Возвращать ошибку контекста, если контекст завершился раньше, чем стал доступен
какой-то результат из реплики.
Напишите функцию getResults, которая запускает конкурентный поиск для каждого набора реплик из replicaKinds,
использую getFirstResul`t, и возвращает результат для каждого набора реплик.
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}
	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}

func fakeSearch(kind string) searh {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	ch := make(chan *result)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, replica := range replicas {
		wg.Add(1)
		go func(replica searh) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Println("cancelled")
			case ch <- replica():
			}
		}(replica)
	}

	select {
	case <-ctx.Done():
		return &result{err: ctx.Err()}
	case r := <-ch:
		return r
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	wg := sync.WaitGroup{}
	ch := make(chan *result)

	for _, replicaKind := range replicaKinds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- getFirstResult(ctx, replicaKind)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res []*result
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				return res
			}
			res = append(res, val)
		}
	}
}