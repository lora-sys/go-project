package concurrency

import (
	"context"
	"time"
)

// DoWithTimeout 在指定的超时时间内执行函数
func DoWithTimeout(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	// 创建带超时的子context
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel() // 确保资源被清理
	
	// 创建结果通道
	resultCh := make(chan error, 1)
	
	// 在goroutine中执行函数
	go func() {
		// 执行函数
		err := fn(ctx)
		
		select {
		case resultCh <- err:
			// 成功发送结果
		case <-ctx.Done():
			// 如果context已取消，丢弃结果
		}
	}()
	
	// 等待结果或超时
	select {
	case <-ctx.Done():
		// 超时或被取消
		return ctx.Err()
	case err := <-resultCh:
		// 函数执行完成
		return err
	}
}