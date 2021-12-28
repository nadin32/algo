

--------------------

## Algorithms.

BIG O sheet: https://www.bigocheatsheet.com/ 

## Sorting ##


Sorting animation:https://www.toptal.com/developers/sorting-algorithms 

|Sorting Algorithm|Average Case|Best Case|Worst Case|description| advantages| disadvantages|space complexity|
|-----------------|------------|---------|----------|-----------| ----------|-----------| -----------| 
|Bubble Sort |O(n^2)| O(n) | O(n^2)|repeatedly compares and swaps(if needed) adjacent elements in every pass. In i-th pass of Bubble Sort (ascending order), last (i-1) elements are already sorted, and i-th largest element is placed at (N-i)-th position, i.e. i-th last position. |1) It is the simplest sorting approach. 2) Best case complexity is of O(N) [for optimized approach] while the array is sorted.3) Using optimized approach, it can detect already sorted array in first pass with time complexity of O(N).4) Stable sort: does not change the relative order of elements with equal keys.In-Place sort.|is comparatively slower algorithm.|O(1)|
|Insertion Sort|O(n^2)| O(n) | O(n^2)|Insertion sort works similarly as we sort cards in our hand in a card game.This is an in-place comparison-based sorting algorithm. A sub-list is maintained which is always sorted. For example, the lower part of an array is maintained to be sorted. An element which is to be 'insert'ed in this sorted sub-list, has to find its appropriate place and then it has to be inserted there.|1)It can be easily computed. 2) Best case complexity is of O(N) while the array is already sorted. 3) Number of swaps reduced than bubble sort. 4) For smaller values of N, insertion sort performs efficiently like other quadratic sorting algorithms. 5) Stable sort. 6) Adaptive: total number of steps is reduced for partially sorted array. 7) In-Place sort. 8) It is more efficient than the Selection sort.|It is generally used when the value of N is small. For larger values of N, it is inefficient. |O(1)| 
|Selection Sort|O(n^2)| O(n^2) | O(n^2)|It is an in-place comparison-based algorithm in which the list is divided into two parts, the sorted part at the left end and the unsorted part at the right end. Initially, the sorted part is empty and the unsorted part is the entire list.The smallest element is selected from the unsorted array and swapped with the leftmost element, and that element becomes a part of the sorted array. This process continues moving unsorted array boundary by one element to the right.| 1) It can   be used on list structures that make add and remove efficient, such as a linked list. Just remove the smallest element of unsorted part and end at the end of sorted part. 2)  The number of swaps reduced. O(N) swaps in all cases. 3) In-Place sort.|Time complexity in all cases is O(N^2), no best case scenario. This algorithm is not suitable for large data sets as its average and worst case complexities are of Ο(n2)  |O(1)| 
|Quick Sort|O(n.log(n))|O(n.log(n))|O(n^2)|Divide & Conquer. There are many different versions of quickSort that pick pivot in different ways. the key is PARTITION of the list around pivot point. Divide: 1) Select a pivot element that will preferably end up close to the center of the sorted pack 2) Move everything onto the “greater than” or “less than” side of the pivot 3)The pivot is now in its final position 4) Recursively repeat the operation on both sides of the pivot .Conquer: Return a sorted array after all elements have been through the pivot operation  | 1)Stable sort 2)When implemented well, it can be somewhat faster than merge sort and about two or three times faster than heapsort.|-----------|O(logn)| 
|Merge Sort|O(n.log(n))|O(n.log(n))|O(n.log(n))| Divide & Conquer. 1) Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted). 2) Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.  |----------|   1) Slower comparative to the other sort algorithms for smaller tasks. 2)  Merge sort algorithm requires an additional memory space of 0(n) for the temporary array. 3)  It goes through the whole process even if the array is sorted.|O(n)| 
|Heap Sort|O(n.log(n))|	O(n.log(n))|	O(n.log(n))|(see descr.below under BinaryHeap DS) Heapsort is a comparison-based sorting algorithm. Heap sort works by presenting  the elements of the array as a special kind of complete binary tree called a heap.Heapsortcan be thought of as an improved selection sort: like selection sort,heapsort divides its input into a sorted and an unsorted region, and ititeratively shrinks the unsorted region by extracting the largest elementfrom it and inserting it into the sorted region. Unlike selection sort,heapsort does not waste time with a linear-time scan of the unsorted region;rather, heap sort maintains the unsorted region in a heap data structureto more quickly find the largest element in each step. |----------|-----------|O(1)| 
|Counting Sort|O(n+k) n-max of the elments k-array size |	O(n+k)|	O(n+k)| It sorts the elements of an array by counting the number of occurrences of each unique element in the array. The count is stored in an auxiliary array and the sorting is done by mapping the count as an index of the auxiliary array.|There is no comparison between any elements, so it is better than comparison based sorting techniques.  1) Counting sort generally performs faster than all comparison-based sorting algorithms, such as merge sort and quicksort, if the range of input is of the order of the number of input 2)Counting sort is easy to code 3)stabele  sort. |  1) Counting sort doesn’t work on decimal values 2) Counting sort is inefficient if the range of values to be sorted is very large|O(k)| 
|Radix Sort| O(n.k) |	O(n.k) |	O(n.k) |The idea of Radix Sort is to do digit by digit sort starting from least significant digit to most significant digit. Radix sort uses counting sort as a subroutine to sort. |Radix sort has linear time complexity which is better than O(nlog n) of comparative sorting algorithms.|If we take very large digit numbers or the number of other bases like 32-bit and 64-bit numbers then it can perform in linear time however the intermediate sort takes large space. This makes radix sort space inefficient. This is the reason why this sort is not used in software libraries.|O(n + k)| 
|Bucket Sort| O(n+k) |	O(n+k) |	O(n2)  |Bucket sort is a sorting algorithm that separate the elements into multiple groups said to be buckets. Elements in bucket sort are first uniformly divided into groups called buckets, and then they are sorted by any other sorting algorithm.|1) Bucket sort is stable, if the underlying sort is also stable, as equal keys are inserted in order to each bucket.2)Bucket sort allows each bucket to be processed independently. As a result, you'll frequently need to sort much smaller arrays as a secondary step after sorting the main array. 3) Bucket sort to be used as an external sorting algorithm. If you need to sort a list that is too large to fit in memory, you may stream it through RAM, split the contents into buckets saved in external files, and then sort each file separately in RAM. | worst case occurs when the elements are of the close range in the array, because of that, they have to be placed in the same bucket. So, some buckets have more number of elements than others.The complexity will get worse when the elements are in the reverse order.|O(n.k)| 
|-----------------|------------|---------|----------|----------|-----------|-----------|-----------| 
|-----------------|------------|---------|----------|----------|-----------|-----------|-----------|
|-----------------|------------|---------|----------|----------|-----------|-----------|-----------|
|-----------------|------------|---------|----------|----------|-----------|-----------|-----------|


	

	 
**Counting sort**

Counting sort is handy while sorting values whose range is not very large.

good illustration : https://www.cs.bgu.ac.il/~ds142/wiki.files/ds142_ps11_updated.pdf
links (pseudo algo.)  : https://www.programiz.com/dsa/counting-sort

The Counting sort typically has 3 steps
```
Input Data : 1, 4, 1, 2, 7, 5, 2

    1. Create Count array to store count of each individual object
     So for the above input data, 1 appears twice so in count array, index 1 stores value 2, for 4, index 4 stores 1 and so on..
     So Count Array becomes —
	index   0 1 2 3 4 5 6 7 8 9
	count   0 2 2 0 1 1 0 1 0 0

    2.  Modify the count array to cumulate the sum with increasing index.
 	 So count at index 1 becomes (0+2)=2, at index 2 becomes (2+2)=4 and so on…
 	 index   0 1 2 3 4 5 6 7 8 9
	 count   0 2 4 4 5 6 6 7 7 7

	This modified count array represents the position of each object.

     3. Now use the original data with count array to determine position of each object in sorted array.
	 So for 1 (original data), at index 1 (in count array) the value is 2, so place 1 at index 2 in sorted array.
	 After every determination reduce the count in count array by 1.
	 So, value at index array after determining position of 1 (original data) becomes 1.
	 Similarly place 4 at position 5, and reduce the value by 1.
	
```

***Radix sort**
What if the elements are in the range from 1 to n2? 
We can’t use counting sort because counting sort will take O(n2) which is worse than comparison-based sorting algorithms. Can we sort such an array in linear time? 

```
The Radix Sort Algorithm 

    Do following for each digit i where i varies from least significant digit to the most significant digit.
        Sort input array using counting sort (or any stable sort) according to the i’th digit.

Example:

Original, unsorted list:
170, 45, 75, 90, 802, 24, 2, 66

Sorting by least significant digit (1s place) gives: 
[*Notice that we keep 802 before 2, because 802 occurred 
before 2 in the original list, and similarly for pairs 
170 & 90 and 45 & 75.]

170, 90, 802, 2, 24, 45, 75, 66

Sorting by next digit (10s place) gives: 
[*Notice that 802 again comes before 2 as 802 comes before 
2 in the previous list.]

802, 2, 24, 45, 66, 170, 75, 90

Sorting by the most significant digit (100s place) gives:
2, 24, 45, 66, 75, 90, 170, 802
```
***Bucket sort**
```
Bucket sort is mainly useful when input is uniformly distributed over a range. For example, consider the following problem. 
Sort a large set of floating point numbers which are in range from 0.0 to 1.0 and are uniformly distributed across the range. How do we sort the numbers efficiently?
A simple way is to apply a comparison based sorting algorithm. The lower bound for Comparison based sorting algorithm (Merge Sort, Heap Sort, Quick-Sort .. etc) is Ω(n Log n), i.e., they cannot do better than nLogn. 
Can we sort the array in linear time? Counting sort can not be applied here as we use keys as index in counting sort. Here keys are floating point numbers. 

Because of the way elements are assigned to buckets, typically using an array where the index is the value, bucket sorting can be extremely fast.
This means that more auxiliary memory for the buckets is required at the expense of running time than more comparison sorts.
By dividing data into small buckets that can be sorted individually, the number of comparisons that must be performed is reduced.



bucketSort(arr[], n)
1) Create n empty buckets (Or lists).
2) Do following for every array element arr[i].
.......a) Insert arr[i] into bucket[n*array[i]]
3) Sort individual buckets using insertion sort.
4) Concatenate all sorted buckets.

For example: Suppose we have the following array, which we have to sort.
11	8	6	1	25	2

Step 1: Create a new array, and this array size is 10. Every field of this array is used as a bucket.  
5 backets: 1-5 6-10 11-15 16-20 21-25 

Step 2: Insert the array element into the new array bucket, and these elements will be added according to the range of the bucket.  
Then sort elemnst in ech backet using insertion sort. 
1,2 => backet 1-5
6,8 => backet 6-10
11 => backet 11-15
25 => backet 21-25
Step 3: Gather the element form each bucket. 
1,2,6,8,11,25
```
Sorting lecture notes:
https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-spring-2020/lecture-notes/MIT6_006S20_lec5.pdf

## Searches ##

### 1.Linear search  ##

Sequenatial search of locating the element in the list.
A simple approach is to do a linear search, i.e  find lemen xin in the list.

    Start from the leftmost element of list[] and one by one compare x with each element of list[]
    If x matches with an element, return the index.
    If x doesn’t match with any of elements, return -1.
    
Performance.
O(n) worst case.

Linear search is rarely used practically because other search algorithms such as the binary search algorithm and hash tables allow significantly faster-searching comparison to Linear search.

### 2.Binary search  ##



### 3.BFS  ##

Links: https://www.khanacademy.org/computing/computer-science/algorithms/breadth-first-search/a/the-breadth-first-search-algorithm

       https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-spring-2020/lecture-notes/MIT6_006S20_lec9.pdf
       
       https://www.programiz.com/dsa/graph-bfs
 
       https://github.com/maximelamure/algorithms/blob/master/datastructure/graph_bfs.go


Breadth–first search (BFS) is a graph traversal algorithm that explores vertices in the

order of their distance from the source vertex,where distance is the minimum length of a path from the source vertex to the node.

**BFS Algorithm Complexity**

The time complexity of the BFS algorithm is represented in the form of O(V + E), where V is the number of nodes and E is the number of edges.

The space complexity of the algorithm is O(V).
 mark and enqueue all (unvisited) neighbours of u
```
Algorithm BFS(G, v)
    create Q ← new empty FIFO queue
    Mark v as visited .
    Q.enqueue(v)   //put v in into Q     
    while Q is not empty
        u ← Q.dequeue() //remove the head  u of Q
        // Perform some operation on u.
        for all unvisited neighbors x of u
            Mark x as visited. 
           Q.enqueue(x) // put x into Q
            
```

**4.DepthFirst Search** 
            

**5.Dijkstra**

Applied for weighthed graph (with positive weights)

Links : https://medium.com/basecs/finding-the-shortest-path-with-a-little-help-from-dijkstra-613149fbdc8e
        https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/

Dijkstra's Algorithm works on the basis that any subpath B -> D of the shortest path A -> D between vertices A and D is also the shortest path between vertices B and D.

***Dijkstra's Algorithm Complexity***

Time Complexity: O(E Log V) where, E is the number of edges and V is the number of vertices.

Space Complexity: O(V)


***Algorithm*** 

1) Create a set sptSet (shortest path tree set) that keeps track of vertices included in the shortest-path tree, i.e., whose minimum distance from the source is calculated and finalized. Initially, this set is empty. 

2) Assign a distance value to all vertices in the input graph. Initialize all distance values as INFINITE. Assign distance value as 0 for the source vertex so that it is picked first. 

3) While sptSet doesn’t include all vertices 
….a) Pick a vertex u which is not there in sptSet and has a minimum distance value. 
….b) Include u to sptSet. 
….c) Update distance value of all adjacent vertices of u. To update the distance values, iterate through all adjacent vertices. For every adjacent vertex v, if the sum of distance value of u (from source) and weight of edge u-v, is less than the distance value of v, then update the distance value of v. 

-------------------
## Detailed Data Structures ##

### Binary Heap ###

Links: https://www.programiz.com/dsa/heap-sort.

We use array as underlying struictire for binary heaps.

**Heap Data Structure**

A ***complete binary tree*** is a binary tree in which all the levels are completely filled except possibly the lowest one, which is filled from the left.

Heap is is a complete binary tree .

All nodes in the tree follow the property that they are greater than their children i.e. the largest element is at the root and both its children and smaller than the root and so on. Such a heap is called a max-heap. If instead, all nodes are smaller than their children, it is called a min-heap

**Relationship between Array Indexes and Tree Elements**

A complete binary tree has an interesting property that we can use to find the children and parents of any node.

If the index of any element in the array is i, the element in the index 2i+1 will become the left child and element in 2i+2 index will become the right child. Also, the parent of any element at index i is given by the lower bound of (i-1)/2.

**How to "heapify" a tree**

Starting from a complete binary tree, we can modify it to become a Max-Heap by running a function called heapify on all the non-leaf elements of the heap.
```
void heapify(int arr[], int n, int i) {
  // Find largest among root, left child and right child 
  int largest = i;  
  int left = 2 * i + 1; 
  int right = 2 * i + 2
  if (left < n && arr[left] > arr[largest])
    largest = left;

  if (right < n && arr[right] > arr[largest])
    largest = right;
    // Swap and continue heapifying if root is not largest
    if (largest != i) {
      swap(&arr[i], &arr[largest]);
      heapify(arr, n, largest);
  }
}
```
**Build max-heap**

To build a max-heap from any tree, we can thus start heapifying each sub-tree from the bottom up and end up with a max-heap after the function is applied to all the elements including the root element.
```
   // Build heap (rearrange array)
    for (int i = n / 2 - 1; i >= 0; i--)
      heapify(arr, n, i);
```
**Use binary heap for heap sort**

Build Max heap (rearrange array elements using heapify () function )
1. Since the tree satisfies Max-Heap property, then the largest item is stored at the root node.
2.  Swap: Remove the root element and put at the end of the array (nth position) Put the last item of the tree (heap) at the vacant place.
3. Remove: Reduce the size of the heap by 1.
4. Heapify: Heapify the root element again so that we have the highest element at root.
5. The process is repeated until all the items of the list are sorted.
 
------------------- 
 
<details>
           <summary><b>Summary DataStructure</b></summary>
           <p>
		   [**Master the Interview Click here for Course Link** ](https://academy.zerotomastery.io/p/master-the-coding-interview-data-structures-algorithms)
		   
	🎁Hash Tables
		
		| space |   O(n)
		| insert|  O(1)
		| lookup|  O(1)
		| delete|  O(1) 
		   
		***could be O(n) with hash collisions and dynamic array resizing but unlikely*
		
		Improve Time Complexity?
			Fast Access O(1), tradeoff: more memory O(n) 
		Collision?
			Linked List
	🎁Graphs
		Shortest Path?
			🚀 Bellman-Ford
			🚀 Dijkstra
		Cyclic or Acyclic?
		Weighted or Unweighted?
		Graph Traversal? **O(n)**
			🚀 Breadth First Search (BFS)
			🚀 Depth First Search (DFS)
				Inorder
				Postorder
				Preorder
		Directed or Undirected?
		🎁Tree
			 Tree Traversal? **O(n)**
				
				
				**Recursion?**
					Be mindful of Space Complexity! (Stack overflow)
			🎁Linked List
				🎁Singly Linked List
					| prepend|  O(1)
		    			| append |  O(1)
					| lookup |  O(n)
					| insert |  O(n)
					| delete |  O(n) 
				🎁Doubly Linked List 
					
			🎁Binary Tree
				🎁Binary Search Tree
					🎁Balanced BST
						🎁AVL Tree
						🎁Red Black Tree				 	    				
						| lookup |  O(logN)
						| insert |  O(logN)
						| delete |  O(logN) 
			🎁Heap
				🎁Binary Heap
		 			| lookup |  O(N)
					| insert |  O(logN)
					| delete |  O(logN) 
				Priority Queue

			🎁Trie
	🎁Arrays
		Sorting? ~ **O(N log N)**
			Radix Sort 🚀
			Quick Sort 🚀
			Heap Sort 🚀
			Bubble Sort 🚀
			Selection Sort 🚀
			Insertion Sort 🚀
			Merge Sort 🚀
			Counting Sort 🚀
		String question? 
			Turn it into an Array ~ split() 🚀
		Static
			 | lookup |    O(1)
			 | push   |    O(1)
			 | insert |    O(n)
		         | delete |    O(n) 
		   
		Dynamic
			 | lookup |    O(1)
			 | append*|    O(1)
			 | insert |    O(n)
		         | delete |    O(n) 
		
			** can be O(n) on expanding memory
	        Searching. Is it sorted?
			Yes - Divide and Conquer - **Binary search O(log N)** 🚀
			No. Will sorting make it faster? If still no, **Linear Search**🚀
			No. Is it a String? See if a **Trie** data structure helps
	🎁Stacks
		Array Stack
		Linked List Stack
		         | lookup |    O(n)
			 | pop    |    O(1)
			 | push   |    O(1)
		         | peek   |    O(1) 
	🎁Queues
		Array Queue (BAD)
		Linked List Queue
			 | lookup |    O(n)
			 | enqueue|    O(1)
			 | dequeue|    O(1)
		         | peek   |    O(1) 
	Dynamic Programming
		🚀Memoization</p>
	
</details> 






