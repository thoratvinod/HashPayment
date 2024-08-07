package specs

type PaymentStatus uint32

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusCanceled
	PaymentStatusFailed
	PaymentStatusCompleted
)
