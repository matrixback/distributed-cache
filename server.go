package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	cache *MemoryCache
}

func NewServer() *Server {
	cache := NewMemoryCache()
	return &Server{cache: cache}
}

func (s *Server) Serve() {
	http.HandleFunc("/cache", CacheHandle)
	fmt.Println("cache server is starting, listen at 6740...")
	http.ListenAndServe(":6740", nil)
}

func CacheHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("hello, matrix"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}