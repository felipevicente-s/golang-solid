package without

import (
	"errors"
	"log"
	"solid/solid/srp"
)

func Pay(payment *srp.Payment) error {
	if payment.PaymentType == "CREDIT_CARD" {
		log.Println("NO_SRP_PAYING_WITH_" + payment.PaymentType)
	} else if payment.PaymentType == "DEBIT_CARD" {
		log.Println("NO_SRP_PAYING_WITH_" + payment.PaymentType)
	} else {
		return errors.New("INVALID_METHOD")
	}
	return nil
}
