package specs

var PaymentStatusToMessageMapping = map[PaymentStatus]string{
	PaymentStatusPending:   "pending",
	PaymentStatusCanceled:  "canceled",
	PaymentStatusFailed:    "failed",
	PaymentStatusCompleted: "completed",
}
