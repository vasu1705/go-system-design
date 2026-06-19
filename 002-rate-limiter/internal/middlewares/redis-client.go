package middlewares

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var rateLimitLuaScript = redis.NewScript(`
	local current = redis.call("GET", KEYS[1])
	if current and tonumber(current) >= tonumber(ARGV[1]) then
		return 0
	else
		current = redis.call("INCR", KEYS[1])
		if tonumber(current) == 1 then
			redis.call("EXPIRE", KEYS[1], ARGV[2])
		end
		return 1
	end
`)

type RedisRateLimiter struct {
	rdb *redis.Client
}

func NewRedisRateLimiter(redisAddr string) *RedisRateLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     "",  // no password by default in dev docker
		DB:           0,   // default DB
		PoolSize:     100, // Capping max concurrent TCP connections!
		MinIdleConns: 10,  // Maintain warm connections ready to use

	})

	return &RedisRateLimiter{
		rdb: rdb,
	}
}

func (rrl *RedisRateLimiter) CheckRateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 50*time.Millisecond)
		defer cancel()
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		key := "rate:" + ip
		limit, window := 5, 60

		allowed, err := rateLimitLuaScript.Run(ctx, rrl.rdb, []string{key}, limit, window).Int()
		if err != nil {
			fmt.Printf("Error executing Lua script: %v", err)
			next.ServeHTTP(w, r)
			return
		}
		if allowed == 0 {
			w.Header().Set("Retry-After", "60")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests) // 429
			return
		}
		fmt.Printf("Request allowed for IP: %s\n", ip)
		// Proceed to the next handler if allowed
		next.ServeHTTP(w, r)

	})
}
