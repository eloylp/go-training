package insertionsort

func InsertionSort(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	copy(list, sorted)
}

func insert(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}
