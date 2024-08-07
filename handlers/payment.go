package handlers

import (
	"encoding/json"
	"fmt"
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

	_, err = services.GetAPIKeyManager().Get(paymentRequest.PaymentGateaway)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paymentModel := specs.PaymentModel{
		PaymentGateaway:  paymentRequest.PaymentGateaway,
		UniqueKey:        uuid.New(),
		OrderName:        paymentRequest.OrderName,
		OrderDescription: paymentRequest.OrderDescription,
		Amount:           paymentRequest.Amount,
		Currency:         paymentRequest.Currency,
	}

	if len(paymentRequest.Metadata) != 0 {
		metadataJson, err := json.Marshal(paymentRequest.Metadata)
		if err != nil {
			http.Error(w, fmt.Sprintf("error while marshaling metadata: %+v", err), http.StatusBadRequest)
			return
		}
		paymentModel.Metadata = string(metadataJson)
	}

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
		paymentModel.Status = specs.PaymentStatusFailed
		resp = specs.CreatePaymentSessionResponse{
			UniqueKey: paymentModel.UniqueKey.String(),
			Error:     err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		paymentModel.Status = specs.PaymentStatusPending
		resp = specs.CreatePaymentSessionResponse{
			UniqueKey:   paymentModel.UniqueKey.String(),
			RedirectURL: returnURL,
		}
	}
	reqJson, err := json.Marshal(paymentRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while marshaling request: %+v", err), http.StatusBadRequest)
		return
	}
	paymentModel.RawRequest = string(reqJson)
	result := database.DB.Create(&paymentModel)
	if result.Error != nil {
		resp = specs.CreatePaymentSessionResponse{
			Error: fmt.Sprintf("error creating payment in DB: %+v", result.Error),
		}
	}
	json.NewEncoder(w).Encode(resp)
}

func GetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uniqueKey := vars["uniqueKey"]

	// TODO validation

	var payment specs.PaymentModel
	if err := database.DB.Where("unique_key = ?", uniqueKey).First(&payment).Error; err != nil {
		http.Error(w, "Payment record not found", http.StatusBadRequest)
		return
	}
	metadata := make(map[string]string)
	if payment.Metadata != "" {
		err := json.Unmarshal([]byte(payment.Metadata), &metadata)
		if err != nil {
			metadata["error-in-unmarshal"] = err.Error()
			metadata["rawMetadata"] = payment.Metadata
		}
	}
	resp := specs.GetPaymentDetailsResponse{
		PaymentGateaway:  payment.PaymentGateaway,
		UniqueKey:        uniqueKey,
		OrderName:        payment.OrderName,
		OrderDescription: payment.OrderDescription,
		Amount:           payment.Amount,
		Currency:         payment.Currency,
		Status:           payment.Status,
		ErrorMsg:         payment.ErrorMsg,
		Metadata:         metadata,
	}

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

	statusMsg := specs.PaymentStatusToMessageMapping[payment.Status]
	resp := specs.CheckPaymentStatusResponse{
		PaymentStatus: statusMsg,
	}

	json.NewEncoder(w).Encode(resp)
}
