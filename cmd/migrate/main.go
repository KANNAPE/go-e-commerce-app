package main

import (
	"github.com/kannape/go-e-commerce-app/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic("Can't init server")
	}
}
