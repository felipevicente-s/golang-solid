package with

type CalculatorStrategy interface {
	Calculate(a int, b int)
}

type CalculatorFacade struct{}

func (c *CalculatorFacade) InferCalc(a int, b int, strategy CalculatorStrategy) {
	strategy.Calculate(a, b)
}
