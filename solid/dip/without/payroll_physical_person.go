package without

import (
	"fmt"
	"log"
)

type PayrollPhysicalPerson struct{}

func (plp *PayrollPhysicalPerson) Pay(salary float64) {
	salary = salary - (salary * 0.275)
	log.Println("SALARY CLT: " + fmt.Sprint(salary))
}
