package main

import (
	"log"
	"net/http"

	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	err := database.InitDatabase()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer database.Close()

	// Initialize payment services

	// Set up router
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong!"))
	})

	// Payment creation and status check endpoints
	r.HandleFunc("/create-payment", handlers.CreatePayment).Methods("POST")
	r.HandleFunc("/check-payment-status/{uniqueKey}", handlers.CheckPaymentStatus).Methods("GET")

	// Webhook endpoints for Stripe and Adyen
	// r.HandleFunc("/stripe-webhook", services.HandleStripeWebhook).Methods("POST")
	// r.HandleFunc("/adyen-webhook", services.HandleAdyenWebhook).Methods("POST")

	// Start server
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
