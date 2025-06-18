package common

import (
	"sync"
	"time"
)

// RateLimiter implements a simple in-memory rate-limiting mechanism.
type RateLimiter struct {
	mu     sync.Mutex
	limits map[string]time.Time
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limits: make(map[string]time.Time),
	}
}

func (r *RateLimiter) Allow(key string, cooldown time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if lastTime, exists := r.limits[key]; exists {
		if now.Sub(lastTime) < cooldown {
			return false
		}
	}
	r.limits[key] = now
	return true
}