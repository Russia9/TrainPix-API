package main

import (
	"os"
	"strconv"
	"trainpix-api/api"
)

func main() {
	port := 8080
	if os.Getenv("API_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("API_PORT"))
	}
	api.Api(port)
}
