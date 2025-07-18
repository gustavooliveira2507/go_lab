package math

import (
	"testing"
)

func TestSumPositiveNumbers(t *testing.T) {
	result := Sum(10.5, 4.5)
	expected := 15.0
	if result != expected {
		t.Errorf("Sum(10.5, 4.5): expected %f, got %f", expected, result)
	}
}

func TestSumNegativeNumbers(t *testing.T) {
	result := Sum(-2.0, -3.0)
	expected := -5.0
	if result != expected {
		t.Errorf("Sum(-2.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestSumPositiveAndNegative(t *testing.T) {
	result := Sum(5.0, -3.0)
	expected := 2.0
	if result != expected {
		t.Errorf("Sum(5.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestSumWithZero(t *testing.T) {
	result := Sum(0.0, 7.0)
	expected := 7.0
	if result != expected {
		t.Errorf("Sum(0.0, 7.0): expected %f, got %f", expected, result)
	}
}
func TestSubtractPositiveNumbers(t *testing.T) {
	result := Subtract(10.5, 4.5)
	expected := 6.0
	if result != expected {
		t.Errorf("Subtract(10.5, 4.5): expected %f, got %f", expected, result)
	}
}

func TestSubtractNegativeNumbers(t *testing.T) {
	result := Subtract(-2.0, -3.0)
	expected := 1.0
	if result != expected {
		t.Errorf("Subtract(-2.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestSubtractPositiveAndNegative(t *testing.T) {
	result := Subtract(5.0, -3.0)
	expected := 8.0
	if result != expected {
		t.Errorf("Subtract(5.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestSubtractWithZero(t *testing.T) {
	result := Subtract(0.0, 7.0)
	expected := -7.0
	if result != expected {
		t.Errorf("Subtract(0.0, 7.0): expected %f, got %f", expected, result)
	}
}
func TestMultiplyPositiveNumbers(t *testing.T) {
	result := Multiply(3.0, 4.0)
	expected := 12.0
	if result != expected {
		t.Errorf("Multiply(3.0, 4.0): expected %f, got %f", expected, result)
	}
}

func TestMultiplyNegativeNumbers(t *testing.T) {
	result := Multiply(-2.0, -5.0)
	expected := 10.0
	if result != expected {
		t.Errorf("Multiply(-2.0, -5.0): expected %f, got %f", expected, result)
	}
}

func TestMultiplyPositiveAndNegative(t *testing.T) {
	result := Multiply(6.0, -3.0)
	expected := -18.0
	if result != expected {
		t.Errorf("Multiply(6.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestMultiplyWithZero(t *testing.T) {
	result := Multiply(0.0, 7.0)
	expected := 0.0
	if result != expected {
		t.Errorf("Multiply(0.0, 7.0): expected %f, got %f", expected, result)
	}
	result = Multiply(5.0, 0.0)
	if result != 0.0 {
		t.Errorf("Multiply(5.0, 0.0): expected 0.0, got %f", result)
	}
}

func TestMultiplyWithOne(t *testing.T) {
	result := Multiply(1.0, 9.0)
	expected := 9.0
	if result != expected {
		t.Errorf("Multiply(1.0, 9.0): expected %f, got %f", expected, result)
	}
	result = Multiply(9.0, 1.0)
	if result != expected {
		t.Errorf("Multiply(9.0, 1.0): expected %f, got %f", expected, result)
	}
}
func TestDividePositiveNumbers(t *testing.T) {
	result, err := Divide(10.0, 2.0)
	expected := 5.0
	if err != nil {
		t.Errorf("Divide(10.0, 2.0): unexpected error %v", err)
	}
	if result != expected {
		t.Errorf("Divide(10.0, 2.0): expected %f, got %f", expected, result)
	}
}

func TestDivideNegativeNumbers(t *testing.T) {
	result, err := Divide(-9.0, -3.0)
	expected := 3.0
	if err != nil {
		t.Errorf("Divide(-9.0, -3.0): unexpected error %v", err)
	}
	if result != expected {
		t.Errorf("Divide(-9.0, -3.0): expected %f, got %f", expected, result)
	}
}

func TestDividePositiveAndNegative(t *testing.T) {
	result, err := Divide(8.0, -2.0)
	expected := -4.0
	if err != nil {
		t.Errorf("Divide(8.0, -2.0): unexpected error %v", err)
	}
	if result != expected {
		t.Errorf("Divide(8.0, -2.0): expected %f, got %f", expected, result)
	}
}

func TestDivideByZero(t *testing.T) {
	result, err := Divide(5.0, 0.0)
	if err == nil {
		t.Errorf("Divide(5.0, 0.0): expected error, got nil")
	}
	if result != 0.0 {
		t.Errorf("Divide(5.0, 0.0): expected result 0.0, got %f", result)
	}
}

func TestDivideZeroNumerator(t *testing.T) {
	result, err := Divide(0.0, 7.0)
	expected := 0.0
	if err != nil {
		t.Errorf("Divide(0.0, 7.0): unexpected error %v", err)
	}
	if result != expected {
		t.Errorf("Divide(0.0, 7.0): expected %f, got %f", expected, result)
	}
}
func TestModulusPositiveNumbers(t *testing.T) {
	result := Modulus(10, 3)
	expected := 1
	if result != expected {
		t.Errorf("Modulus(10, 3): expected %d, got %d", expected, result)
	}
}

func TestModulusNegativeDividend(t *testing.T) {
	result := Modulus(-10, 3)
	expected := -10 % 3
	if result != expected {
		t.Errorf("Modulus(-10, 3): expected %d, got %d", expected, result)
	}
}

func TestModulusNegativeDivisor(t *testing.T) {
	result := Modulus(10, -3)
	expected := 10 % -3
	if result != expected {
		t.Errorf("Modulus(10, -3): expected %d, got %d", expected, result)
	}
}

func TestModulusBothNegative(t *testing.T) {
	result := Modulus(-10, -3)
	expected := -10 % -3
	if result != expected {
		t.Errorf("Modulus(-10, -3): expected %d, got %d", expected, result)
	}
}

func TestModulusZeroDividend(t *testing.T) {
	result := Modulus(0, 5)
	expected := 0
	if result != expected {
		t.Errorf("Modulus(0, 5): expected %d, got %d", expected, result)
	}
}

func TestModulusByZero(t *testing.T) {
	result := Modulus(10, 0)
	expected := 0
	if result != expected {
		t.Errorf("Modulus(10, 0): expected %d, got %d", expected, result)
	}
}
func TestMaxFirstIsGreater(t *testing.T) {
	result := Max(10.0, 5.0)
	expected := 10.0
	if result != expected {
		t.Errorf("Max(10.0, 5.0): expected %f, got %f", expected, result)
	}
}

func TestMaxSecondIsGreater(t *testing.T) {
	result := Max(3.0, 7.0)
	expected := 7.0
	if result != expected {
		t.Errorf("Max(3.0, 7.0): expected %f, got %f", expected, result)
	}
}

func TestMaxBothEqual(t *testing.T) {
	result := Max(4.5, 4.5)
	expected := 4.5
	if result != expected {
		t.Errorf("Max(4.5, 4.5): expected %f, got %f", expected, result)
	}
}

func TestMaxWithNegativeNumbers(t *testing.T) {
	result := Max(-2.0, -5.0)
	expected := -2.0
	if result != expected {
		t.Errorf("Max(-2.0, -5.0): expected %f, got %f", expected, result)
	}
}

func TestMaxWithZero(t *testing.T) {
	result := Max(0.0, -3.0)
	expected := 0.0
	if result != expected {
		t.Errorf("Max(0.0, -3.0): expected %f, got %f", expected, result)
	}
}
func TestMinFirstIsSmaller(t *testing.T) {
	result := Min(3.0, 7.0)
	expected := 3.0
	if result != expected {
		t.Errorf("Min(3.0, 7.0): expected %f, got %f", expected, result)
	}
}

func TestMinSecondIsSmaller(t *testing.T) {
	result := Min(8.0, 2.0)
	expected := 2.0
	if result != expected {
		t.Errorf("Min(8.0, 2.0): expected %f, got %f", expected, result)
	}
}

func TestMinBothEqual(t *testing.T) {
	result := Min(5.5, 5.5)
	expected := 5.5
	if result != expected {
		t.Errorf("Min(5.5, 5.5): expected %f, got %f", expected, result)
	}
}

func TestMinWithNegativeNumbers(t *testing.T) {
	result := Min(-4.0, -1.0)
	expected := -4.0
	if result != expected {
		t.Errorf("Min(-4.0, -1.0): expected %f, got %f", expected, result)
	}
}

func TestMinWithZero(t *testing.T) {
	result := Min(0.0, 3.0)
	expected := 0.0
	if result != expected {
		t.Errorf("Min(0.0, 3.0): expected %f, got %f", expected, result)
	}
	result = Min(-2.0, 0.0)
	expected = -2.0
	if result != expected {
		t.Errorf("Min(-2.0, 0.0): expected %f, got %f", expected, result)
	}
}
