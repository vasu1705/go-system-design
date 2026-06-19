package middlewares

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mutex     sync.Mutex
	limit     int
	cache_map map[string]userRateLimit
}

type userRateLimit struct {
	count int
	time  int64
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		limit:     limit,
		cache_map: make(map[string]userRateLimit),
	}
}

func (rl *RateLimiter) CheckIfAllowed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		rl.mutex.Lock()

		if userRate, exists := rl.cache_map[ip]; exists {
			if time.Now().Unix()-userRate.time > 10 {
				rl.cache_map[ip] = userRateLimit{count: 0, time: time.Now().Unix()}
			}
			userRate = rl.cache_map[ip]
			if userRate.count >= rl.limit {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				rl.mutex.Unlock()
				return
			}
		} else {
			rl.cache_map[ip] = userRateLimit{count: 0, time: time.Now().Unix()}
		}
		rl.cache_map[ip] = userRateLimit{count: rl.cache_map[ip].count + 1, time: rl.cache_map[ip].time}
		rl.mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}
