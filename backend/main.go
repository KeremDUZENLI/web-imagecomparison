package main

import (
	"log"
	"net/http"

	"web-imagecomparison/common/env"
	"web-imagecomparison/controller"
	"web-imagecomparison/database"
	"web-imagecomparison/repository"
	"web-imagecomparison/service"
)

func main() {
	env.Load()

	db, err := database.DatabaseDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := repository.InitTable(db); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewProjectRepository(db)
	svc := service.NewProjectService(repo)
	ctrl := controller.NewProjectController(svc)

	http.Handle("/", http.FileServer(http.Dir("../docs")))
	http.HandleFunc("/api/votes", ctrl.HandleEntry)

	log.Println("✅ Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
