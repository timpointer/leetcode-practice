package main

func climbStairs(n int) int {
	return helper(n, map[int]int{})
}

func helper(n int, kv map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	cn1, ok := kv[n-1]
	if ok == false {
		cn1 = helper(n-1, kv)
		kv[n-1] = cn1
	}

	cn2, ok := kv[n-2]
	if ok == false {
		cn2 = helper(n-2, kv)
		kv[n-2] = cn2
	}

	kv[n] = cn1 + cn2
	return kv[n]
}
