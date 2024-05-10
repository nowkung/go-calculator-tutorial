package main

import (
	"fmt"
	calculator "go-calculator-tutorial/calculate"
	computer "go-calculator-tutorial/computer"
)

func main() {

	iphoneCalculator := calculator.IphoneCalculator{}
	digitalCalculator := calculator.DigitalCalculator{}

	com1 := computer.NewComputer(iphoneCalculator)
	com2 := computer.NewComputer(digitalCalculator)

	result1 := com1.MultiplyNumber(2, 4)
	result2 := com2.MultiplyNumber(2, 4)

	fmt.Println(result1, result2)

}