package with

import (
	"fmt"
	"log"
)

type PayrollLegalPersonImpl struct{}

func (plp *PayrollLegalPersonImpl) Pay(salary float64) {
	log.Println("SALARY PJ: " + fmt.Sprint(salary))
}
