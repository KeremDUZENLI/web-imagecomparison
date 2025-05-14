package main

import (
	"log"
	"net/http"

	"web-imagecomparison/app"
	"web-imagecomparison/database"
	"web-imagecomparison/env"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer db.Close()

	if err := app.InitTableVotes(db); err != nil {
		log.Fatalf("failed to init votes table: %v", err)
	}
	if err := app.InitTableRatings(db); err != nil {
		log.Fatalf("failed to init ratings table: %v", err)
	}

	repo := app.NewProjectRepository(db)
	svc := app.NewProjectService(repo)
	controller := app.NewProjectController(svc)
	router := app.NewRouter(controller)

	log.Printf("✅ http://localhost:%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
