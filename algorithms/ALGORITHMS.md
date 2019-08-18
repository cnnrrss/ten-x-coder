## Algorithms

## Big O
Big O Notation of all datastructures
- Insert, Delete, Access, Search, Traversal \
Big O Notation of sorting algorithms \
Dominance relations


![](../assets/big_O_complexity_chart.png)
Link: http://bigocheatsheet.com


If you're not changing the data, the other main option is to change the code.

The biggest improvement is likely to come from an algorithmic change. This is the equivalent of replacing *bubble sort* (O(n^2)) with *quicksort* (O(n log n)) or replacing a linear *scan through an array* (O(n)) with a *binary search* (O(log n)) or a *map lookup* (O(1)).

This is how software becomes slow. Structures originally designed for one use is repurposed for something it wasn't designed for. This happens gradually.

It's important to have an intuitive grasp of the different big-O levels. Choose the right data structure for your problem. You don't have to always shave cycles, but this just prevents dumb performance issues that might not be noticed until much later.


The basic classes of complexity are:
## Time Complexities
`n! -> 2^n -> n^3 -> n^2 -> n * log(n) -> n -> log(n) -> 1`

#### O(1)
A field access, array or map lookup

Advice: don't worry about it (but keep in mind the constant factor.)

**Constant functions 		f(n) = 1** \
	- No dependence on the parameter n \

#### O(log n): binary search

Advice: only a problem if it's in a loop

**Logarithmic functions	f(n) = log(n)** \
	- Binary search \
	- Functions grow quite slowly as n gets big


#### O(n): simple loop

Advice: you're doing this all the time

**Linear functions		f(n) = n** \
	- Measure the cost of looking at each item once (or twice or ten times) \
	- Identify the biggest item, smallest item, compute average value

#### O(n log n): divide-and-conquer, sorting

Advice: still fairly fast

**Superlinear functions	f(n) = n * log(n)** \
	- This important class of functions arises in algorithms such as Quicksort and Mergesort. \
	- They grow just a little faster than linear

#### O(n*m): nested loop / quadratic

Advice: be careful and constrain your set sizes

**Quadratic functions		f(n) = n^2** \
	- Measure the cost of looking at most or all _pairs_ of items in an n-element universe

**Cubic functions			f(n) = n^3** \
	- Such functions enumerate through all triples of items in an n-element universe \
	- Also arise in certain dynamic programing algorithms

#### Anything else between quadratic and subexponential

Advice: don't run this on a million rows

#### O(b ^ n), O(n!): exponential and up

Advice: good luck if you have more than a dozen or two data points

**Exponential functions	f(n) = c^n** _// for a given constant c > 1_ \ 
	- Functions like 2n arise when enumerating all subsets of n items. As we have seen, exponential algorithms become useless fast, but not as fast as. . .

**Factorial functions		f(n) = n!** \
Arise when generating all permutations or orderings of n items

![](../assets/big_O_datastructure_operations.png)
![](../assets/big_O_array_sorting_algorithms.png)