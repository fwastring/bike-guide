package database

import (
	"bike-guide-api/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisDB represents a Redis database connection
type RedisDB struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisDB creates a new Redis database connection
func NewRedisDB(cfg *config.Config) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisAddr(),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx := context.Background()

	// Test the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}

	return &RedisDB{
		client: client,
		ctx:    ctx,
	}, nil
}

// Close closes the Redis connection
func (r *RedisDB) Close() error {
	return r.client.Close()
}

// SetSession stores a session in Redis
func (r *RedisDB) SetSession(sessionID string, sessionData []byte, expiration time.Duration) error {
	key := fmt.Sprintf("session:%s", sessionID)
	return r.client.Set(r.ctx, key, sessionData, expiration).Err()
}

// GetSession retrieves a session from Redis
func (r *RedisDB) GetSession(sessionID string) ([]byte, error) {
	key := fmt.Sprintf("session:%s", sessionID)
	return r.client.Get(r.ctx, key).Bytes()
}

// DeleteSession deletes a session from Redis
func (r *RedisDB) DeleteSession(sessionID string) error {
	key := fmt.Sprintf("session:%s", sessionID)
	return r.client.Del(r.ctx, key).Err()
}

// SetKeyValue sets a key-value pair in Redis
func (r *RedisDB) SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

// GetKeyValue retrieves a value from Redis by key
func (r *RedisDB) GetKeyValue(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// DeleteKeyValue deletes a key-value pair from Redis
func (r *RedisDB) DeleteKeyValue(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// SetHash sets a hash in Redis
func (r *RedisDB) SetHash(key string, values map[string]interface{}) error {
	return r.client.HMSet(r.ctx, key, values).Err()
}

// GetHash retrieves a hash from Redis
func (r *RedisDB) GetHash(key string) (map[string]string, error) {
	return r.client.HGetAll(r.ctx, key).Result()
}

// DeleteHash deletes a hash from Redis
func (r *RedisDB) DeleteHash(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
