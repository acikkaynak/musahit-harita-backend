package cache

import (
	log "github.com/acikkaynak/musahit-harita-backend/pkg/logger"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type RedisStore struct {
	Client *redis.Client
}

func NewRedisStore() *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &RedisStore{Client: client}
}

func (r *RedisStore) SetKey(key string, value []byte, ttl time.Duration) {
	err := r.Client.Set(r.Client.Context(), key, value, ttl).Err()
	if err != nil {
		log.Logger().Info("Unable to set key in redis" + key + err.Error())
	}
}

func (r *RedisStore) Get(key string) []byte {
	get := r.Client.Get(r.Client.Context(), key)

	resp, err := get.Bytes()
	if err != nil {
		log.Logger().Info("Unable to get key in redis" + key + err.Error())
	}

	return resp
}

func (r *RedisStore) Delete(key string) error {
	return r.Client.Del(r.Client.Context(), key).Err()
}

func (r *RedisStore) DeleteAll() error {
	return r.Client.FlushDB(r.Client.Context()).Err()
}

func (r *RedisStore) Close() {
	r.Client.Close()
}
