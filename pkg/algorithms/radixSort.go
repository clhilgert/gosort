package algorithms

func RadixSort(arr []string) []string {
	maxLen := 0
	for _, str := range arr {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	for pos := maxLen - 1; pos >= 0; pos-- {
		arr = sortByChar(arr, pos)
	}
	return arr
}

func sortByChar(arr []string, pos int) []string {
	count := make([]int, 256)
	output := make([]string, len(arr))

	for _, str := range arr {
		char := byte(0)
		if pos < len(str) {
			char = str[pos]
		}
		count[char]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		str := arr[i]
		char := byte(0)
		if pos < len(str) {
			char = str[pos]
		}
		output[count[char]-1] = str
		count[char]--
	}
	return output
}
