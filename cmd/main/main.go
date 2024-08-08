package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/routes"
	"github.com/thoratvinod/HashPayment/specs"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}

	err = database.InitDatabase()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer database.CloseDB()

	proto := os.Getenv("SERVER_PROTOCOL")
	if proto == "" {
		proto = "http"
	}
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	specs.ServerBaseURL = fmt.Sprintf("%s://%s:%s", proto, host, port)
	r := routes.InitRoutes()
	log.Printf("Server started on %s\n", specs.ServerBaseURL)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r))
}
