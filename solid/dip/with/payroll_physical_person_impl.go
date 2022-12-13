package with

import (
	"fmt"
	"log"
)

type PayrollPhysicalPersonImpl struct{}

func (ppp *PayrollPhysicalPersonImpl) Pay(salary float64) {
	salary = salary - (salary * 0.275)
	log.Println("SALARY CLT: " + fmt.Sprint(salary))
}
