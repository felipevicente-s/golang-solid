package main

import (
	"log"
	with6 "solid/solid/dip/with"
	without5 "solid/solid/dip/without"
	"solid/solid/isp"
	with4 "solid/solid/isp/with"
	without4 "solid/solid/isp/without"
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
	//END IMPLEMENTED WITH LSP

	//START IMPLMENTED WITHOUT ISP
	order := isp.Order{
		Status: "PENDING",
	}

	approved := without4.OrderApproved{}
	approved.Next(&order)

	canceled := without4.OrderCanceled{}
	canceled.Next(&order)

	log.Println(order.Status)
	//END IMPLEMENTED WITHOUT ISP

	//START IMPLEMENTED WITH ISP
	approvedWith := with4.OrderApprovedWith{}
	approvedWith.Next(&order)
	log.Println(order.Status)
	approvedWith.Prev(&order)
	log.Println(order.Status)

	failedWith := with4.OrderFailedWith{}
	failedWith.Cancel(&order)
	log.Println(order.Status)

	failedWith.Prev(&order)
	log.Println(order.Status)
	//END IMPLEMENTED WITH ISP

	//START IMPLEMENTED WITHOUT DIP
	person := without5.PayrollLegalPerson{}
	person.Pay(5000.00)

	physicalPerson := without5.PayrollPhysicalPerson{}
	physicalPerson.Pay(5000.00)
	//END IMPLEMENTED WITHOUT DIP

	//START IMPLEMENTED WITH DIP
	payrollFacade := with6.PayrollFacade{}
	payrollFacade.Infer(&with6.PayrollLegalPersonImpl{}, 5000.00)

	payrollFacade.Infer(&with6.PayrollPhysicalPersonImpl{}, 5000.00)
}
