package main

import (
	"fmt"
	"go-calculator-tutorial/calculate"
	"strings"
)

func main() {

	type calculator struct{
		firstNumber float64;
        secondNumber float64;
        mathOperation string;
	}

	c := calculator{firstNumber: 0, secondNumber: 0, mathOperation:""};

	fmt.Println("First Number:")
	fmt.Scan(&c.firstNumber);
	fmt.Println("Second Number:")
	fmt.Scan(&c.secondNumber);
	fmt.Println("Math Operation:")
	fmt.Scan(&c.mathOperation)

	switch strings.ToLower(c.mathOperation) {
		case "addition":
			ans, _ := calculate.Addition(c.firstNumber, c.secondNumber)
			fmt.Printf("Answer: ")
			fmt.Println(ans)
		case "subtraction":
			ans, _ := calculate.Subtraction(c.firstNumber, c.secondNumber)
			fmt.Printf("Answer: ")
			fmt.Println(ans)
		case "multiplication":
			ans, _ := calculate.Multiplication(c.firstNumber, c.secondNumber)
			fmt.Printf("Answer: ")
			fmt.Println(ans)
		case "division":
			ans, _ := calculate.Division(c.firstNumber, c.secondNumber)
			fmt.Printf("Answer: ")
			fmt.Println(ans)
		default:
			fmt.Println("Invalid Operation")
	}

}