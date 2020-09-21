package nthfib

// FibBrute is the most inefficient solution, it doesn't save values
func FibBrute(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}

	return FibBrute(n-1) + FibBrute(n-2)
}

// FibHash uses a hashmap to store already found values, so it is more efficient than constantly recalculating
func FibHash(n int) int {
	mem := FibHashHelper(n, map[int]int{
		1: 0,
		2: 1,
	})
	return mem
}

// FibHashHelper is a part of FibHash, this is where it becomes recursive
func FibHashHelper(n int, hash map[int]int) int {
	if val, ok := hash[n]; ok {
		return val
	}

	hash[n] = FibHashHelper(n-2, hash) + FibHashHelper(n-1, hash)
	return hash[n]
}

// FibHashCount is the most effiecient in terms of space complexity as it only saves needed values
func FibHashCount(n int) int {
	init := []int{0, 1}
	count := 3

	for count <= n {
		fib := init[0] + init[1]
		init[0], init[1] = init[1], fib
		count++
	}

	if n == 1 {
		return init[0]
	}

	return init[1]
}
