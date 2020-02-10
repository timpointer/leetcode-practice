package main

import "strconv"

import "fmt"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	b := []byte(s)

	for i := 0; i < len(b)/2; i++ {
		fmt.Printf("%d : %d\n", i, len(b)-1-i)
		if b[i] != b[len(b)-1-i] {
			return false
		}
	}
	return true
}
