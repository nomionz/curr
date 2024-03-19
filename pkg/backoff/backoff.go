package backoff

import (
	"math"
	"time"
)

// Backoff is an interface for simple exp backoff implementation
type Backoff interface {
	Next() time.Duration
	Reset()
}

type backoff struct {
	factor float64
	max    time.Duration
	count  int
}

func NewBackoff() Backoff {
	return &backoff{
		factor: 2,
		max:    20 * time.Second,
		count:  0,
	}
}

func (b *backoff) Next() time.Duration {
	d := time.Duration(math.Pow(b.factor, float64(b.count))) * time.Second

	b.count++

	if d > b.max {
		d = b.max
	}

	return d
}

func (b *backoff) Reset() {
	b.count = 0
}
