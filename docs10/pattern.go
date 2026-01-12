package patterns

import (
	"context"
	"errors"
	"sync"
	"time"
)

func PipelineDoubleThenAdd(in <-chan int) <-chan int{
	statage1 := func(input <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range input {
			out <- v * 2
		}
	}()
	return out
	}
   statage2 := func(input <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v:= range input{
			out <- v+1
		}
	}()
	return out
   }
   return statage2(statage1(in))
}

func FanOutSquare(ctx context.Context,in <- chan int,workerCount int) <- chan int{
	out := make(chan int)
	if workerCount <= 0 {
     close(out)
	 return out
	}
    var wg sync.WaitGroup
	worker := func() {
		//保证一定会关闭
		defer wg.Done()
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v*v:
			}
		}
	}
    wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}



func SendWithTimeout(ctx context.Context, ch chan<- int, v int, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = time.Nanosecond
	}
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case ch <- v:
		return nil
	case <-timer.C:
		return errors.New("send timeout")
	}
}