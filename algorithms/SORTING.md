# Array Sorting Algorithms

## O (n log n ) Sort Algorithms

#### Quicksort
- Comparison sort (meaning it can sort items of any type for which a "less than" relation is defined)
- Worst O(n^2) | Avg O(n log n) | Best O(n log n)
- Can operate in-place on an array requiring no add. memory
- Similar to selection sort
- Divide and conquer algo
    1. Pick element called a pivot from array
    2. Partitioning: reorder the array so that all elements value less than 
the pivot
    3. Recursively* apply the above steps to the sub-array of elements with smaller values and sub-arr with greater values

#### Mergesort 
Efficient gen-purpose comparison based sorting algo
- Most implementations produce a stable sort
- Divide and conquer
    1. Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
    2. Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.
- Avg & Worst O(n log n)
- Top-Down or Bottom-Up
- Space Complex O(n)

#### Heapsort
Space Complex O(1)

#### Radix Sort
(comparison based) - fast but store in memory
- O(wn) w = word length; n = keys;
- If w not constant, then O(n log n) <- Worse then other comp algos O(log n)

## Basic Sorts
#### Bubble Sort
Sorting by exchage
- simple
	- step through list
	- compare adjacent pairs 
	- swaps if they are in the wrong order
	- The pass through the list is repeated until the list is sorted.
- (Comp Sort)
- Named because smaller/larger elements will "bubble" to the top of list
- Simple but too slow and impractical for most problems (naive)
- Could be practical if input is in mostly sorted order.***
- O(n^2)

#### Insertion Sort
Sorting by insertion

#### Selection Sort
Sort by selecting


## Other Sorts

#### Timsort 
- Hybrid sort 
- Insertion sort to combine runs smaller than minimum run size and merge sort otherwise.



#### Tree Sort
Insertion sort

A tree sort is a sort algorithm that builds a binary search tree from the elements to be sorted, and then traverses the tree (in-order) so that the elements come out in sorted order. 

Its typical use is sorting elements online: after each insertion, the set of elements seen so far is available in sorted order.
- Avg O(log n)
- Trees require memory to be allocated on the heap, which is a performance hit

#### Splaysort 
Adaptive comparison sorting algorithm basedo on the splay tree data structure.


#### Shell Sort
Insertion sort


#### Bucket Sort
TODO

#### Counting Sort

#### Cubesort


## Vocabulary
**Comparison sort** - meaning it can sort items of any type for which a "less than" relation is defined
