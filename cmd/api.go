package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	repo "github.com/kannape/go-e-commerce-app/internal/adapters/postgresql/sqlc"
	"github.com/kannape/go-e-commerce-app/internal/products"
)

type application struct {
	config config
	db     *pgx.Conn
}

func (app *application) mount() http.Handler {
	router := http.DefaultServeMux

	router.HandleFunc("GET /info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	router.HandleFunc("GET /products", productHandler.ListProducts)
	router.HandleFunc("GET /product", productHandler.FindProductByID)

	return router
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Listening on %s...", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	domainStr string
}
