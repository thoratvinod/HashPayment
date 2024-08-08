package services

import (
	"context"
	"fmt"

	"github.com/adyen/adyen-go-api-library/v11/src/adyen"
	"github.com/adyen/adyen-go-api-library/v11/src/checkout"
	"github.com/adyen/adyen-go-api-library/v11/src/common"
	"github.com/thoratvinod/HashPayment/specs"
)

func CreateAdyenPaymentSession(uniqueKey string, req *specs.CreatePaymentSessionRequest) (string, string, error) {
	returnURL := fmt.Sprintf(
		"%v/webhook?uniqueKey=%v&paymentGateway=adyen&successWebhookURL=%v&failureWebhookURL=%v",
		specs.ServerBaseURL, uniqueKey,
		req.SuccessWebhookURL,
		req.FailureWebhookURL,
	)

	plainAPIKey, err := GetAPIKeyManager().Get("adyen")
	if err != nil {
		return "", "", fmt.Errorf("API key for Adyen is not set, set it by calling /setapikeys API: %+v", err)
	}

	client := adyen.NewClient(&common.Config{
		ApiKey:      string(plainAPIKey),
		Environment: common.TestEnv,
	})

	service := client.Checkout()

	body := checkout.CreateCheckoutSessionRequest{
		Reference: req.OrderName,
		Mode:      common.PtrString("hosted"),
		Amount: checkout.Amount{
			Value:    req.Amount,
			Currency: req.Currency,
		},
		ReturnUrl: returnURL,
		// set lineItems required for some payment methods (ie Klarna)
		LineItems: []checkout.LineItem{
			{Quantity: common.PtrInt64(1), AmountIncludingTax: common.PtrInt64(req.Amount), Description: common.PtrString(req.OrderDescription)},
		},
		MerchantAccount: req.AdyenMerchantAccount,
	}

	if len(req.PaymentMethodTypes) == 0 {
		body.AllowedPaymentMethods = req.PaymentMethodTypes
	}

	sessionReq := service.PaymentsApi.SessionsInput().CreateCheckoutSessionRequest(body)
	res, _, err := service.PaymentsApi.Sessions(context.Background(), sessionReq)
	if err != nil {
		return "", "", fmt.Errorf("failed to create Adyen checkout session: err=%+v", err)
	}
	return res.Id, *res.Url, nil
}

func GetAdyenPaymentStatus(sessionId, sessionResult string) {

}
