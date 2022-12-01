package without

import "log"

type Vehicle struct {
	Wheel  int
	Engine string
	Color  string
	Model  string
	Brand  string
}

func (v Vehicle) TurnOn() {
	log.Println("Vruuumm")
}
