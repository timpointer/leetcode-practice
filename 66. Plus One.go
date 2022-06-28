package main

func plusOne(digits []int) []int {
	// get last int and add 1, get result
	// if result == 0 , set adder is 1

	// get next last int add 1, get result
	// if result == 0 , set adder is 1

	// util have no value in the list
	// if adder is 1 , add one in the head

	adder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + adder
		if v == 10 {
			digits[i] = v - 10
			adder = 1
		} else {
			digits[i] = v
			adder = 0
		}
	}
	if adder == 0 {
		return digits
	} else {
		return append([]int{1}, digits...)
	}
}
