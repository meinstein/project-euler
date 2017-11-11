package main

import (
	"fmt"
)

func main() {
	i := 20

	for {
		if isDivisible(i) {
			break
		}
		i = i + 20
	}
	fmt.Println(i)
}

func isDivisible(n int) bool {
	for i := 2; i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
