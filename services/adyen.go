package services

import (
	"context"
	"fmt"

	"github.com/adyen/adyen-go-api-library/v11/src/adyen"
	"github.com/adyen/adyen-go-api-library/v11/src/checkout"
	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

type AdyenGatewayInfo struct {
	Reference           string              `json:"reference"`
	LineItems           []checkout.LineItem `json:"lineItems"`
	MerchantAccountName string              `json:"merchantAccountName"`
}

func CreateAdyenPaymentSession(amount int64, currency, webhookURL string) (string, string, error) {

	client := adyen.NewClient(&common.Config{
		ApiKey:      "AQExhmfxK4zObxxLw0m/n3Q5qf3Vb4ZMBJ9rW2ZZ03a/zTUeL2Vi2ZEeWzsTT2G96p8q+xDBXVsNvuR83LVYjEgiTGAH-aUNKkHPjul/Z9yhWTLHDIYTkrTSI922rhW+UDuZoarM=-i1i7>$S98tZkhgw$g{F",
		Environment: common.TestEnv,
	})

	service := client.Checkout()

	body := checkout.CreateCheckoutSessionRequest{
		Reference: "test order ref",
		Mode:      common.PtrString("hosted"),
		Amount: checkout.Amount{
			Value:    amount,
			Currency: "USD",
		},
		ReturnUrl: "http://localhost:8000/webhook/success",
		// set lineItems required for some payment methods (ie Klarna)
		LineItems: []checkout.LineItem{
			{Quantity: common.PtrInt64(1), AmountIncludingTax: common.PtrInt64(5000), Description: common.PtrString("Sunglasses")},
			{Quantity: common.PtrInt64(1), AmountIncludingTax: common.PtrInt64(5000), Description: common.PtrString("Headphones")},
		},
		MerchantAccount: "GmailAccount335ECOM",
	}

	req := service.PaymentsApi.SessionsInput().CreateCheckoutSessionRequest(body)
	res, httpRes, err := service.PaymentsApi.Sessions(context.Background(), req)
	if err != nil {
		return "", "", fmt.Errorf("failed to create Adyen checkout session: err=%+v", err)
	}
	print(httpRes)
	return res.Id, *res.Url, nil
}
