package main

import (
	"IOT-Smart-Agriculture/internal/config"
	"IOT-Smart-Agriculture/internal/database"
	"IOT-Smart-Agriculture/internal/router"
	"IOT-Smart-Agriculture/migration"
	"IOT-Smart-Agriculture/utils/dependency"
	"log"
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

	defer db.Close()

	if err := migration.Run(cfg.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	depInject := dependency.CreateNewDI(db)

	r := router.Setup(depInject)

	log.Printf("server started at :%s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
