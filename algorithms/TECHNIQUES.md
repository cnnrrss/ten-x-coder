# Algorithm Techniques

[Two pointers](#Two-pointer-Technique)
Sliding Window
Recursion
Iteration - Queue, Stack, Array
Search - Binary, DFS, BFS, Tree traversal
[Dynamic Programming](#Dynamic-Programming)
[Backtracking](#Backtracking)
[Sorting - Quick sort, Merge Sort](#Sorting)
[Divide and Conquer](#Divide-and-conquer)

## Two-pointer Technique
Array and Linked List. Array is fundamental blocks in algorithms (String is just array of characters).
Two pointers:
- Slow-runner and fast-runner
- Classic example is to remove duplicates from a sorted array.
- Another variation is one pointer at beginning, other p at the end moving toward each other until the both meet. (Reverse characters in a string)
**Classic problems**: Remove Duplicates from Sorted Array, Two Sum II - Input array is sorted, Reverse Words in a String II, Rotate Array, Valid Palindrome, Container with Most Water, Product of Array Except Self

## Divide and conquer
- Sorting
- Binary Search
- Fast Heap Construction
- Matrix Multiplication

## Union Find
- Propogation
- Connectivity

## Dynamic Programming
- Pascals triangle
- Approx. String matching
- Longest Common Subsequence
- Longest Increasing Sequence
- Coin change
- Rod cutting
- Class 

**Strategies** \
Tabulation \
Recursion \
Memoization vs Top Down vs Bottom Up \
Greedy Algorithm


## Searching
- Binary Search

## Backtracking
**Problems**
- All subsets
- All permutations
- All paths in graph
- Search pruning
- Sudoku
- Knight's tour problem
- Rat in a maze
	- Rat in a maze with multiple steps or jumps allowed
- N Queen
	- O(n) space
- Subset Sum
- m Coloring
- Hamiltonian Cycle
- Magnet Puzzle

## Sorting
- Merge Sort O(n log(n))
- Quick Sort O(n log(n))
- Other O(n) sorting (insertion sort, bubble sort, selection sort)
- Bucket sort
- Convex hulls
- Increasing or decreasing order
- What to do with equal keys

## Bit Manipulation

## Greedy Method
For example, a greedy strategy for the traveling salesman problem (which is of a high computational complexity) is the following heuristic: "At each step of the journey, visit the nearest unvisited city." This heuristic does not intend to find a best solution, but it terminates in a reasonable number of steps; finding an optimal solution to such a complex problem typically requires unreasonably many steps. 

In mathematical optimization, greedy algorithms optimally solve combinatorial problems having the properties of matroids, and give constant-factor approximations to optimization problems with submodular structure.

Greedy method algorithms have five components:
1. A candidate set, from which a solution is created
2. A selection function, which chooses the best candidate to be added to the solution
3. A feasibility function, that is used to determine if a candidate can be used to contribute to a solution
4. An objective function, which assigns a value to a solution, or a partial solution, and
5. A solution function, which will indicate when we have discovered a complete solution


The activity selection problem is characteristic to this class of problems, where the goal is to pick the maximum number of activities that do not clash with each other.

In the Macintosh computer game Crystal Quest the objective is to collect crystals, in a fashion similar to the travelling salesman problem. The game has a demo mode, where the game uses a greedy algorithm to go to every crystal. The artificial intelligence does not account for obstacles, so the demo mode often ends quickly.

The matching pursuit is an example of greedy algorithm applied on signal approximation.

A greedy algorithm finds the optimal solution to Malfatti's problem of finding three disjoint circles within a given triangle that maximize the total area of the circles; it is conjectured that the same greedy algorithm is optimal for any number of circles.

A greedy algorithm is used to construct a Huffman tree during Huffman coding where it finds an optimal solution.

In decision tree learning, greedy algorithms are commonly used, however they are not guaranteed to find the optimal solution.

One popular such algorithm is the ID3 algorithm for decision tree construction.

Dijkstra's algorithm and the related A* search algorithm are verifiably optimal greedy algorithms for graph search and shortest path finding.

A* search is conditionally optimal, requiring an "admissible heuristic" that will not overestimate path costs.

Kruskal's algorithm and Prim's algorithm are greedy algorithms for constructing minimum spanning trees of a given connected graph. They always find an optimal solution, which may not be unique in general.

https://github.com/atambol/Leetcode/blob/master/Notes.md

### Heaps Algorithm
- Heapsort

- All permutations of a string
- Ex: aaB -> [ 'aaB', 'aaB', 'Baa', 'aBa', 'aBa', 'Baa' ]
- Problem to find only the permutations where no matching neighbor chars
- Ex: aaB -> [ 'aBa', 'aBa' ]

- Use regular expression to match repeated consecutive characters.
- The string str is split into an array of characters, arr.
- 0 is returned if str contains same characters.
- The function swap() is used for the purpose of swapping the contents of two variable’s contents.
- Heap’s algorithm to generate arrays of permutations in permutations.
- Filter permutations to include only permutations with non-repeated characters.
- Filtered.length returns the number of total permutations of the provided string that don’t have repeated consecutive letters. ANSWER