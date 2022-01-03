package dynamic

/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Ref. https://www.geeksforgeeks.org/count-ways-reach-nth-stair-using-step-1-2-3/?ref=lbp
*/

import (
	"fmt"
	"testing"
)

type st struct {
	name   string
	steps  int
	output int
}

var stTestcases = []st{
	st{"first", 0, 0},
	st{"second", 1, 1},
	st{"third", 2, 2},
	st{"fourth", 3, 3},
}

var caches map[int]int
var caches2 []int

//DP :Top-Down Approach:
//To keep the recursive structure intact and just store
//the value in a HashMap and whenever the function is called again return the value store without computing ().
func steps(n int) int {
	caches[0] = 0
	caches[1] = 1
	caches[2] = 2
	if n < 3 {
		return caches[n]
	} else {
		_, ok := caches[n]
		if !ok {
			caches[n] = steps(n-1) + steps(n-2)
		}
	}
	return caches[n]
}

func stepsDP1(n int) func(n int) int {
	caches = make(map[int]int)
	return steps
}

func stepsDP2(n int) func(n int) int {
	caches2 = make([]int, n+1)
	return steps2
}

//DP Bottom-Up Approach: The second way is to take an extra space of size n and start computing values of states
// from 1, 2 .. to n, i.e.
// compute values of i, i+1  and then use them to calculate the value of i+2.
func steps2(n int) int {
	if n < 3 {
		return n
	}
	caches2[0] = 0
	caches2[1] = 1
	caches2[2] = 2
	for j := 3; j < n+1; j++ {
		caches2[j] = caches2[j-1] + caches2[j-2]
	}
	return caches2[n]
}

func TestStepsDP1(t *testing.T) {
	for _, test := range stTestcases {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
			f := stepsDP1(test.steps)
			ns := f(test.steps)
			if test.output != ns {
				t.Errorf(" Expected: %d, Computed: %d", test.output, ns)
			}
		})

	}
}

func TestStepsDP2(t *testing.T) {
	for _, test := range stTestcases {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
			f := stepsDP2(test.steps)
			ns := f(test.steps)
			if test.output != ns {
				t.Errorf(" Expected: %d, Computed: %d", test.output, ns)
			}
		})

	}
}
