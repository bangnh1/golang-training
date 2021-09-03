package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var inputString string
	fmt.Println("Enter a number >2 and <= 100000: ")
	fmt.Scan(&inputString)
	input := inputString2Float(inputString)
	PrintPrimeNumbers(input)
}

/*
Input validation and convert from string to int type
*/
func inputString2Float(inputString string) int {
	input, err := strconv.Atoi(inputString)
	if err != nil || input < 2 || input > 100000 {
		fmt.Println("Please enter a valid number")
		os.Exit(1)
	}

	return input
}

/*
Find prime numbers between 0 and num
*/
func PrintPrimeNumbers(num int) []int {
	var primes []int
	start := 2

	for start <= num {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(start))); i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", start)
			primes = append(primes, start)
		}
		start++
	}
	return primes
}
