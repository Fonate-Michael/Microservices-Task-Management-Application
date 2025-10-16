package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {

	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("URL parsing error: %v", err)

	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Custom director to modify the request path
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Strip the /task prefix from the path
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/task")
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/users")
	}

	return func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}

}
