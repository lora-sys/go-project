package concurrency

import (
	"sync"

)

type SafeCounter struct {
	mu  sync.Mutex // 创建互斥锁
	n int
}
// Inc increments the counter safely.
func (c *SafeCounter) Inc() {
	c.mu.Lock() // 上锁
	c.n++
	c.mu.Unlock() // 释放锁
}
// Value returns the current count.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// 互斥锁
// 应用场景，网站统计次数，线程池计数，限流