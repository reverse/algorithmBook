package twosum

// TSBrute is the brute force solution to the Two Sum Algorithm
func TSBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for num := 0; num < len(nums); num++ {
			if nums[num]+nums[i] == target && (num != i) {
				return []int{num, i}
			}
		}
	}
	return []int{}
}

// TSHashMap utilizes a hashmap in a onceover to solve the Two Sum Algorithm
func TSHashMap(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := hashMap[complement]; ok && (val != i) {
			return []int{val, i}
		}
		hashMap[nums[i]] = i
	}
	return []int{}
}

// TSHashMapLong utilizes a hash map but it iterates over the data twice to solve the Two Sum Algorithm
func TSHashMapLong(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := hashMap[complement]; ok && (val != i) {
			return []int{val, i}
		}
	}
	return []int{}
}
