package main

import "fmt"

// Define an interface for a payment method
type PaymentMethod interface {
	MakePayment(amount float64) error
	GetPaymentDetails() string
}

// Define a type for a credit card
type CreditCard struct {
	number string
	expiry string
}

// Implement the PaymentMethod interface for CreditCard
func (c CreditCard) MakePayment(amount float64) error {
	// Simulate a payment being made
	fmt.Println("Charging $", amount, " to credit card", c.number)
	return nil
}

func (c CreditCard) GetPaymentDetails() string {
	return "Credit Card: " + c.number
}

// Define a type for PayPal
type PayPal struct {
	email string
}

// Implement the PaymentMethod interface for PayPal
func (p PayPal) MakePayment(amount float64) error {
	// Simulate a payment being made
	fmt.Println("Sending $", amount, " to PayPal account", p.email)
	return nil
}

func (p PayPal) GetPaymentDetails() string {
	return "PayPal: " + p.email
}

// Use the PaymentMethod interface as a type
func processPayment(p PaymentMethod, amount float64) {
	err := p.MakePayment(amount)
	if err != nil {
		fmt.Println("Error making payment:", err)
	} else {
		fmt.Println("Payment successful. Details:", p.GetPaymentDetails())
	}
}

func main() {
	creditCard := CreditCard{number: "1234-5678-9012-3456"}
	processPayment(creditCard, 100)

	payPal := PayPal{email: "example@example.com"}
	processPayment(payPal, 50)
}
