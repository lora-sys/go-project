package main

import (
	"errors"
	"fmt"
)

// average 计算可变参数的平均值，若为空返回错误
func Average(nums ...float64) (float64,error) {
	if(len(nums)==0) {
		return 0 , errors.New("no numbers provided")
	}
	var sum float64
	for _,n := range nums {
		sum +=n
	}

	return sum / float64(len(nums)),nil
}


// Pred 谓词函数类型， 用以过滤切片
type Pred[T any] func(T) bool

func Filter[T any](xs []T,pred Pred[T]) []T {
	var out []T
	for _,v := range xs{
		if pred(v) {
			out = append(out,v)
		}
	}
	return out
}
// NewCounter 返回一个闭包，每次调用返回递增值
func NewCounter(start int) func() int {
	current := start
	return func() int {
		val := current
		current++
		return val
	}
}


func main() {
	avg,err := Average(1,2,3,4,5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(avg)
	}

	evens := Filter([]int{1,2,3,4,5}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens)

	counter := NewCounter(1)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
