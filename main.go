package main

import (
	"github.com/matrixback/distributed-cache/tcp"
	"github.com/matrixback/distributed-cache/cache"
)

func main() {
	cache := cache.NewMemoryCache()
	server := tcp.NewServer(cache)
	server.Serve()
}