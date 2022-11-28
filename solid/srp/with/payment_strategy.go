package with

import (
	"errors"
	"solid/solid/srp"
)

type PaymentStrategy interface {
	Pay()
}

type PaymentFactory struct {
	PayCreditCard PaymentCreditCard
	PayDebitCard  PaymentDebitCard
}

func (p *PaymentFactory) Factory(payment *srp.Payment) (PaymentStrategy, error) {
	if payment.PaymentType == "CREDIT_CARD" {
		return &p.PayCreditCard, nil
	}

	if payment.PaymentType == "DEBIT_CARD" {
		return &p.PayDebitCard, nil
	}

	return nil, errors.New("INVALID_METHOD")

}
