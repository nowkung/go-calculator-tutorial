package calculate

import "errors"

var result float64;

func Addition(firstNumber float64, secondNumber float64) (float64, error) {

	result = firstNumber + secondNumber;

	return result, nil;
}

func Subtraction(firstNumber float64, secondNumber float64) (float64, error) {

	result = firstNumber - secondNumber;

	return result, nil;
}

func Multiplication(firstNumber float64, secondNumber float64) (float64, error) {

	result = firstNumber * secondNumber;

	return result, nil;
}

func Division(firstNumber float64, secondNumber float64) (float64, error) {

	result = firstNumber / secondNumber;

	if (secondNumber == 0) {
        return 0, errors.New("undefined division");
    } else {
		return result, nil;
	}
}