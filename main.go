package main

import (
	"github.com/matrixback/distributed-cache/http"
)
func main() {
	server := http.NewServer()
	server.Serve()
}