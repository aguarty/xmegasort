package xmegasort

import "math/rand"

// mergeCfg -
type QuickCfg struct {
	Threshold     int
	MaxGoroutines int
}

// DefaultQuickConfig -
func DefaultQuickConfig() *QuickCfg {
	return &QuickCfg{
		Threshold:     1000,
		MaxGoroutines: 100,
	}
}

// QuickSortX - Mega async Quicksort algorithm for sorting the elements of array in ascending order, accelerated by InsertionSort.
// Fastest on the big arrays (>1.000.000 elements)
// recommended MAXGOROUTINES=100
// "threshold" is that size of array for activate InsertionSort. Recomended - 1000
func QuickSortX(items []int, cfg *QuickCfg) []int {
	if len(items) <= 1 {
		return items
	}
	if cfg == nil {
		cfg = DefaultQuickConfig()
	}
	workers := make(chan int, cfg.MaxGoroutines-1)
	for i := 0; i < (cfg.MaxGoroutines - 1); i++ {
		workers <- 1
	}
	return xQuick(items, nil, workers, cfg.Threshold)
}

// QuicksortIns standart quicksort algorithm for sorting the elements of array in ascending order, accelerated by InsertionSort
// "threshold" is that size of array for activate InsertionSort. Recomended - 1000
func QuicksortIns(items []int, cfg *QuickCfg) []int {
	if cfg == nil {
		cfg = DefaultQuickConfig()
	}
	if len(items) < cfg.Threshold {
		return InsertionSort(items)
	}
	p := pivot(items)
	low, high := partition(items, p)
	QuicksortIns(items[:low], cfg)
	QuicksortIns(items[high:], cfg)
	return items
}

// Quicksort standart quicksort algorithm for sorting the elements of array in ascending order.
func QuickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	p := pivot(items)
	low, high := partition(items, p)
	QuickSort(items[:low])
	QuickSort(items[high:])
	return items
}

// xQuick sorting worker
func xQuick(items []int, done chan int, workers chan int, threshold int) []int {
	// report to caller that we're finished
	if done != nil {
		defer func() { done <- 1 }()
	}
	if len(items) < threshold {
		return InsertionSort(items)
	}
	p := pivot(items)
	low, high := partition(items, p)
	doneChannel := make(chan int, 2)
	select {
	case <-workers:
		go xQuick(items[:low], doneChannel, workers, threshold)
	default:
		xQuick(items[:low], nil, workers, threshold)
		doneChannel <- 1
	}
	select {
	case <-workers:
		go xQuick(items[high:], doneChannel, workers, threshold)
	default:
		xQuick(items[high:], nil, workers, threshold)
		doneChannel <- 1
	}
	<-doneChannel
	<-doneChannel
	return items
}

// Partition reorders the elements of items so that:
// - all elements in items[:low] are less than p,
// - all elements in items[low:high] are equal to p,
// - all elements in items[high:] are greater than p.
func partition(items []int, p int) (low, high int) {
	low, high = 0, len(items)
	for mid := 0; mid < high; {
		// Initemsariant:
		//  - items[:low] < p
		//  - items[low:mid] = p
		//  - items[mid:high] are unknown
		//  - items[high:] > p
		//
		//         < p       = p        unknown       > p
		//     -----------------------------------------------
		// items: |          |          |a            |           |
		//     -----------------------------------------------
		//                ^          ^             ^
		//               low        mid           high
		switch a := items[mid]; {
		case a < p:
			items[mid] = items[low]
			items[low] = a
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			items[mid] = items[high-1]
			items[high-1] = a
			high--
		}
	}
	return
}

// pivot - get pivot element
func pivot(items []int) int {
	n := len(items)
	return median(items[rand.Intn(n)],
		items[rand.Intn(n)],
		items[rand.Intn(n)])
}

// median - choose median
func median(a, b, c int) int {
	if a < b {
		switch {
		case b < c:
			return b
		case a < c:
			return c
		default:
			return a
		}
	}
	switch {
	case a < c:
		return a
	case b < c:
		return c
	default:
		return b
	}
}
