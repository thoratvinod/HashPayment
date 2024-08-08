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
	// Custom data
	Metadata map[string]string `json:"metadata"`
}

type CreatePaymentSessionResponse struct {
	UniqueKey   string `json:"uniqueKey,omitempty"`
	RedirectURL string `json:"redirectURL,omitempty"`
	Error       string `json:"error,omitempty"`
}

type CheckPaymentStatusResponse struct {
	PaymentStatus string `json:"paymentStatus"`
}

type SetAPIKeysRequest struct {
	APIKeys map[string]string `json:"apiKeys"`
}

type GetPaymentDetailsResponse struct {
	PaymentGateaway  string            `json:"paymentGateaway"`
	UniqueKey        string            `json:"uniqueKey"`
	OrderName        string            `json:"orderName"`
	OrderDescription string            `json:"orderDescription"`
	Amount           int64             `json:"amount"`
	Currency         string            `json:"currency"`
	Status           PaymentStatus     `json:"status"`
	ErrorMsg         string            `json:"errorMsg,omitempty"`
	Metadata         map[string]string `json:"metadata"`
}
