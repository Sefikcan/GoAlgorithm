package main

import "fmt"

// This interface contains method for processing payment transactions
type PaymentProcessor interface {
	ProcessPayment(amount float64) (string, error)
}

// Define credit card info
type CreditCard struct {
	CardNumber     string
	ExpirationDate string
	Cvv            string
	BankStrategy   BankStrategy
}

// This interface represents bank strategies
type BankStrategy interface {
	ConnectToBank() string
}

// This struct represents the default bank strategy
type DefaultBankStrategy struct{}

// This method handle to default bank strategy
func (d *DefaultBankStrategy) ConnectToBank() string {
	return "Varsayılan banka stratejisiyle bağlanılıyor..."
}

// This struct represents the bank A strategy
type BankAStrategy struct{}

// This method handle to bank A strategy
func (b *BankAStrategy) ConnectToBank() string {
	return "Bank A'ya özgü banka stratejisiyle bağlanılıyor..."
}

// set the bank strategy of the credit card
func (c *CreditCard) SetBankStrategy(strategy BankStrategy) {
	c.BankStrategy = strategy
}

// Handle credit card payment operation
func (c *CreditCard) ProcessPayment(amount float64) (string, error) {
	// connect bank strategy
	conn := c.BankStrategy.ConnectToBank()

	return conn + "\n payment successfully proceded", nil
}

func main() {
	creditCard := CreditCard{
		CardNumber:     "1234567812345678",
		ExpirationDate: "12/24",
		Cvv:            "123",
	}

	result, err := creditCard.ProcessPayment(100.0)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	fmt.Println(result)

	// handle bankA payment operation
	creditCard.SetBankStrategy(&BankAStrategy{})
	result, err = creditCard.ProcessPayment(150.0)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	fmt.Println(result)
}
