package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

var redisC *redis.Client

func launchClient() {
	redisC = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// CacheResponse is the ModifyResponse of all proxies;
// The method is called by the proxy when it receives a response
func CacheResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK && r.Request.Method == http.MethodGet {
		key := r.Request.URL.EscapedPath() + "?" + r.Request.URL.RawQuery
		value := responseBody(r)
		err := redisC.Set(key, value, 10*time.Second).Err()
		CheckErr(err, "Error caching data")
		fmt.Printf("\nCached: %v", key)
	}
	return nil
}
