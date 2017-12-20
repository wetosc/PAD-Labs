package main

import (
	"fmt"
	"net/http"
	// "github.com/go-redis/redis"
)

var i = Index{I: 0, Max: 2}

func main() {
	launchClient()
	generateProxies()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.EscapedPath() + "?" + r.URL.RawQuery
		fmt.Printf("\nRead for key: %v", url)

		val, err := redisC.Get(url).Result()
		if CheckErr(err, "Error reading from Redis") {
			fmt.Println("\nFrom REDIS")
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, val)
			return
		}

		fmt.Println("\nFrom SERVER")
		proxies[i.Inc()].ServeHTTP(w, r)
	})

	http.ListenAndServe(":3000", nil)
}

func CheckErr(err error, m string) bool {
	if err != nil {
		fmt.Printf("\n| Error | %v : %v", m, err)
		return false
	}
	return true
}
