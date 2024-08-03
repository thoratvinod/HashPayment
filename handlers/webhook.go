package handlers

import (
	"net/http"

	"github.com/thoratvinod/HashPayment/database"
	"github.com/thoratvinod/HashPayment/specs"
)

func SuccessWebhook(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	uniqueKey := params.Get("uniqueKey")
	redirectURL := params.Get("redirectURL")

	updatePaymentStatus(uniqueKey, specs.PaymentStatusSuccess)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func CanceledWebhook(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uniqueKey := params.Get("uniqueKey")
	redirectURL := params.Get("redirectURL")

	updatePaymentStatus(uniqueKey, specs.PaymentStatusCanceled)
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
