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
	envCfg, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	sqlDB, err := database.ConnectDB(envCfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer sqlDB.Close()

	if err := app.InitTableVotes(sqlDB); err != nil {
		log.Fatalf("Error initializing votes table: %v", err)
	}
	if err := app.InitTableRatings(sqlDB); err != nil {
		log.Fatalf("Error initializing ratings table: %v", err)
	}

	repository := app.NewProjectRepository(sqlDB)
	service := app.NewProjectService(repository)
	controller := app.NewProjectController(service)

	middlewareCfg := app.MiddlewareConfig{
		EnableLogging:      true,
		DisableStaticCache: true,
	}

	router := app.NewRouter(controller, middlewareCfg)

	server := &http.Server{
		Addr:    ":" + envCfg.ServerPort,
		Handler: router,
	}
	utils.StartServerWithGracefulShutdown(server, 5*time.Second)
}
