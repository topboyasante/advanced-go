package payment

import "fmt"

type Paypal struct{}

func (p *Paypal) MakeSinglePayment(amount float64, account string) error {
	fmt.Printf("Processing PayPal payment of GHC %v to %v \n", amount, account)
	return nil
}

func (p *Paypal) MakeBatchPayment(amount float64, accounts []string) error {
	for i := 0; i < len(accounts); i++ {
		fmt.Printf("Processing PayPal payment of GHC %v to %v \n", amount, accounts[i])
	}
	return nil
}
