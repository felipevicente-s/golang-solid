package with

import "log"

type MotorVehicle struct {
	Wheel  int
	Engine string
	Color  string
	Model  string
	Brand  string
}

func (m MotorVehicle) TurnOn() {
	log.Println("Vruuumm")
}
