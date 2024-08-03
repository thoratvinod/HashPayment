package services

import (
	"fmt"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"github.com/thoratvinod/HashPayment/configs"
	"github.com/thoratvinod/HashPayment/specs"
	// "github.com/thoratvinod/HashPayment/specs"
)

const (
	baseURL = "http://localhost:8000"
)

type StripeTenantInfo struct {
	CustomerID string                                  `json:"customerID"`
	LineItems  []*stripe.CheckoutSessionLineItemParams `json:"lineItems"`
}

// func ProcessPayment(amount int64, metadata map[string]string, currency, customerID, description, secretKey string) (*stripe.PaymentIntent, error) {

// 	stripe.Key = secretKey

// 	paymentIntentParams := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(int64(amount * 100)), // Stripe amounts are in cents
// 		Currency: stripe.String(currency),
// 		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
// 			Enabled: stripe.Bool(true),
// 		},
// 	}

// 	stripe.CheckoutSessionParams

// 	// paymentIntentParams := &stripe.PaymentIntentParams{
// 	// 	Amount:   stripe.Int64(1000), // Amount in the smallest currency unit (e.g., cents for USD)
// 	// 	Currency: stripe.String(string(stripe.CurrencyUSD)),
// 	// 	Customer: stripe.String(customerID),
// 	// 	AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
// 	// 		Enabled: stripe.Bool(true),
// 	// 	},
// 	// 	PaymentMethod: ,
// 	// 	// PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
// 	// 	// 	Type: stripe.String("card"),
// 	// 	// 	Card: &stripe.PaymentMethodCardParams{
// 	// 	// 		Number:   stripe.String("4242424242424242"),
// 	// 	// 		ExpMonth: stripe.String("12"),
// 	// 	// 		ExpYear:  stripe.String("2025"),
// 	// 	// 		CVC:      stripe.String("123"),
// 	// 	// 	},
// 	// 	// },
// 	// 	// Confirm: stripe.Bool(true),
// 	// }

// 	paymentIntent, err := paymentintent.New(paymentIntentParams)
// 	if err != nil {
// 		log.Fatalf("Failed to create PaymentIntent: %v", err)
// 	}

// 	fmt.Printf("PaymentIntent created: %s\n", paymentIntent.ID)

// 	return nil, nil

// 	// paymentIntent, err := sc.PaymentIntents.New(paymentIntentParams)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to create PaymentIntent: %v", err)
// 	// }

// 	// fmt.Printf("PaymentIntent created: %s\n", paymentIntent.ID)

// 	// params := &stripe.PaymentIntentParams{
// 	// 	Amount:   stripe.Int64(amount),
// 	// 	Currency: stripe.String(currency),
// 	// 	Customer: stripe.String(customerID),
// 	// 	// PaymentMethodTypes: stripe.StringSlice([]string{"cash_balance"}),
// 	// 	// Confirm:   stripe.Bool(true),
// 	// 	// ReturnURL: stripe.String("http://localhost:8000/handle-next-action"),
// 	// }

// 	// params := &stripe.PaymentIntentParams{
// 	// 	Amount:   stripe.Int64(1000), // amount in cents
// 	// 	Currency: stripe.String(string(stripe.CurrencyUSD)),
// 	// 	PaymentMethodTypes: stripe.StringSlice([]string{
// 	// 		"customer_balance",
// 	// 	}),
// 	// 	PaymentMethod: stripe.String("customer_balance"),
// 	// 	PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
// 	// 		CustomerBalance: &stripe.PaymentIntentPaymentMethodOptionsCustomerBalanceParams{
// 	// 			FundingType: stripe.String("bank_transfer"),
// 	// 		},
// 	// 	},
// 	// 	Confirm:   stripe.Bool(true), // Automatically confirm the payment intent
// 	// 	ReturnURL: stripe.String("http://localhost:8000/handle-next-action"),
// 	// }

// 	// for key, value := range metadata {
// 	// 	params.AddMetadata(key, value)
// 	// }

// 	// pi, err := paymentintent.New(params)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("failed to create charge: %w", err)
// 	// }
// 	// return pi, nil
// }

// func extractStripeGatewayInfo(request *specs.ProcessPaymentRequest, )

// func CreateStripePaymentSession(request *specs.ProcessPaymentRequest) (string, string, error) {

// }

// func decodeStripeTenantInfo(tenantInfoBytes []byte) (*StripeTenantInfo, error) {
// 	tenantInfo := StripeTenantInfo{}
// 	err := json.Unmarshal(tenantInfoBytes, &tenantInfo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &tenantInfo, err
// }

func CreateStripePaymentSession(uniqueKey string, req *specs.CreatePaymentSessionRequest) (string, string, error) {

	successURL := fmt.Sprintf(
		"%v/webhook/success?uniqueKey=%v&redirectURL=%v",
		baseURL, uniqueKey,
		req.SuccessWebhookURL,
	)

	cancelURL := fmt.Sprintf(
		"%v/webhook/cancel?uniqueKey=%v&redirectURL=%v",
		baseURL, uniqueKey,
		req.FailureWebhookURL,
	)

	stripe.Key = configs.SecretTestKey
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice(req.PaymentMethodTypes),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(req.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String(req.OrderName),
						Description: stripe.String(req.OrderDescription),
					},
					UnitAmount: stripe.Int64(req.Amount), // Stripe amounts are in cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
		Customer:   stripe.String(req.CustomerID),
	}
	session, err := session.New(params)
	if err != nil {
		return "", "", fmt.Errorf("failed to create Stripe checkout session: %+v", err)
	}
	return session.ID, session.URL, nil
}
