package diffsquares

import "math"

/*
func findPrimes(input int) []int {
	var primes []int
	primes = append(primes, 1)
	for i := 1; i <= input; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func SquareOfSum(input int) int {
	primes := findPrimes(input)
	sum := 0
	for _, num := range primes {
		sum = sum + num
	}
	return int(math.Pow(float64(sum), 2))
}
func SumOfSquares(input int) int {
	primes := findPrimes(input)
	sum := 0
	for _, num := range primes {
		sum = sum + int(math.Pow(float64(num), 2))
	}
	return sum
}
*/

// SquareOfSum returns the square of the sum of the first N natural numbers (input = N)
func SquareOfSum(input int) int {
	sum := 0
	for i := 1; i <= input; i++ {
		sum = sum + i
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns the sum of the squares of the first N natural numbers (input = N)
func SumOfSquares(input int) int {
	sum := 0
	for i := 1; i <= input; i++ {
		sum = sum + int(math.Pow(float64(i), 2))
	}
	return sum
}

// Difference returns the differences between SquareOfSum and SumOfSquares
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
