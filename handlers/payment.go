package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/services"
	"github.com/thoratvinod/HashPayment/specs"
)

func CreatePaymentSession(w http.ResponseWriter, r *http.Request) {
	var paymentRequest specs.CreatePaymentSessionRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paymentModel := specs.PaymentModel{}
	paymentModel.UniqueKey = uuid.New()
	paymentModel.Status = specs.PaymentStatusCreated


	var returnURL string

	// TODO validation

	switch paymentRequest.PaymentGateaway {
	case "stripe":
		_, returnURL, err = services.CreateStripePaymentSession(
			paymentModel.UniqueKey.String(),
			&paymentRequest,
		)
	case "adyen":
		_, returnURL, err = services.CreateAdyenPaymentSession(
			paymentModel.UniqueKey.String(),
			&paymentRequest,
		)
	default:
		http.Error(w, "Invalid payment gateway provided", http.StatusBadRequest)
		return
	}

	var resp specs.CreatePaymentSessionResponse

	if err != nil {
		paymentModel.ErrorMsg = err.Error()
		resp = specs.CreatePaymentSessionResponse{
			UniqueKey: paymentModel.UniqueKey.String(),
			Error:     err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp = specs.CreatePaymentSessionResponse{
			UniqueKey:   paymentModel.UniqueKey.String(),
			RedirectURL: returnURL,
		}
	}
	database.DB.Create(&paymentModel)
	json.NewEncoder(w).Encode(resp)
}

func CheckPaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uniqueKey := vars["uniqueKey"]

	// TODO validation

	var payment specs.PaymentModel
	if err := database.DB.Where("unique_key = ?", uniqueKey).First(&payment).Error; err != nil {
		http.Error(w, "Payment record not found", http.StatusBadRequest)
		return
	}
	resp := specs.CheckPaymentStatusResponse{
		PaymentStatus: payment.Status,
	}

	json.NewEncoder(w).Encode(resp)
}
