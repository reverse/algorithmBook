package threesum

import "sort"

// ThreeBrute uses a brute force method to solve the Three Sum Problem
func ThreeBrute(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for e := i + 1; e < len(nums); e++ {
			if e > i+1 && nums[e] == nums[e-1] {
				continue
			}
			for f := e + 1; f < len(nums); f++ {
				if f > e+1 && nums[f] == nums[f-1] {
					continue
				}
				if nums[i]+nums[e]+nums[f] == 0 {
					result = append(result, []int{nums[i], nums[e], nums[f]})
				}
			}
		}
	}
	return result
}

// ThreeEfficient is the best solution with O(n^2) complexity, much faster than brute force. 2 pointer method
func ThreeEfficient(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		right := len(nums) - 1
		second := i + 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for second < right {
			sum := nums[right] + nums[i] + nums[second]
			if sum > 0 {
				right--
			} else if sum < 0 {
				second++
			} else {
				result = append(result, []int{nums[i], nums[second], nums[right]})
				second++
				for second < right && nums[second] == nums[second-1] {
					second++
				}
			}
		}
	}
	return result
}
