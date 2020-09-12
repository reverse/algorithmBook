package twosum

// TSTwoPointer solves two sum but it is quicker than the hashmap method due to the input slice being sorted (numbers is sorted)
func TSTwoPointer(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		complement := numbers[left] + numbers[right]
		if complement > target {
			right--
		} else if complement < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
