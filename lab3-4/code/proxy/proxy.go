package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var proxies []*httputil.ReverseProxy

func generateProxies() {
	proxies = make([]*httputil.ReverseProxy, 0)

	u, _ := url.Parse("http://localhost:3001")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ModifyResponse = CacheResponse
	proxies = append(proxies, proxy)

	u, _ = url.Parse("http://localhost:3002")
	proxy = httputil.NewSingleHostReverseProxy(u)
	proxy.ModifyResponse = CacheResponse
	proxies = append(proxies, proxy)
}

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

func responseBody(r *http.Response) string {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyString
}
