package main

import (
	"solid/solid/ocp/with"
)

func main() {
	//START IMPLEMENTED PAYMENT WITHOUT SRP

	/*	err := without.Pay(&srp.Payment{PaymentType: "CREDIT_CARD"})

		if err != nil {
			log.Println(err.Error())
		}*/
	
	//END IMPLEMENTED PAYMENT WITHOUT SRP

	//IMPLEMENTED PAYMENT WITH SRP

	/*	factory := with.PaymentFactory{}

		payment, err := factory.Factory(&srp.Payment{PaymentType: "CREDIT_CARD"})

		if err != nil {
			log.Println(err.Error())
		} else {
			payment.Pay()
		}*/

	//END IMPLEMENTED PAYMENT WITH SRP

	//START IMPLEMENTED CALCULATOR WITHOUT OCP

	/*calculator := without2.Calculator{}
	calculator.Addition(1, 2)
	calculator.Subtraction(1, 2)*/

	//END IMPLEMENTED WITHOUT OCP

	//START IMPLEMENTED CALCULATOR WITH OCP

	facade := with.CalculatorFacade{}
	facade.InferCalc(1, 2, &with.Addition{})
}
