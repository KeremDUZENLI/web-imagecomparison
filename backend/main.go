package main

import (
	"log"
	"net/http"
	"web-imagecomparison/app"
	"web-imagecomparison/database"
	"web-imagecomparison/env"
)

func main() {
	env.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	application, err := app.InitApp(db)
	if err != nil {
		log.Fatal(err)
	}

	router := app.NewRouter(application.Service)
	log.Printf("✅ http://localhost:%s", env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+env.SERVER_PORT, router))
}
