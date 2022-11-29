package with

import "log"

type Subtraction struct{}

func (s *Subtraction) Calculate(a int, b int) {
	log.Println(a - b)
}
