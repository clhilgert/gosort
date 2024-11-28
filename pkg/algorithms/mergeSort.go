package algorithms

func MergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []string) []string {
	sorted := []string{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}
	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)
	return sorted
}
