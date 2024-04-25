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
	result := true
	if x < 0 {
		result = false
	} else {
		digits := make([]int, 0, 10)
		for x > 0 {
			digits = append(digits, x%10)
			x = x / 10
		}
		for i := 0; i < len(digits)/2; i++ {
			if digits[i] != digits[len(digits)-i-1] {
				result = false
				break
			}
		}
	}
	return result
}
