package with

import "log"

type Multiplication struct {
}

func (m *Multiplication) Perform(a int, b int) {
	log.Println(a * b)
}
