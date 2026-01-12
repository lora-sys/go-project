package concurrency

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	var wg sync.WaitGroup
	c := &SafeCounter{}
	total := 1000

	wg.Add(total) // 等待的100个协程
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done() // 完成后调用Done
			c.Inc()  // 安全地递增计数器
		}()
	}
	wg.Wait() // 等待所有协程完成

	// 检查计数器的最终值是否正确

	if got := c.Value(); got != total {
		t.Fatalf("counter = %d, want %d", got, total)
	}
}