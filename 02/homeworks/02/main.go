package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Solver: ax^2 + bx + c = 0")
	fmt.Print("Enter the a, b, c of Quadratic equation = ")

	var inputA, inputB, inputC string
	fmt.Scanln(&inputA, &inputB, &inputC)

	a := inputString2Float(inputA)
	b := inputString2Float(inputB)
	c := inputString2Float(inputC)

	if a == 0 {
		fmt.Println("This is not a quadratic equation !!!")
		os.Exit(1)
	}

	fmt.Printf("Solver: %f*x^2 + %f*x + %f = 0 \n", a, b, c)
	ResolveQuadraticEquation(a, b, c)

}

/*
Input validation and convert from string to float64 type
*/
func inputString2Float(inputString string) float64 {
	input, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		fmt.Println("Please enter a valid number")
		os.Exit(1)
	}
	return input
}

func ResolveQuadraticEquation(a, b, c float64) {
	delta := b*b - 4*a*c

	switch {
	case delta > 0:
		fmt.Println(delta)
		x1 := (-b + math.Sqrt(delta)/(2*a))
		x2 := (-b - math.Sqrt(delta)/(2*a))
		fmt.Println("Two Distinct Real Roots Exist: x1=", x1, "; x2=", x2)
	case delta == 0:
		x := -b / (2 * a)
		fmt.Println("Two Equal and Real Roots Exist:: x=", x)
	case delta < 0:
		x := -b / (2 * a)
		imaginary := math.Sqrt(-delta) / (2 * a)
		fmt.Printf("Two Distinct Complex Roots Exist: x1=%f + %fi; x2=%f - %fi", x, imaginary, x, imaginary)
	}
}
