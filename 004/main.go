package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := 0

	for j := 999; j > 99; j-- {
		for i := 999; i > 99; i-- {
			n := i * j
			if reverse(n) {
				if n > p {
					p = n
				}
			}
		}
	}
	fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is %d\n", p)
}

func reverse(value int) bool {
	intString := strconv.Itoa(value)
	newString := ""

	for x := len(intString); x > 0; x-- {
		newString += string(intString[x-1])
	}

	newInt, _ := strconv.Atoi(newString)
	return newInt == value
}
