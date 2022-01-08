package http

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"

	"github.com/matrixback/distributed-cache/cache"
)

type Server struct {
	*cache.MemoryCache
}

func NewServer() *Server {
	return &Server{MemoryCache: cache.NewMemoryCache()}
}

func (s *Server) Serve() {
	http.Handle("/cache/", &CacheHandler{Server: s})
	fmt.Println("cache server is starting, listen at 6740...")
	http.ListenAndServe(":6740", nil)
}

type CacheHandler struct {
	*Server
}

func (ch *CacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		val, err := ch.Server.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if val == nil || len(val) <= 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(val)
		return
	}

	if r.Method == http.MethodPut {
		val, _ := ioutil.ReadAll(r.Body)
		if len(val) <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := ch.Server.Set(key, val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}