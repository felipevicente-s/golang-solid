package with

import (
	"log"
)

type PaymentCreditCard struct{}

func (p *PaymentCreditCard) Pay() {
	log.Println("SRP_PAYING_WITH_CREDIT_CARD")
}
