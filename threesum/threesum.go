package threesum

// ThreeBrute uses a brute force method to solve the Three Sum Problem
func ThreeBrute(input []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(input); i++ {
		if i > 0 && input[i] == input[i-1] {
			continue
		}
		for e := i + 1; e < len(input); e++ {
			if e > i+1 && input[e] == input[e-1] {
				continue
			}
			for f := e + 1; f < len(input); f++ {
				if f > e+1 && input[f] == input[f-1] {
					continue
				}
				if input[i]+input[e]+input[f] == 0 {
					result = append(result, []int{input[i], input[e], input[f]})
				}
			}
		}
	}
	return result
}
