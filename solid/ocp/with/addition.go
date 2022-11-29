package with

import "log"

type Addition struct{}

func (ad *Addition) Perform(a int, b int) {
	log.Println(a + b)
}
