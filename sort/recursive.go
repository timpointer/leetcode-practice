package sort

func Reversal(txt string) string {
	if len(txt) == 1 {
		return txt
	}
	return Reversal(txt[1:]) + txt[0:1]
}

func CountX(txt string) int {
	if len(txt) == 0 {
		return 0
	}
	if txt[0] == 'x' {
		return 1 + CountX(txt[1:])
	} else {
		return CountX(txt[1:])
	}
}

func CountStep(stair int) int {
	if stair <= 0 {
		return 0
	} else if stair == 1 {
		return 1
	} else if stair == 2 {
		return 2
	} else if stair == 3 {
		return 4
	}

	paths := CountStep(stair - 1)
	paths += CountStep(stair - 2)
	paths += CountStep(stair - 3)
	return paths
}

func CountStep2(stair int) int {
	if stair < 0 {
		return 0
	} else if stair == 1 || stair == 0 {
		return 1
	}

	paths := CountStep2(stair - 1)
	paths += CountStep2(stair - 2)
	paths += CountStep2(stair - 3)
	return paths
}

func Anagram(txt string) []string {
	if len(txt) == 0 {
		return []string{}
	} else if len(txt) == 1 {
		return []string{txt}
	}

	var result []string
	for i, char := range txt {
		subTxts := Anagram(txt[0:i] + txt[i+1:])
		for _, subTxt := range subTxts {
			result = append(result, string(char)+subTxt)
		}
	}
	return result
}
