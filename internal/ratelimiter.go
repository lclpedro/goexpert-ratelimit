package internal

import (
	"time"
)

type RateLimiter struct {
	store      Store
	rate       int64
	expiration time.Duration
}

func NewRateLimiter(store Store, expiration time.Duration) *RateLimiter {
	return &RateLimiter{
		store:      store,
		expiration: expiration,
	}
}

func (l *RateLimiter) SetRate(rate int64) {
	l.rate = rate
}

func (l *RateLimiter) Allow(identifier string) bool {
	count, err := l.store.Increment(identifier, l.expiration)
	if err != nil {
		return false
	}

	return count <= l.rate
}
