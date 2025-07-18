package math
import "errors"

// Returns the sum of two float64.
func Sum(a float64, b float64) float64 {
	return a + b
}
// Returns the difference of two float64.
func Subtract(a float64, b float64) float64 {
	return a - b
}
// Returns the product of two float64.
func Multiply(a float64, b float64) float64 {
	return a * b
}
// Returns the quotient of two float64.
func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil	
}
// Returns the modulus of two integers.
func Modulus(a int, b int) int {
	if b == 0 {
		return 0 // Handle modulus by zero
	}
	return a % b
}
// Returns the maximum of two float64.
func Max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
// Returns the minimum of two float64.
func Min(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
