package main

import (
	"fmt"
	"math"
	)


func getPrimes(N int) []int {
	// Eratosphen algo
	isComposite := make([]bool, N)
	primes := []int{}
	for i := 2; i < N; i ++ {
		if !isComposite[i] {
			primes = append(primes, i)
			for x := i+i; x < N; x += i {
				isComposite[x] = true
			}
		}
	}
	return primes
}


func getPrimeFactors(n int, primes []int) map[int]float64 {
	primeFactors := make(map[int]float64) // prime : power
	for n != 1 {
		for i := 0; i < len(primes); i ++ {
			if n % primes[i] == 0 {
				primeFactors[primes[i]] += 1.0
				n = n/primes[i]
				break
			}
		}
	}
	return primeFactors
}

func eulerFunc(n int) int{
	eulerResult := 1.0
	for prime, power := range getPrimeFactors(n, getPrimes(n)){
			eulerResult *= math.Pow(float64(prime), power) - 
			math.Pow(float64(prime), power-1)
		}
	return int(eulerResult)
}

func primeTest(value int, primeFactors map[int]int) bool{
	primeCheck := 1.0
	for key, value := range primeFactors {
    	primeCheck *= math.Pow(float64(key), float64(value))
	}
	return (value == int(primeCheck))
}

func main() {
	number := 200
	fmt.Println(eulerFunc(number))
}