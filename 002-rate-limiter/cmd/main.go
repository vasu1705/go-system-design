package main

import (
	"net/http"

	"github.com/vasu1705/go-system-design/002-rate-limiter/internal/middlewares"
	"github.com/vasu1705/go-system-design/002-rate-limiter/internal/typicode"
)

func main() {
	server := http.NewServeMux()
	typicodeService := typicode.NewTypicodeService(&http.Client{})
	typicodeHandler := typicode.NewTypicodeHandler(typicodeService)
	rateLimiter := middlewares.NewRateLimiter(10)
	rateLimiterRedis := middlewares.NewRedisRateLimiter("localhost:6379")
	rateLimiterRedisHandler := rateLimiterRedis.CheckRateLimit(http.HandlerFunc(typicodeHandler.GetUsers))
	server.Handle("/users", rateLimiter.CheckIfAllowed(rateLimiterRedisHandler))

	http.ListenAndServe(":8080", server)
}
