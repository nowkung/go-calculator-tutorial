package calculate

type DigitalCalculator struct {
}

func (c DigitalCalculator) Add(a, b float64) float64 {
	return a + b + 1;
}

func (c DigitalCalculator) Subtract(a, b float64) float64 {
	return a - b + 1;
}

func (c DigitalCalculator) Multiply(a, b float64) float64 {
	return a * b + 1;
}

func (c DigitalCalculator) Divide(a, b float64) float64 {

	if (b == 0) {
        return 0;
    } else {
		return a/b + 1;
	}
}