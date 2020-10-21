package binarySearch

// Search uses binary search to return index of target input
func Search(array []int, target int) int {
	return help(array, target, 0, len(array)-1)
}

func help(array []int, target int, left int, right int) int {
	for left <= right {
		midpoint := (left + right) / 2
		num := array[midpoint]

		if num == target {
			return midpoint
		} else if target < num {
			right = midpoint - 1
		} else {
			left = midpoint + 1
		}
	}
	return -1
}
