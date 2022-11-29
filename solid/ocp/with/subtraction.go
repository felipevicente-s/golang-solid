package with

import "log"

type Subtraction struct{}

func (s *Subtraction) Perform(a int, b int) {
	log.Println(a - b)
}
