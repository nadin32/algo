<details>
           <summary>Title 1</summary>
           <p>Content 1 Content 1 Content 1 Content 1 Content 1</p>
         </details>
# Algorythms.

## Searches ##

### BFS  ## 

Algorithm BFS(G, v)

    Q ← new empty FIFO queue
    
    Mark v as visited.
    
    Q.enqueue(v)
    
    while Q is not empty
    
        a ← Q.dequeue()
        
        // Perform some operation on a.
        
        for all unvisited neighbors x of a
        
            Mark x as visited.
            
            Q.enqueue(x)
            
