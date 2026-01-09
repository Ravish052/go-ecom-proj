package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Ravish052/goEcon/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbconfig{
			dsn: env.GetString("GOOSE_DB_STRING", "host=localhost user=postgres password=postgres dbname=go_ecom_db sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//database configuration
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
		// slog.Error("server Failed to start", "error", err)
		// os.Exit(1)
	}

	defer conn.Close(ctx)

	logger.Info("database connected successfully", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	//structured logging

	if err := api.run(api.mount()); err != nil {
		slog.Error("server Failed to start", "error", err)
		os.Exit(1)
	}
}
