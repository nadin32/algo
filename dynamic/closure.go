//https://www.geeksforgeeks.org/closures-in-golang/

// example of using closure for data isolation in dynmaic programming
// for memoization function (  caching intermediate results )
//
package dynamic

import "fmt"

func calc(n int) int {
	return n + 1
}
func MemoizationCalc() func(int) int {
	cachemap := make(map[int]int)
	fmt.Println("init cache once")
	return func(n int) int {
		if _, ok := cachemap[n]; !ok {
			fmt.Println("new cache entry:", n)
			cachemap[n] = calc(n)
		}
		return cachemap[n]
	}
}
