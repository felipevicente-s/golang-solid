package with

import "log"

type PulledVehicles struct {
	Name  string
	Power string
}

func (t PulledVehicles) Pull() {
	log.Println("Pull")
}
