package concurrency

import (
	"context"
	"testing"
	"time"
	"fmt"
)

func TestProcessWithPoolEmptyInput(t *testing.T) {
	results, err := ProcessWithPool(context.Background(), []int{}, 5)
	if err != nil {
		t.Fatalf("unexpected error for empty input: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected empty results, got %v", results)
	}
}

func TestProcessWithPoolMoreWorkersThanInputs(t *testing.T) {
	inputs := []int{1, 2}
	workers := 5 // 比输入数量多
	
	results, err := ProcessWithPool(context.Background(), inputs, workers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expected := []int{1, 4}
	for i, v := range expected {
		if results[i] != v {
			t.Fatalf("result[%d]=%d want %d", i, results[i], v)
		}
	}
}

func TestProcessWithPoolConcurrentSafety(t *testing.T) {
	// 测试大并发量
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i + 1
	}
	
	start := time.Now()
	results, err := ProcessWithPool(context.Background(), inputs, 10)
	elapsed := time.Since(start)
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	// 验证所有结果
	for i, v := range inputs {
		if results[i] != v*v {
			t.Fatalf("result[%d]=%d want %d", i, results[i], v*v)
		}
	}
	
	t.Logf("Processed 1000 items with 10 workers in %v", elapsed)
}

func TestProcessWithPoolEarlyCancellation(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    
    // 更早取消（1ms后）
    go func() {
        time.Sleep(time.Millisecond * 1)
        cancel()
    }()
    
    // 创建更多数据（确保处理不完）
    inputs := make([]int, 100000)  // 10万条数据
    for i := 0; i < 100000; i++ {
        inputs[i] = i
    }
    
    // 使用较少的worker，让处理更慢
    _, err := ProcessWithPool(ctx, inputs, 2)
    if err == nil {
        t.Fatal("expected cancellation error")
    }
    if err != context.Canceled {
        t.Fatalf("expected context.Canceled, got %v", err)
    }
}

// 基准测试
func BenchmarkProcessWithPool(b *testing.B) {
	ctx := context.Background()
	
	for n := 0; n < b.N; n++ {
		inputs := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			inputs[i] = i
		}
		
		_, err := ProcessWithPool(ctx, inputs, 10)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkProcessWithPoolDifferentWorkerCounts(b *testing.B) {
	ctx := context.Background()
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}
	
	workerCounts := []int{1, 2, 4, 8, 16, 32, 64}
	
	for _, workers := range workerCounts {
		b.Run(fmt.Sprintf("Workers%d", workers), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := ProcessWithPool(ctx, inputs, workers)
				if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}