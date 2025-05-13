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

	application, err := app.RunApp(db)
	if err != nil {
		log.Fatal(err)
	}
	application.Routes()

	log.Printf("✅ http://localhost:%s", env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+env.SERVER_PORT, nil))
}
