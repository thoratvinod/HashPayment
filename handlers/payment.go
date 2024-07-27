package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/models"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment.UniqueKey = uuid.New()
	payment.Status = models.PaymentStatusNone

	database.DB.Create(&payment)

	// TODO payment integraiton logic here

	payment.RedirectURL = "redirect-url"
	database.DB.Save(&payment)

	json.NewEncoder(w).Encode(payment)
}

func CheckPaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uniqueKey := vars["uniqueKey"]

	var payment models.Payment
	if err := database.DB.Where("unique_key = ?", uniqueKey).First(&payment).Error; err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}
