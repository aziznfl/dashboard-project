package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/config"
	"github.com/durianpay/fullstack-boilerplate/internal/infrastructure/cache"
	"github.com/durianpay/fullstack-boilerplate/internal/infrastructure/database"
	srv "github.com/durianpay/fullstack-boilerplate/internal/service/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.Load()

	// 1. Database Setup
	db, err := sql.Open("sqlite3", cfg.DbSource+"?_foreign_keys=1")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer func() {
		log.Println("closing database connection...")
		db.Close()
	}()

	if err := database.InitDB(db); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// 2. Cache Setup
	redisCache, err := cache.NewRedisCache(cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Printf("warning: redis connection failed: %v", err)
	}
	// Note: redisCache might have a Close method we should defer if applicable

	// 3. API & Server Setup
	apiHandler, err := api.InitAPIHandler(db, redisCache, cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize API handler: %w", err)
	}

	server := srv.NewServer(apiHandler, cfg.OpenapiYamlLocation, cfg.AppEnv, cfg.JwtSecret)

	// 4. Start Server
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("starting server on %s", cfg.HttpAddress)
		if err := server.Start(cfg.HttpAddress); err != nil && err != http.ErrServerClosed {
			serverErrors <- fmt.Errorf("server error: %w", err)
		}
	}()

	// 5. Graceful Shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-shutdown:
		log.Printf("received signal %v, shutting down...", sig)

		defer log.Println("server stopped cleanly ✔")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
	}

	return nil
}
