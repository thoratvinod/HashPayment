package services

import (
	"fmt"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"github.com/thoratvinod/HashPayment/specs"
	// "github.com/thoratvinod/HashPayment/specs"
)

func CreateStripePaymentSession(uniqueKey string, req *specs.CreatePaymentSessionRequest) (string, string, error) {

	successURL := fmt.Sprintf(
		"%v/webhook?uniqueKey=%v&redirectURL=%v&paymentStatus=success&paymentGateway=stripe",
		specs.ServerBaseURL, uniqueKey,
		req.SuccessWebhookURL,
	)
	cancelURL := fmt.Sprintf(
		"%v/webhook?uniqueKey=%v&redirectURL=%v&paymentStatus=cancel&paymentGateway=stripe",
		specs.ServerBaseURL, uniqueKey,
		req.FailureWebhookURL,
	)

	plainAPIKey, err := GetAPIKeyManager().Get("stripe")
	if err != nil {
		return "", "", fmt.Errorf("API key for Stripe is not set, set it by calling /setapikeys API: %+v", err)
	}
	stripe.Key = plainAPIKey

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(req.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String(req.OrderName),
						Description: stripe.String(req.OrderDescription),
					},
					UnitAmount: stripe.Int64(req.Amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
		Customer:   stripe.String(req.CustomerID),
	}
	if len(req.PaymentMethodTypes) != 0 {
		params.PaymentMethodTypes = stripe.StringSlice(req.PaymentMethodTypes)
	}
	session, err := session.New(params)
	if err != nil {
		return "", "", fmt.Errorf("failed to create Stripe checkout session: %+v", err)
	}
	return session.ID, session.URL, nil
}
