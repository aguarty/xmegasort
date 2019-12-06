# Xmegasort
The set of algorithms for sorting the elements of array in ascending order.
- **InsertionSort**   
Is a simple sorting algorithm that is relatively efficient for small lists and mostly sorted lists, and is often used as part of more sophisticated algorithms.
- **HeapSort**    
Is a much more efficient version of selection sort. It also works by determining the largest (or smallest) element of the list, placing that at the end (or beginning) of the list, then continuing with the rest of the list, but accomplishes this task efficiently by using a data structure called a heap, a special type of binary tree.
- **QuickSort**   
Is a divide and conquer algorithm which relies on a partition operation: to partition an array, an element called a pivot is selected. All elements smaller than the pivot are moved before it and all greater elements are moved after it. This can be done efficiently in linear time and in-place. The lesser and greater sublists are then recursively sorted.
- **QuickSortIns**    
Is a custom quicksort, accelerated by InsertionSort.
- **QuickSortX**    
Is a custom quicksort with Async work and accelerated by InsertionSort.
- **MergeSort**   
Merge sort takes advantage of the ease of merging already sorted lists into a new sorted list.
- **MergeSortAsync**    
Is a custom mergesort with Async work

### Benchmarks
In this benchs SortInts is a standart golang library "sort".    
Algorithm QuickSortX is a winner in average. 
```
goos: linux
goarch: amd64
pkg: xmegasort

[1:InsertionSort, 2:QuickSortIns, 3:QuickSortX]
Benchmark_1k_SortInts-6                30187       36308 ns/op        32 B/op        1 allocs/op
Benchmark_1k_InsertionSort-6         1534983         771 ns/op         0 B/op        0 allocs/op
Benchmark_1k_HeapSort-6                16882       71241 ns/op         0 B/op        0 allocs/op
Benchmark_1k_QuickSort-6               18548       64268 ns/op         0 B/op        0 allocs/op
Benchmark_1k_QuickSortIns-6           516468        2292 ns/op         0 B/op        0 allocs/op
Benchmark_1k_QuickSortX-6             191682       17668 ns/op      1008 B/op        2 allocs/op
Benchmark_1k_MergeSort-6               13161      103250 ns/op     81536 B/op      999 allocs/op
Benchmark_1k_MergeSortAsync-6          15753       71025 ns/op     81713 B/op     1003 allocs/op

[1:InsertionSort, 2:QuickSortX, 3:QuickSortIns]
Benchmark_10k_SortInts-6                 2200      538027 ns/op        32 B/op        1 allocs/op
Benchmark_10k_InsertionSort-6         154311        7736 ns/op         0 B/op        0 allocs/op
Benchmark_10k_HeapSort-6                1302      922701 ns/op         0 B/op        0 allocs/op
Benchmark_10k_QuickSort-6               2074      563877 ns/op         0 B/op        0 allocs/op
Benchmark_10k_QuickSortIns-6           15192       80169 ns/op         0 B/op        0 allocs/op
Benchmark_10k_QuickSortX-6             15120       80127 ns/op      2701 B/op       17 allocs/op
Benchmark_10k_MergeSort-6                855     2046507 ns/op   1127169 B/op     9999 allocs/op
Benchmark_10k_MergeSortAsync-6          2906      422680 ns/op   1129786 B/op    10093 allocs/op

[1:QuickSortX, 2:QuickSortIns, 3:QuickSort]
Benchmark_100k_SortInts-6                194     6022563 ns/op        32 B/op        1 allocs/op
Benchmark_100k_InsertionSort-6             1  1311486513 ns/op         0 B/op        0 allocs/op
Benchmark_100k_HeapSort-6                109    10793682 ns/op         0 B/op        0 allocs/op
Benchmark_100k_QuickSort-6               466     2515772 ns/op         0 B/op        0 allocs/op
Benchmark_100k_QuickSortIns-6            834     1313140 ns/op         0 B/op        0 allocs/op
Benchmark_100k_QuickSortX-6             1387      783037 ns/op     19748 B/op      169 allocs/op
Benchmark_100k_MergeSort-6                84    24803366 ns/op  14306304 B/op    99999 allocs/op
Benchmark_100k_MergeSortAsync-6          208     5215945 ns/op  12780479 B/op   100738 allocs/op

[1:QuickSortX, 2:QuickSort, 3:QuickSortIns]
Benchmark_1m_SortInts-6                   18    57470669 ns/op        32 B/op        1 allocs/op
Benchmark_1m_HeapSort-6                    9   122462235 ns/op         0 B/op        0 allocs/op
Benchmark_1m_QuickSort-6                  61    18465592 ns/op         0 B/op        0 allocs/op
Benchmark_1m_QuickSortIns-6               57    19525665 ns/op         0 B/op        0 allocs/op
Benchmark_1m_QuickSortX-6                 88    13640823 ns/op    171538 B/op     1524 allocs/op
Benchmark_1m_MergeSort-6                   8   130147308 ns/op  164767744 B/op    999999 allocs/op
Benchmark_1m_MergeSortAsync-6             37    29177067 ns/op  116915139 B/op   1005814 allocs/op

[1:QuickSortX, 2:MergeSortAsync, 3:QuickSort]
Benchmark_10m_SortInts-6                   1  1294825307 ns/op        32 B/op        1 allocs/op
Benchmark_10m_HeapSort-6                   1  2909912017 ns/op         0 B/op        0 allocs/op
Benchmark_10m_QuickSort-6                  4   266210205 ns/op         0 B/op        0 allocs/op
Benchmark_10m_QuickSortIns-6               4   269342834 ns/op         0 B/op        0 allocs/op
Benchmark_10m_QuickSortX-6                 7   152180517 ns/op    908941 B/op     8108 allocs/op
Benchmark_10m_MergeSort-6                  1  1363606702 ns/op  1949677568 B/op  9999999 allocs/op
Benchmark_10m_MergeSortAsync-6             6   203232707 ns/op  996696000 B/op  10093423 allocs/op
```
