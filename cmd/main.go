package main

import (
	"log"
	with2 "solid/solid/ocp/with"
	without2 "solid/solid/ocp/without"
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

	//START IMPLEMENTED CALCULATOR WITHOUT OCP

	calculator := without2.Calculator{}
	calculator.Addition(1, 2)
	calculator.Subtraction(1, 2)

	//END IMPLEMENTED WITHOUT OCP

	//START IMPLEMENTED CALCULATOR WITH OCP

	facade := with2.CalculatorFacade{}
	facade.Calculate(1, 2, &with2.Addition{})
	facade.Calculate(10, 25, &with2.Multiplication{})
}
