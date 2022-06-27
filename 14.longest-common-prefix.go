package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	isSame := true
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < 1 {
			isSame = false
			break
		}

		if strs[0][0:1] != strs[i][0:1] {
			isSame = false
			break
		}
	}
	if isSame == false {
		return ""
	}

	subStrs := []string{}
	for _, str := range strs {
		subStrs = append(subStrs, str[1:])
	}
	return strs[0][0:1] + longestCommonPrefix(subStrs)
}
