package twosum

// TSBrute is the brute force solution to the Two Sum Algorithm
func TSBrute(input []int, target int) []int {
	for i := 0; i < len(input); i++ {
		for num := 0; num < len(input); num++ {
			if input[num]+input[i] == target {
				return []int{num, i}
			}
		}
	}
	return []int{}
}

// TSHashMap utilizes a hashmap in a onceover to solve the Two Sum Algorithm
func TSHashMap(input []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(input); i++ {
		complement := target - input[i]
		if val, ok := hashMap[complement]; ok && (val != i) {
			return []int{val, i}
		}
		hashMap[input[i]] = i
	}
	return []int{}
}

// TSHashMapLong utilizes a hash map but it iterates over the data twice to solve the Two Sum Algorithm
func TSHashMapLong(input []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(input); i++ {
		hashMap[input[i]] = i
	}

	for i := 0; i < len(input); i++ {
		if val, ok := hashMap[input[i]]; ok && (val != i) {
			return []int{val, i}
		}
	}
	return []int{}
}
