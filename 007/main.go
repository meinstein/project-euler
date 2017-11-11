package main

import "fmt"

func main() {
	i, counter := 2, 0

	for {
		if isPrime(i) {
			counter++
			if counter == 10001 {
				break
			}
		}
		i++
	}

	fmt.Printf("The 10001st prime number is: %d\n", i)
}

func isPrime(n int) bool {
	startingInt := 2

	if n%2 != 0 {
		startingInt = 3
	}

	for i := startInt; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
