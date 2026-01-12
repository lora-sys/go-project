package concurrency

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDoWithTimeoutSuccess(t *testing.T) {
	err := DoWithTimeout(context.Background(), 50*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-time.After(10 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoWithTimeoutExceeded(t *testing.T) {
	err := DoWithTimeout(context.Background(), 10*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-time.After(50 * time.Millisecond):
			return errors.New("should not happen")
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestDoWithTimeoutCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	
	// 立即取消
	cancel()
	
	err := DoWithTimeout(ctx, 100*time.Millisecond, func(ctx context.Context) error {
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled, got %v", err)
	}
}

func TestDoWithTimeoutFunctionReturnsError(t *testing.T) {
	expectedErr := errors.New("function error")
	
	err := DoWithTimeout(context.Background(), 100*time.Millisecond, func(ctx context.Context) error {
		return expectedErr
	})
	
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestDoWithTimeoutImmediateReturn(t *testing.T) {
	start := time.Now()
	err := DoWithTimeout(context.Background(), 100*time.Millisecond, func(ctx context.Context) error {
		return nil // 立即返回
	})
	elapsed := time.Since(start)
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	// 应该立即返回，不应该等待超时
	if elapsed > 10*time.Millisecond {
		t.Fatalf("function should return immediately, took %v", elapsed)
	}
}

func TestDoWithTimeoutLongRunningNoTimeout(t *testing.T) {
	err := DoWithTimeout(context.Background(), 200*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-time.After(100 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}