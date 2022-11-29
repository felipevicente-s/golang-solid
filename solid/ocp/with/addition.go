package with

import "log"

type Addition struct{}

func (ad *Addition) Calculate(a int, b int) {
	log.Println(a + b)
}
