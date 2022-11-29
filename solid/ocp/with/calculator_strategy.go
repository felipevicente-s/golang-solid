package with

type CalculatorStrategy interface {
	Perform(a int, b int)
}

type CalculatorFacade struct{}

func (c *CalculatorFacade) Calculate(a int, b int, strategy CalculatorStrategy) {
	strategy.Perform(a, b)
}
