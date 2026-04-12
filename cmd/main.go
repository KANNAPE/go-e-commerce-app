package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kannape/go-e-commerce-app/internal/env"
)

func main() {
	fmt.Println("こんいちは、ワールド")

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			domainStr: env.GetString("GOOSE_DBSTRING", "user=postgres password=postgres host=localhost port=5432 dbname=go-e-commerce-app sslmode=disable"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.domainStr)
	if err != nil {
		slog.Error("Database connection failed!", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	logger.Info("Database connection successful!")

	api := application{
		config: cfg,
		db:     conn,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server has failed to start!", "error", err)
		os.Exit(1)
	}
}
