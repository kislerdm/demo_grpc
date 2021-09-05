package main

import (
	"frontned/server"
	"os"
)

var s *server.Server

func init() {
	s = server.New()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}
