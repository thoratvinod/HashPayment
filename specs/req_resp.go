package specs

type CreatePaymentSessionRequest struct {
	// PaymentGateway can be "stripe" or "adyen"
	PaymentGateaway    string   `json:"paymentGateway"`
	Amount             int64    `json:"amount"`
	Currency           string   `json:"currency"`
	SuccessWebhookURL  string   `json:"successWebhookURL"`
	FailureWebhookURL  string   `json:"failureWebhookURL"`
	OrderName          string   `json:"orderName"`
	OrderDescription   string   `json:"orderDescription"`
	PaymentMethodTypes []string `json:"paymentMethodTypes"`
	// CustomerID for stripe
	CustomerID string `json:"customerID"`
	// Adyen specific merchant account
	AdyenMerchantAccount string `json:"adyenMerchantAccount"`
}

type CreatePaymentSessionResponse struct {
	UniqueKey   string `json:"uniqueKey"`
	RedirectURL string `json:"redirectURL,omitempty"`
	Error       string `json:"error,omitempty"`
}

type CheckPaymentStatusResponse struct {
	PaymentStatus string `json:"paymentStatus"`
}

type SetAPIKeysRequest struct {
	APIKeys map[string]string `json:"apiKeys"`
}
