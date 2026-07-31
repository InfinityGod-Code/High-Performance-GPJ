package usecases

import (
	"fmt"
)

// PaymentProcessor represents an entity that can4 process payments.
// It should be an interface so implementations can provide the Pay method.
type PaymentProcessor interface {
	Pay(amount float32) string
}

// Go follows the Duck Typing in order to implement the Interface. If it walks like a Duck and quack like
// a Duct, it should be the duck.

type DebitCardPayment struct {
	CardNumber string
	CardOwner  string
}

// since this method ressembles the signature of the interface therefore, DebitCardPayment implements
// the PaymentProcessor.
func (py *DebitCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("%s with %s has paid %f", py.CardOwner, py.CardNumber, amount)
}

type CreditCardPayment struct {
	CardNumber string
	CardOwner  string
}

func (ct *CreditCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("%s with %s has paid %f", ct.CardOwner, ct.CardNumber, amount)
}
