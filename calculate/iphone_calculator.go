package calculate

type IphoneCalculator struct {
}

func (c IphoneCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c IphoneCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c IphoneCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c IphoneCalculator) Divide(a, b float64) float64{

	if b == 0 {
		return 0
	} else {
		return a / b
	}
}
