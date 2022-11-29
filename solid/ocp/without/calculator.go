package without

import "log"

type Calculator struct{}

func (c *Calculator) Addition(a int, b int) {
	log.Println(a + b)
}

func (c *Calculator) Subtraction(a int, b int) {
	log.Println(a - b)
}
