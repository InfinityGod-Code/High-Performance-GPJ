package main

import (
	"Intermediate/usecases"
	"fmt"
)

func ProcessPayment(processor usecases.PaymentProcessor, amount float32) {
	fmt.Println(processor.Pay(amount))
}

func main() {
	creditCard := usecases.CreditCardPayment{
		CardNumber: "5343-2443-4234-9884",
		CardOwner:  "Shubham Srivastava",
	}

	debitCard := usecases.DebitCardPayment{
		CardNumber: "9345-3434-3449-9835",
		CardOwner:  "Priyanka Chopra",
	}

	ProcessPayment(&creditCard, 56)
	ProcessPayment(&debitCard, 99)
}
