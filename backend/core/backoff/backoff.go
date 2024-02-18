package backoff

import (
	"math"
	"time"
)

type Backoff func(time.Duration) time.Duration

func ExpBackoff(wait time.Duration) time.Duration {
	return wait * 2
}

func NoopBackoff(wait time.Duration) time.Duration {
	return wait
}

func (backoff Backoff) WithUpperLimit(high time.Duration) Backoff {
	return backoff.WithLimit(0, high)
}

func (backoff Backoff) WithLowerLimit(low time.Duration) Backoff {
	return backoff.WithLimit(low, math.MaxInt64)
}

func (backoff Backoff) WithLimit(low time.Duration, high time.Duration) Backoff {
	return func(d time.Duration) time.Duration {
		duration := backoff(d)
		if duration < low {
			return low
		} else if duration > high {
			return high
		}
		return duration
	}
}
