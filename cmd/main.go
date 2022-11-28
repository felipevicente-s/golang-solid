package main

import (
	"log"
	"solid/solid/srp"
	"solid/solid/srp/with"
	"solid/solid/srp/without"
)

func main() {
	//START IMPLEMENTED PAYMENT WITHOUT SRP
	err := without.Pay(&srp.Payment{PaymentType: "CREDIT_CARD"})

	if err != nil {
		log.Println(err.Error())
	}
	//END IMPLEMENTED PAYMENT WITHOUT SRP

	//IMPLEMENTED PAYMENT WITH SRP
	factory := with.PaymentFactory{}

	payment, err := factory.Factory(&srp.Payment{PaymentType: "CREDIT_CARD"})

	if err != nil {
		log.Println(err.Error())
	} else {
		payment.Pay()
	}
	//END IMPLEMENTED PAYMENT WITH SRP
}
