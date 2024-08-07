package specs

var PaymentStatusToMessageMapping = map[PaymentStatus]string{
	PaymentStatusCreated:   "created",
	PaymentStatusPending:   "pending",
	PaymentStatusCanceled:  "canceled",
	PaymentStatusFailed:    "failed",
	PaymentStatusCompleted: "completed",
}
