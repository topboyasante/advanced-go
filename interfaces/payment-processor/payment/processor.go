package payment

type PaymentProcessor interface {
	MakeSinglePayment(amount float64, account string) error
	MakeBatchPayment(amount float64, accounts []string) error
}

