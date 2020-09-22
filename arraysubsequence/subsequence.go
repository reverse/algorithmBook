package arraysubsequence

// ArraySub is o(n) time complexity and returns true if the given sequence is a subsequence of the array
func ArraySub(array, sequence []int) bool {
	arrayTicker := 0
	sequenceTicker := 0

	for sequenceTicker < len(sequence) && arrayTicker < len(array) {
		if sequence[sequenceTicker] == array[arrayTicker] {
			sequenceTicker++
		}
		arrayTicker++
	}

	return sequenceTicker == len(sequence)
}

// StringSub returns true if given string (sub) is a substring of other given string (initial)
func StringSub(initial, sub string) bool {
	startInitial := 0
	startSub := 0

	for startInitial < len(initial) && startSub < len(sub) {
		if initial[startInitial] == sub[startSub] {
			startSub++
		}
		startInitial++
	}

	return startSub == len(sub)
}
