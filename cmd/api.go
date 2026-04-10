package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	router := http.DefaultServeMux

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})

	//http.ListenAndServe(":8080", router)

	return router
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Listening on %s...", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db 	 dbConfig
}

type dbConfig struct {
	domainStr string
}
