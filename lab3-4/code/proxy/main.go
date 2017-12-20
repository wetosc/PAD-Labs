package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:3001")
	proxy1 := httputil.NewSingleHostReverseProxy(u)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy1.ServeHTTP(w, r)
	})
	http.ListenAndServe(":3000", nil)
}
