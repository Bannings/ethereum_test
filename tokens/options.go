package tokens

import (
	"math/rand"
	"time"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/core"
)

var (
	defaultRetriableErrors = []error{core.ErrInsufficientFunds}

	defaultOptions = &options{
		max: 1,
		retryIf: func(err error) bool {
			for _, e := range defaultRetriableErrors {
				if e == err {
					return true
				}
			}
			return false
		},
		backoff: backoffUniformRandom(50*time.Millisecond, 0.10),
	}
)

type RetryIfFunc func(error) bool

type BackoffFunc func(attempt uint) time.Duration

type options struct {
	max     uint
	retryIf RetryIfFunc
	backoff BackoffFunc
}

type ExecuteOption func(*options)

func WithRetryDisable() ExecuteOption {
	return WithRetryMax(0)
}

func WithRetryMax(maxRetries uint) ExecuteOption {
	return func(o *options) {
		o.max = maxRetries
	}
}

func WithRetryBackoff(bf BackoffFunc) ExecuteOption {
	return func(o *options) {
		o.backoff = bf
	}
}

func WithRetryIf(retryIf RetryIfFunc) ExecuteOption {
	return func(c *options) {
		c.retryIf = retryIf
	}
}

func backoffUniformRandom(backoff time.Duration, jitter float64) BackoffFunc {
	return func(attempt uint) time.Duration {
		multiplier := jitter * (rand.Float64() - 0.5) * 2
		return time.Duration(float64(backoff) * (1 + multiplier))
	}
}

func waitRetryBackoff(attempt uint, opts *options) {
	var waitTime time.Duration = 0
	if attempt > 0 {
		waitTime = opts.backoff(attempt)
	}
	if waitTime > 0 {
		log.Infof("executor retry attempt: %d, backoff for %v", attempt, waitTime)
		select {
		case <-time.Tick(waitTime):
		}
	}
}
