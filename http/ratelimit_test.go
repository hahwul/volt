package http

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(100 * time.Millisecond)

	key := "test-key"

	start := time.Now()
	rl.Block(key)
	elapsed := time.Since(start)

	if elapsed > 10*time.Millisecond {
		t.Errorf("expected initial block to be instantaneous, but took %v", elapsed)
	}

	start = time.Now()
	rl.Block(key)
	elapsed = time.Since(start)

	if elapsed < 90*time.Millisecond || elapsed > 110*time.Millisecond {
		t.Errorf("expected block to be approximately 100ms, but took %v", elapsed)
	}

	time.Sleep(100 * time.Millisecond)

	start = time.Now()
	rl.Block(key)
	elapsed = time.Since(start)

	if elapsed > 10*time.Millisecond {
		t.Errorf("expected block after delay to be instantaneous, but took %v", elapsed)
	}
}
