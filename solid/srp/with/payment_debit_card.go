package with

import (
	"log"
)

type PaymentDebitCard struct{}

func (p *PaymentDebitCard) Pay() {
	log.Println("SRP_PAYING_WITH_DEBIT_CARD")
}
