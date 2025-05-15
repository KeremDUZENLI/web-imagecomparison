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
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer db.Close()

	if err := app.InitTableVotes(db); err != nil {
		log.Fatalf("Error initializing votes table: %v", err)
	}
	if err := app.InitTableRatings(db); err != nil {
		log.Fatalf("Error initializing ratings table: %v", err)
	}

	repo := app.NewProjectRepository(db)
	svc := app.NewProjectService(repo)
	ctrl := app.NewProjectController(svc)

	logCfg := app.MiddlewareConfig{EnableLogging: true}
	router := app.NewRouter(ctrl, logCfg)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}
	utils.StartServerWithGracefulShutdown(srv, 5*time.Second)
}
