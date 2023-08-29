package main

import (
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

// go.uber.org/ratelimit TestQPS
//
//	rl_test.go:30: Received number of events is lesser than expected. Got: 1905835 (95.29%); Expected: 2000000
//	rl_test.go:30: Received number of events is lesser than expected. Got: 3921179 (98.03%); Expected: 4000000
//	rl_test.go:30: Received number of events is lesser than expected. Got: 5597944 (93.30%); Expected: 6000000
func TestQPS(t *testing.T) {
	testLimiterQPS(t, 1000000)
	testLimiterQPS(t, 2000000)
	testLimiterQPS(t, 3000000)
}

func testLimiterQPS(t *testing.T, rate int) {
	limiter := ratelimit.New(rate)
	// limiter := ratelimit.NewLimiter(ratelimit.Limit(rate), 1)
	deadline := time.After(time.Millisecond * 2000)
	i := 0
	for {
		select {
		case <-deadline:
			expEvents := 2 * rate
			if i > expEvents {
				t.Errorf("Received number of events is bigger than expected. Got: %d; Expected: %d", i, expEvents)
			}

			expEventsPercent := (float64(i) / float64(expEvents)) * 100
			if expEventsPercent < 100 {
				t.Errorf("Received number of events is lesser than expected. Got: %d (%.2f%%); Expected: %d", i, expEventsPercent, expEvents)
			}
			return
		default:
			limiter.Take()
			i++
		}
	}
}
