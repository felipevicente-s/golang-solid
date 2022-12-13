package without

import (
	"fmt"
	"log"
)

type PayrollLegalPerson struct{}

func (plp *PayrollLegalPerson) Pay(salary float64) {
	log.Println("SALARY PJ: " + fmt.Sprint(salary))
}
