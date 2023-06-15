package utils

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

var Cache *redis.Client

func NewCache() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		log.Fatal(err, "Error converting CACHE_DB to int")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_URL"),
		Password: os.Getenv("CACHE_PASSWORD"), // no password set
		DB:       db,                          // use default DB
	})

	Cache = rdb

	return rdb
}

func SetCache(key string, values interface{}, duration time.Duration) {
	Cache.Set(context.Background(), key, values, duration)
}

func GetFromCache(key string, values interface{}) interface{} {
	val, err := Cache.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return nil
	}

	if err != nil {
		log.Fatal(err, "Error getting cache value")
	}

	err = json.Unmarshal([]byte(val), values)
	if err != nil {
		log.Fatal(err, "Error unmarshalling cache value")
	}

	return values
}
