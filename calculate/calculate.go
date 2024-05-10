package calculate

type Calculator interface {
	Add(a float64, b float64) float64;
	Subtract(a float64, b float64) float64;
	Multiply(a float64, b float64) float64;
	Divide(a float64, b float64) float64;
}