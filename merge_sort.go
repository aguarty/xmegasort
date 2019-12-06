package xmegasort

import "sync"

// mergeCfg -
type MergeCfg struct {
	Threshold     int
	MaxGoroutines int
}

// DefaultMergeConfig -
func DefaultMergeConfig() *MergeCfg {
	return &MergeCfg{
		Threshold:     500,
		MaxGoroutines: 100,
	}
}

// MergeSort standart mergesort algorithm for sorting the elements of array in ascending order.
func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)
}

// MergeSortAsync async MergeSort algorithm for sorting the elements of array in ascending order.
// recommended MAXGOROUTINES=100
// "threshold" is that size of array for activate InsertionSort. Recomended - 500
func MergeSortAsync(items []int, cfg *MergeCfg) []int {
	if len(items) <= 1 {
		return items
	}
	if cfg == nil {
		cfg = DefaultMergeConfig()
	}
	sem := make(chan struct{}, cfg.MaxGoroutines)
	return xMerge(items, sem, cfg.Threshold)
}

// xMerge sorting worker
func xMerge(items []int, mergeSem chan struct{}, threshold int) []int {
	if len(items) <= 1 {
		return items
	}
	n := len(items) / 2
	if len(items) <= threshold {
		return MergeSort(items)
	}
	var l []int
	var r []int
	wg := sync.WaitGroup{}
	wg.Add(2)
	select {
	case mergeSem <- struct{}{}:
		go func() {
			l = xMerge(items[:n], mergeSem, threshold)
			<-mergeSem
			wg.Done()
		}()
	default:
		l = xMerge(items[:n], mergeSem, threshold)
		wg.Done()
	}
	select {
	case mergeSem <- struct{}{}:
		go func() {
			r = xMerge(items[n:], mergeSem, threshold)
			<-mergeSem
			wg.Done()
		}()
	default:
		l = xMerge(items[:n], mergeSem, threshold)
		wg.Done()
	}
	wg.Wait()
	return merge(l, r)
}

// merge -
func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}
	return r
}
