package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition1(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := 4.00

	ans, err := Addition(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, 7.00, ans)
	}
	
}

func TestAddition2(t *testing.T) {
	firstNumber  := -3.00
	secondNumber := 4.00

	ans, err := Addition(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, 1.00, ans)
	}
	
}

func TestSubtraction1(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := 4.00

	ans, err := Subtraction(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, -1.00, ans)
	}
}

func TestSubtraction2(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := -4.00

	ans, err := Subtraction(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, 7.00, ans)
	}
}

func TestMultiplication1(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := 4.00

	ans, err := Multiplication(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, 12.00, ans)
	}
}

func TestMultiplication2(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := -4.00

	ans, err := Multiplication(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, -12.00, ans)
	}
}

func TestDivision1(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := 4.00

	ans, err := Division(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err)
    } else {
		assert.Equal(t, 0.75, ans)
	}
}

func TestDivision1Error(t *testing.T) {
	firstNumber  := 3.00
	secondNumber := 0.00

	ans, err := Division(firstNumber, secondNumber)

	if err != nil {
        assert.Error(t, err, "undefined division")
    } else {
		assert.Equal(t, -1.00, ans)
	}
}
