package main

import (
	"IOT-Smart-Agriculture/internal/config"
	"IOT-Smart-Agriculture/internal/database"
	"IOT-Smart-Agriculture/internal/router"
	"IOT-Smart-Agriculture/migration"
	"IOT-Smart-Agriculture/utils/dependency"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration loaded")

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")

	defer func() {
		db.Close()
		log.Println("Database connection closed")
	}()

	if err := migration.Run(cfg.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	depInject := dependency.CreateNewDI(db, *cfg)

	r := router.Setup(depInject)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server started at :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
