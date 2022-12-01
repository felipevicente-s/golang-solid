package main

import (
	"log"
	with3 "solid/solid/lsp/with"
	without3 "solid/solid/lsp/without"
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
	facade.Calculate(20, 10, &with2.Division{})

	//END IMPLEMENTED WITH OCP

	//START IMPLEMENTED WITHOUT LSP
	car := without3.Car{Door: 4, Vehicle: without3.Vehicle{Wheel: 4, Engine: "1.0", Color: "white", Model: "gol", Brand: "vw"}}
	horse := without3.Horse{Legs: 4, Vehicle: without3.Vehicle{Wheel: 4, Engine: "1.0", Color: "white", Model: "gol", Brand: "vw"}}

	car.Vehicle.TurnOn()
	horse.Vehicle.TurnOn()
	//END IMPLEMENTED WITHOUT LSP

	//START IMPLEMENTED WITH LSP
	carWith := with3.CarWith{Door: 4, MotorVehicle: with3.MotorVehicle{Wheel: 4, Engine: "1.0", Color: "white", Model: "gol", Brand: "vw"}}
	horseWith := with3.HorseWith{Legs: 4, PulledVehicles: with3.PulledVehicles{Name: "Pé de pano", Power: "10N"}}

	carWith.TurnOn()
	horseWith.Pull()

}
