package main

import (
	"log"
	"net/http"

	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/routes"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer database.CloseDB()

	r := routes.InitRoutes()
 	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
