package xmegasort

// InsertionSort insertion sort algorithm
func InsertionSort(items []int) []int {
	for j := 1; j < len(items); j++ {
		key := items[j]
		i := j - 1
		for i >= 0 && items[i] > key {
			items[i+1] = items[i]
			i--
		}
		items[i+1] = key
	}
	return items
}
