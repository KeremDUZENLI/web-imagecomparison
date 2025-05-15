package main

import (
	"log"
	"net/http"
	"time"

	"web-imagecomparison/app"
	"web-imagecomparison/database"
	"web-imagecomparison/env"
	"web-imagecomparison/utils"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("\t\u274c Error loading config: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("\t\u274c Error connecting to database: %v", err)
	}
	defer db.Close()

	if err := app.InitTableVotes(db); err != nil {
		log.Fatalf("\t\u274c Error initializing votes table: %v", err)
	}
	if err := app.InitTableRatings(db); err != nil {
		log.Fatalf("\t\u274c Error initializing ratings table: %v", err)
	}

	repo := app.NewProjectRepository(db)
	service := app.NewProjectService(repo)
	controller := app.NewProjectController(service)
	router := app.NewRouter(controller)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	utils.StartServerWithGracefulShutdown(server, 5*time.Second)
}
