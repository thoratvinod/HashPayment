package handlers

import (
	"context"
	"fmt"
	"net/http"

	// "github.com/adyen/adyen-go-api-library/v11/src/adyen/models/checkout"

	"github.com/adyen/adyen-go-api-library/v11/src/adyen"
	"github.com/adyen/adyen-go-api-library/v11/src/common"

	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/services"
	"github.com/thoratvinod/HashPayment/specs"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	paymentGateway := r.URL.Query().Get("paymentGateway")
	// if paymentGateway == "" {

	// }

	switch paymentGateway {
	case "stripe":
		handleStripeWebhook(w, r)
	case "adyen":
		handleAdyenWebhook(w, r)
	default:
	}

}

func handleStripeWebhook(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	uniqueKey := queryParams.Get("uniqueKey")
	redirectURL := queryParams.Get("redirectURL")
	paymentStatus := queryParams.Get("paymentStatus")

	var paymentStatusEnum specs.PaymentStatus
	switch paymentStatus {
	case "success":
		paymentStatusEnum = specs.PaymentStatusCompleted
	case "cancel":
		paymentStatusEnum = specs.PaymentStatusCanceled
	default:
		// handle
	}
	updatePaymentStatus(uniqueKey, paymentStatusEnum)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func handleAdyenWebhook(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	uniqueKey := queryParams.Get("uniqueKey")
	successWebhookURL := queryParams.Get("successWebhookURL")
	failureWebhookURL := queryParams.Get("failureWebhookURL")
	sessionId := queryParams.Get("sessionId")
	sessionResult := queryParams.Get("sessionResult")

	plainAPIKey, err := services.GetAPIKeyManager().Get("adyen")
	if err != nil {
		httpErr := fmt.Errorf("API key for Adyen is not set, set it by calling /setapikeys API: %+v", err)
		http.Error(w, httpErr.Error(), http.StatusBadRequest)
		return
	}

	client := adyen.NewClient(&common.Config{
		ApiKey:      string(plainAPIKey),
		Environment: common.TestEnv,
	})

	service := client.Checkout()

	getResultPaymentSession := service.PaymentsApi.GetResultOfPaymentSessionInput(sessionId)
	getResultPaymentSession = getResultPaymentSession.SessionResult(sessionResult)
	res, _, err := client.Checkout().PaymentsApi.GetResultOfPaymentSession(context.Background(), getResultPaymentSession)
	if err != nil {
		//handle
	}
	var statusToUpdate specs.PaymentStatus
	var redirectURL string
	switch *res.Status {
	case "completed":
		redirectURL = successWebhookURL
		statusToUpdate = specs.PaymentStatusCompleted
	case "paymentPending":
		redirectURL = failureWebhookURL
		statusToUpdate = specs.PaymentStatusPending
	case "refused", "canceled":
		redirectURL = failureWebhookURL
		statusToUpdate = specs.PaymentStatusCanceled
	case "expired":
		redirectURL = failureWebhookURL
		statusToUpdate = specs.PaymentStatusFailed
	}
	updatePaymentStatus(uniqueKey, statusToUpdate)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func updatePaymentStatus(uniqueKey string, status specs.PaymentStatus) error {
	var payment specs.PaymentModel
	if err := database.DB.Where("unique_key = ?", uniqueKey).First(&payment).Error; err != nil {
		return err
	}
	payment.Status = status
	return database.DB.Save(&payment).Error
}
