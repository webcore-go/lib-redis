package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/webcore-go/webcore/infra/config"
)

// Redis represents shared Redis connection
type Redis struct {
	Client *redis.Client
}

// NewRedis creates a new Redis connection
func NewRedis(config config.RedisConfig) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return &Redis{Client: client}
}

func (r *Redis) Install(args ...any) error {
	// Tidak melakukan apa-apa
	return nil
}

func (r *Redis) Connect() error {
	// Test connection
	_, err := r.Client.Ping(r.Client.Context()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return nil
}

// Close closes the Redis connection
func (r *Redis) Disconnect() error {
	return r.Client.Close()
}

func (r *Redis) Uninstall() error {
	// Tidak melakukan apa-apa
	return nil
}

func (r *Redis) Set(key string, value any, ttl time.Duration) error {
	ctx := r.Client.Context()
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(key string) (any, bool) {
	ctx := r.Client.Context()
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	}
	if err != nil {
		return nil, false
	}
	return val, true
}
