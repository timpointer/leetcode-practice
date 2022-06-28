package main

func isValid(str string) bool {
	list := make([]rune, len(str), len(str))
	count := -1
	for _, s := range str {
		if s == '(' {
			count++
			list[count] = ')'
		} else if s == '{' {
			count++
			list[count] = '}'
		} else if s == '[' {
			count++
			list[count] = ']'
		}
		if s == ')' || s == '}' || s == ']' {
			if count < 0 {
				return false
			}
			en := list[count]
			if s != en {
				return false
			}
			count--
		}
	}
	if count != -1 {
		return false
	}
	return true
}
