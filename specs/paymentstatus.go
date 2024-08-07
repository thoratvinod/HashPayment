package specs

const (
	PaymentStatusCreated PaymentStatus = iota
	PaymentStatusPending
	PaymentStatusCanceled
	PaymentStatusFailed
	PaymentStatusCompleted
)
