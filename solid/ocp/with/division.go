package with

import "log"

type Division struct {
}

func (d *Division) Perform(a int, b int) {
	log.Println(a / b)
}
