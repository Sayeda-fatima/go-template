package common

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type FrequencyLimiter struct {
	requests  map[string][]time.Time
	cooldowns *cache.Cache
	window    time.Duration
	limit     int
	cooldown  time.Duration
}

func NewFrequencyLimiter(limit int, window time.Duration, cooldown time.Duration) *FrequencyLimiter {
	return &FrequencyLimiter{
		requests:  make(map[string][]time.Time),
		cooldowns: cache.New(cooldown, 1*time.Minute),
		window:    window,
		limit:     limit,
		cooldown:  cooldown,
	}
}

func (f *FrequencyLimiter) Allow(key string) bool {
	now := time.Now()

	if _, found := f.cooldowns.Get(key); found {
		return false
	}

	times := f.requests[key]
	cutoff := now.Add(-f.window)
	pruned := times[:0]
	for _, t := range times {
		if t.After(cutoff) {
			pruned = append(pruned, t)
		}
	}

	if len(pruned) >= f.limit {
		f.cooldowns.Set(key, struct{}{}, f.cooldown)
		f.requests[key] = pruned
		return false
	}

	f.requests[key] = append(pruned, now)
	return true
}
