package main

import (
	"./http"
)
func main() {
	server := http.NewServer()
	server.Serve()
}