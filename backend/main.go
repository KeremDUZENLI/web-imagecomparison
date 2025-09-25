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

	var initQueries = []string{
		app.CreateTableSurveysQuery,
		app.CreateTableVotesQuery,
		app.CreateTableRatingsQuery,
	}
	for _, q := range initQueries {
		if _, err := sqlDB.Exec(q); err != nil {
			log.Fatalf("failed to init table: %v", err)
		}
	}

	middlewareCfg := app.MiddlewareConfig{
		EnableLogging:      true,
		DisableStaticCache: true,
	}

	repository := app.NewProjectRepository(sqlDB)
	service := app.NewProjectService(repository)
	controller := app.NewProjectController(service)
	router := app.NewRouter(controller, middlewareCfg)

	server := &http.Server{
		Addr:    ":" + envCfg.ServerPort,
		Handler: router,
	}
	utils.StartServerWithGracefulShutdown(server, 5*time.Second)
}
