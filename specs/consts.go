package specs

var PaymentStatusToMessageMapping = map[PaymentStatus]string{
	PaymentStatusPending:   "pending",
	PaymentStatusCanceled:  "canceled",
	PaymentStatusFailed:    "failed",
	PaymentStatusCompleted: "completed",
}

// ServerBaseURL will get set in main when spawning URL
var ServerBaseURL string
