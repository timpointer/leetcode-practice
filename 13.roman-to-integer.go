package main

func romanToInt(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	b := []byte(s)

	r := 0
	for i := range b {
		if i == len(b)-1 {
			break
		}
		if mapping[b[i]] < mapping[b[i+1]] {
			r -= mapping[b[i]]
		} else {
			r += mapping[b[i]]
		}
	}

	return r + mapping[b[len(b)-1]]
}
