package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Index struct {
	sync.Mutex
	I   int32
	Max int32
}

func (i *Index) Inc() int32 {
	i.Lock()
	defer i.Unlock()

	i.I++
	if i.I >= i.Max {
		i.I = 0
	}
	return i.I
}

var i = Index{I: 0, Max: 2}

func main() {
	proxies := make([]*httputil.ReverseProxy, 0)
	u, _ := url.Parse("http://localhost:3001")
	proxies = append(proxies, httputil.NewSingleHostReverseProxy(u))
	u, _ = url.Parse("http://localhost:3002")
	proxies = append(proxies, httputil.NewSingleHostReverseProxy(u))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		proxies[i.Inc()].ServeHTTP(w, r)
	})
	http.ListenAndServe(":3000", nil)
}
