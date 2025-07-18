package main

import (
	"go_lab/v1/math"
	"net/http"
)

var messageHello string = "Welcome to Go Lab!"

func main() {
	println(getMessageHello())
	learnOperationMethod(5.0, 0.0)
	learnHttpMethod()
}

// Return welcome message
func getMessageHello() string {
	return messageHello
}

// Perform some math operations and print the results.
func learnOperationMethod(localA float64, localB float64) {
	println("Performing math operations...")
	println("Local A: ", localA)
	println("Local B: ", localB)
	println("Performing operations with A and B...")
	resultSum := math.Sum(localA, localB)
	println("Result sum ", resultSum)
	resultSubtract := math.Subtract(localA, localB)
	println("Result subtract ", resultSubtract)
	resultMultiply := math.Multiply(localA, localB)
	println("Result multiply ", resultMultiply)
	resultDivide, error := math.Divide(localA, localB)
	if error != nil {
		println("Error dividing: ", error.Error())
	} else {
		println("Result divide ", resultDivide)
	}
	resultModulus := math.Modulus(int(localA), int(localB))
	println("Result modulus ", resultModulus)
	resultMax := math.Max(localA, localB)
	println("Result max ", resultMax)
	resultMin := math.Min(localA, localB)
	println("Result min ", resultMin)
}

func learnHttpMethod() {
	println("Learning HTTP method...")
	res, err := http.Get("http://google.com")
	if err != nil {
		println("Error: ", err)
		return
	}
	defer res.Body.Close()
	println("Response status: ", res.Status)
}
