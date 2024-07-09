package main

import (
	"fmt"

	"github.com/topboyasante/advanced-go/payment-processor/payment"
)

func MakeSinglePayment(processor payment.PaymentProcessor, amount float64, account string) {
	err := processor.MakeSinglePayment(amount, account)
	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}
func MakeBatchPayment(processor payment.PaymentProcessor, amount float64, accounts []string) {
	err := processor.MakeBatchPayment(amount, accounts)
	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}

func main() {
	paypal := &payment.Paypal{}

	MakeSinglePayment(paypal, 500, "Adolf")
	MakeBatchPayment(paypal, 500, []string{"Adolf", "Bruce","Vector"})
}
