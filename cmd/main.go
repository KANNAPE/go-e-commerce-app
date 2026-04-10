package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("こんいちは、ワールド")

	cfg := config{
		addr: ":8080",
		db: dbConfig{},
	}

	api := application{
		config: cfg,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		log.Fatal("Server has failed to start!")
		os.Exit(1)
	}
}
