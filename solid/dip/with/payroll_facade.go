package with

type PayrollFacade struct{}

func (pf *PayrollFacade) Infer(strategy Payroll, salary float64) {
	strategy.Pay(salary)
}
