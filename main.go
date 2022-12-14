package main

import (
	"fmt"
	"math"
)

func main() {
	printPrimeNumbers(5, 19)
	printPrimeNumbers(0, 2)
	printPrimeNumbers(13, 100)
}

func printPrimeNumbers(num1, num2 int) {
	//Step 1: Define a function that accepts two numbers, num1 and num2, type is int.
	//Step 2: Iterate between num1 and num2.
	//Step 3: If the number is prime, then print that number, else break.

	if num1 < 2 || num2 < 2 {
		fmt.Println("Error: Numbers must be greater than 2.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}
