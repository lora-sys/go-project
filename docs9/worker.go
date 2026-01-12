package concurrency

import (
	"context"
	"errors"
	"sync"
)

// ProcessWithPool 使用工作池处理输入数据
func ProcessWithPool(ctx context.Context, inputs []int, workers int) ([]int, error) {
	if workers <= 0 {
		return nil, errors.New("number of workers must be greater than zero")
	}
	
	// 如果输入为空，直接返回空结果
	if len(inputs) == 0 {
		return []int{}, nil
	}
	
	// 如果worker数大于输入数，调整为输入数
	if workers > len(inputs) {
		workers = len(inputs)
	}
	
	type job struct {
		idx int
		val int
	}
	
	jobs := make(chan job)
	results := make([]int, len(inputs))
	var wg sync.WaitGroup
	
	// 工作函数
	worker := func() {
		defer wg.Done()
		for j := range jobs {
			select {
			case <-ctx.Done():
				return
			default:
				// 模拟一些处理时间（实际使用时可能是网络请求或复杂计算）
				// time.Sleep(time.Millisecond * 10)
				results[j.idx] = j.val * j.val
			}
		}
	}
	
	// 启动worker池
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker()
	}
	
	// 发送任务
	for i, v := range inputs {
		// 检查context是否已取消
		if err := ctx.Err(); err != nil {
			close(jobs)
			wg.Wait()
			return results, err
		}
		
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return results, ctx.Err()
		case jobs <- job{idx: i, val: v}:
			// 成功发送任务
		}
	}
	
	// 所有任务发送完成，关闭通道
	close(jobs)
	
	// 等待所有worker完成
	wg.Wait()
	
	// 再次检查context错误（可能在worker处理期间取消）
	if err := ctx.Err(); err != nil {
		return results, err
	}
	
	return results, nil
}
