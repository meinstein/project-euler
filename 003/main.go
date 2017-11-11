package main

import (
	"fmt"
	"math"
)

const Val = 600851475143

func main() {
	var i int64 = int64(math.Sqrt(float64(Val)))
	for ; i > 1; i-- {
		if Val%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}

func isPrime(x int64) bool {
	var i int64 = 2
	for ; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
