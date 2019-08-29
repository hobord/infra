package ctxrouter

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func loadConfig() {

}

func RouterHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin, _ := url.Parse("http://netlab.hu")

		proxy := NewSingleHostReverseProxy(origin)
		r.Host = r.URL.Host

		proxy.ServeHTTP(w, r)
	})
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = join([]string{target.Path, req.URL.Path})
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}
	return &httputil.ReverseProxy{Director: director}
	// proxy := httputil.NewSingleHostReverseProxy(target)
	// return proxy
}

func join(elem []string) string {
	for i, e := range elem {
		if e != "" {
			return strings.Join(elem[i:], "/")
		}
	}
	return ""
}
