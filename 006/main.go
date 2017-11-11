package main

import "fmt"

func main() {
	sumOfSquares, squareOfSum, diff := 0, 0, 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squareOfSum += i
	}

	squareOfSum = squareOfSum * squareOfSum
	diff = squareOfSum - sumOfSquares
	fmt.Println(diff)
}
