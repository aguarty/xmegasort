package xmegasort

// heapify -
func heapify(items []int, idx int, size int) {
	l := 2*idx + 1
	r := 2*idx + 2
	var largest int
	if l <= size && items[l] > items[idx] {
		largest = l
	} else {
		largest = idx
	}
	if r <= size && items[r] > items[largest] {
		largest = r
	}
	if largest != idx {
		t := items[idx]
		items[idx] = items[largest]
		items[largest] = t
		heapify(items, largest, size)
	}
}

// HeapSort heap sort
func HeapSort(items []int) []int {
	ln := len(items) //heap size
	m := int(ln / 2) //middle
	for i := m; i >= 0; i-- {
		heapify(items, i, ln-1)
	}
	f := ln - 1
	for j := f; j >= 0; j-- {
		t := items[0]
		items[0] = items[j]
		items[j] = t
		f = f - 1
		heapify(items, 0, f)
	}
	return items
}
