package computer

import (
	calculator "go-calculator-tutorial/calculate"
)

type Computer struct {
	cal calculator.Calculator
}

func NewComputer(cal calculator.Calculator) Computer {
    return Computer {
		cal: cal,
	}
}

func (com Computer) Add3Number (a, b, c float64) float64 {
	tmp := com.cal.Add(a,b)
	return com.cal.Add(tmp,c)
}

func (com Computer) SubtractNumber (a, b float64) float64 {
	return com.cal.Subtract(a,b)
}

func (com Computer) MultiplyNumber (a, b float64) float64 {
	return com.cal.Multiply(a, b)
}

func (com Computer) DivideNumber (a, b float64) float64 {
	return com.cal.Divide(a, b)
}