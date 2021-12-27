

--------------------

## Algorithms.

BIG O sheet: https://www.bigocheatsheet.com/ 

## Sorting ##


Sorting animation:https://www.toptal.com/developers/sorting-algorithms 

|Sorting Algorithm|Average Case|Best Case|Worst Case|description| advantages| disadvantages|space complexity|
|-----------------|------------|---------|----------|-----------| ----------|-----------| -----------| 
|Bubble Sort |O(n^2)| O(n) | O(n^2)|repeatedly compares and swaps(if needed) adjacent elements in every pass. In i-th pass of Bubble Sort (ascending order), last (i-1) elements are already sorted, and i-th largest element is placed at (N-i)-th position, i.e. i-th last position. |1) It is the simplest sorting approach. 2) Best case complexity is of O(N) [for optimized approach] while the array is sorted.3) Using optimized approach, it can detect already sorted array in first pass with time complexity of O(N).4) Stable sort: does not change the relative order of elements with equal keys.In-Place sort.|is comparatively slower algorithm.|O(1)|
|Insertion Sort|O(n^2)| O(n) | O(n^2)|This is an in-place comparison-based sorting algorithm. A sub-list is maintained which is always sorted. For example, the lower part of an array is maintained to be sorted. An element which is to be 'insert'ed in this sorted sub-list, has to find its appropriate place and then it has to be inserted there.|1)It can be easily computed. 2) Best case complexity is of O(N) while the array is already sorted. 3) Number of swaps reduced than bubble sort. 4) For smaller values of N, insertion sort performs efficiently like other quadratic sorting algorithms. 5) Stable sort. 6) Adaptive: total number of steps is reduced for partially sorted array. 7) In-Place sort. 8) It is more efficient than the Selection sort.|It is generally used when the value of N is small. For larger values of N, it is inefficient. |O(1)| 
|Selection Sort|O(n^2)| O(n^2) | O(n^2)|It is an in-place comparison-based algorithm in which the list is divided into two parts, the sorted part at the left end and the unsorted part at the right end. Initially, the sorted part is empty and the unsorted part is the entire list.The smallest element is selected from the unsorted array and swapped with the leftmost element, and that element becomes a part of the sorted array. This process continues moving unsorted array boundary by one element to the right.| 1) It can   be used on list structures that make add and remove efficient, such as a linked list. Just remove the smallest element of unsorted part and end at the end of sorted part. 2)  The number of swaps reduced. O(N) swaps in all cases. 3) In-Place sort.|Time complexity in all cases is O(N^2), no best case scenario. This algorithm is not suitable for large data sets as its average and worst case complexities are of Ο(n2)  |O(1)| 
|Quick Sort|O(n.log(n))|O(n.log(n))|O(n^2)|Divide & Conquer. There are many different versions of quickSort that pick pivot in different ways. the key is PARTITION of the list around pivot point. Divide: 1) Select a pivot element that will preferably end up close to the center of the sorted pack 2) Move everything onto the “greater than” or “less than” side of the pivot 3)The pivot is now in its final position 4) Recursively repeat the operation on both sides of the pivot .Conquer: Return a sorted array after all elements have been through the pivot operation  | 1)Stable sort 2)When implemented well, it can be somewhat faster than merge sort and about two or three times faster than heapsort.|-----------|O(logn)| 
|Merge Sort|O(n.log(n))|O(n.log(n))|O(n.log(n))| Divide & Conquer. 1) Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted). 2) Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.  |----------|-----------|O(n)| 
|Heap Sort|O(n.log(n))|	O(n.log(n))|	O(n.log(n))||----------|-----------|O(1)| 
|Counting Sort|O(n+k)|	O(n+k)|	O(n+k)| | ----------|-----------|O(k)| 
|Radix Sort|O(n.k)|	O(n.k)|	O(n.k)|  |----------|-----------|O(n + k)| 
|Bucket Sort|O(n+k)|	O(n+k)|	O(n2) ||----------|-----------|O(n)| 
|-----------------|------------|---------|----------|  |----------|-----------|-----------| 
|-----------------|------------|---------|----------|  |----------|-----------|-----------| 
|-----------------|------------|---------|----------|  |----------|-----------|-----------| 
|-----------------|------------|---------|----------|  |----------|-----------|-----------| 

			


			

	

	

	 

	

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
 
<details>
           <summary><b>DataStructure</b></summary>
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
