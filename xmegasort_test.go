package xmegasort

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestInsertionSort(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	r := InsertionSort(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, r)
}

func TestHeapSort(t *testing.T) {
	s := CreateTestSlice(9999)
	r := HeapSort(s)
	if !sort.IntsAreSorted(r) {
		t.Errorf("array is not sorted")
	}
}

func TestQuicksort(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	r := QuickSort(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, r)
}

func TestQuicksortIns(t *testing.T) {
	s := CreateTestSlice(9999)
	r := QuicksortIns(s, nil)
	if !sort.IntsAreSorted(r) {
		t.Errorf("array is not sorted")
	}
}

func TestQuicksortX(t *testing.T) {
	s := CreateTestSlice(9999)
	r := QuickSortX(s, nil)
	if !sort.IntsAreSorted(r) {
		t.Errorf("array is not sorted")
	}
}

func TestMergesort(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	r := MergeSort(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, r)
}

func TestMergesortAsync(t *testing.T) {
	s := CreateTestSlice(9999)
	r := MergeSortAsync(s, nil)
	if !sort.IntsAreSorted(r) {
		t.Errorf("array is not sorted")
	}
}

// =========== BENCHS ==========================================================

var (
	result []int
)

func benchmarkInsertionSort(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = InsertionSort(s)
	}
	result = r
}

func benchmarkHeapSort(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = HeapSort(s)
	}
	result = r
}

func benchmarkQuickSort(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = QuickSort(s)
	}
	result = r
}

func benchmarkQuickSortIns(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = QuicksortIns(s, nil)
	}
	result = r
}

func benchmarkQuickSortX(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = QuickSortX(s, nil)
	}
	result = r
}

func benchmarkMergeSort(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = MergeSort(s)
	}
	result = r
}

func benchmarkMergeSortAsync(b *testing.B, count int) {
	var r []int
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = MergeSortAsync(s, nil)
	}
	result = r
}

func benchmarkSortInts(b *testing.B, count int) {
	s := CreateTestSlice(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

// === 1k ===
func Benchmark_1k_InsertionSort(b *testing.B) {
	benchmarkInsertionSort(b, 1000)
}
func Benchmark_1k_HeapSort(b *testing.B) {
	benchmarkHeapSort(b, 1000)
}
func Benchmark_1k_QuickSort(b *testing.B) {
	benchmarkQuickSort(b, 1000)
}
func Benchmark_1k_QuickSortIns(b *testing.B) {
	benchmarkQuickSortIns(b, 1000)
}
func Benchmark_1k_QuickSortX(b *testing.B) {
	benchmarkQuickSortX(b, 1000)
}
func Benchmark_1k_MergeSort(b *testing.B) {
	benchmarkMergeSort(b, 1000)
}
func Benchmark_1k_MergeSortAsync(b *testing.B) {
	benchmarkMergeSortAsync(b, 1000)
}
func Benchmark_1k_SortInts(b *testing.B) {
	benchmarkSortInts(b, 1000)
}

// === 10k ===
func Benchmark_10k_InsertionSort(b *testing.B) {
	benchmarkInsertionSort(b, 10000)
}
func Benchmark_10k_HeapSort(b *testing.B) {
	benchmarkHeapSort(b, 10000)
}
func Benchmark_10k_QuickSort(b *testing.B) {
	benchmarkQuickSort(b, 10000)
}
func Benchmark_10k_QuickSortIns(b *testing.B) {
	benchmarkQuickSortIns(b, 10000)
}
func Benchmark_10k_QuickSortX(b *testing.B) {
	benchmarkQuickSortX(b, 10000)
}
func Benchmark_10k_MergeSort(b *testing.B) {
	benchmarkMergeSort(b, 10000)
}
func Benchmark_10k_MergeSortAsync(b *testing.B) {
	benchmarkMergeSortAsync(b, 10000)
}
func Benchmark_10k_SortInts(b *testing.B) {
	benchmarkSortInts(b, 10000)
}

// === 100k ===
func Benchmark_100k_InsertionSort(b *testing.B) {
	benchmarkInsertionSort(b, 100000)
}
func Benchmark_100k_HeapSort(b *testing.B) {
	benchmarkHeapSort(b, 100000)
}
func Benchmark_100k_QuickSort(b *testing.B) {
	benchmarkQuickSort(b, 100000)
}
func Benchmark_100k_QuickSortIns(b *testing.B) {
	benchmarkQuickSortIns(b, 100000)
}
func Benchmark_100k_QuickSortX(b *testing.B) {
	benchmarkQuickSortX(b, 100000)
}
func Benchmark_100k_MergeSort(b *testing.B) {
	benchmarkMergeSort(b, 100000)
}
func Benchmark_100k_MergeSortAsync(b *testing.B) {
	benchmarkMergeSortAsync(b, 100000)
}
func Benchmark_100k_SortInts(b *testing.B) {
	benchmarkSortInts(b, 100000)
}

// === 1m ===
func Benchmark_1m_HeapSort(b *testing.B) {
	benchmarkHeapSort(b, 1000000)
}
func Benchmark_1m_QuickSort(b *testing.B) {
	benchmarkQuickSort(b, 1000000)
}
func Benchmark_1m_QuickSortIns(b *testing.B) {
	benchmarkQuickSortIns(b, 1000000)
}
func Benchmark_1m_QuickSortX(b *testing.B) {
	benchmarkQuickSortX(b, 1000000)
}
func Benchmark_1m_MergeSort(b *testing.B) {
	benchmarkMergeSort(b, 1000000)
}
func Benchmark_1m_MergeSortAsync(b *testing.B) {
	benchmarkMergeSortAsync(b, 1000000)
}
func Benchmark_1m_SortInts(b *testing.B) {
	benchmarkSortInts(b, 1000000)
}

// === 10m ===
func Benchmark_10m_HeapSort(b *testing.B) {
	benchmarkHeapSort(b, 10000000)
}
func Benchmark_10m_QuickSort(b *testing.B) {
	benchmarkQuickSort(b, 10000000)
}
func Benchmark_10m_QuickSortIns(b *testing.B) {
	benchmarkQuickSortIns(b, 10000000)
}
func Benchmark_10m_QuickSortX(b *testing.B) {
	benchmarkQuickSortX(b, 10000000)
}
func Benchmark_10m_MergeSort(b *testing.B) {
	benchmarkMergeSort(b, 10000000)
}
func Benchmark_10m_MergeSortAsync(b *testing.B) {
	benchmarkMergeSortAsync(b, 10000000)
}
func Benchmark_10m_SortInts(b *testing.B) {
	benchmarkSortInts(b, 10000000)
}
