<details>
           <summary>##DataStructures##</summary>
           <p>
		   [**Master the Interview Click here for Course Link** ](https://academy.zerotomastery.io/p/master-the-coding-interview-data-structures-algorithms)
		   
	🎁Hash Tables
		![image](attached://008d8d5fbe7759015b102fb5e61eedf7 125x150)             
		
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
					 ![image](attached://7df06ec4ad27f2f697ec682ddb5aed6c 125x150)
				🎁Doubly Linked List 
					
			🎁Binary Tree
				🎁Binary Search Tree
					🎁Balanced BST
						🎁AVL Tree
						🎁Red Black Tree
						![image](attached://1cbd9dae9e063cb5b0294abfc41bc600 125x150)
			🎁Heap
				🎁Binary Heap
					Priority Queue
					![image](attached://26b117268285c9e89ade0640e6ca4147 125x150)
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
			 ![image](attached://6aa87955d9d6283c9bd60c7f4e0c6d8c 125x150)
		Dynamic
			 ![image](attached://0f574b0c427c0cf20e9c03ccb546d71c 125x150)
			
			** can be O(n) on expanding memory
		Searching. Is it sorted?
			Yes - Divide and Conquer - **Binary search O(log N)** 🚀
			No. Will sorting make it faster? If still no, **Linear Search**🚀
			No. Is it a String? See if a **Trie** data structure helps
	🎁Stacks
		Array Stack
		Linked List Stack
		![image](attached://a03fa110a9697e07ac4979d9e8a61d57 125x150)
	🎁Queues
		Array Queue (BAD)
		Linked List Queue
		![image](attached://afc27bd2c74ddef65f4905665d511006 125x150)
	Dynamic Programming
		🚀Memoization</p>
	
</details>

--------------------

# Algorithms.

## Sorting ##


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

            
