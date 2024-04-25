package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121) == true)
	fmt.Println(isPalindrome(-121) == false)
	fmt.Println(isPalindrome(-333) == false)
	fmt.Println(isPalindrome(333) == true)
	fmt.Println(isPalindrome(123321) == true)
	fmt.Println(isPalindrome(1235321) == true)
}

func isPalindrome(x int) bool {
	input := x
	output := 0
	for input > 0 {
		r := input % 10
		output = output*10 + r
		input = input / 10
	}
	if output != x {
		return false
	} else {
		return true
	}
}
