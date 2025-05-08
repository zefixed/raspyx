package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"raspyx/config"
	v1 "raspyx/internal/delivery/http/v1"
	"sync"
)

type RateLimiterStorage struct {
	limits map[string]*rate.Limiter
	mu     sync.Mutex
}

func NewRateLimiterStorage() *RateLimiterStorage {
	return &RateLimiterStorage{
		limits: make(map[string]*rate.Limiter),
	}
}

func (s *RateLimiterStorage) GetOrCreate(key string, r rate.Limit, b int) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limiter, exists := s.limits[key]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(r, b)
	s.limits[key] = limiter
	return limiter
}

func RateLimiter(ctx context.Context, rl config.RateLimiter, storage *RateLimiterStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		//storage.GetOrCreate(c.ClientIP(), rate.Limit(rl.Limit), rl.Burst).Wait(ctx)
		if !storage.GetOrCreate(c.ClientIP(), rate.Limit(rl.Limit), rl.Burst).Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, v1.RespError("too many requests, please try again later"))
			return
		}
		c.Next()
	}
}
