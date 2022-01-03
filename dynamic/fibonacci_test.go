package dynamic

// 0,1,1,2,3,5,8,13,21,34,
import (
	"fmt"
	"testing"
)

type ft struct {
	name   string
	input  int
	output int
}

var fibTestcases = []ft{
	ft{"first", 0, 0},
	ft{"second", 1, 1},
	ft{"third", 6, 8},
	ft{"fourth", 9, 34},
}

// cache for dynamic version
var cache map[int]int
var count int

//O(2^n)
func fibRecursive(ind int) int {
	count++
	if ind < 2 {
		return ind
	}

	return fibRecursive(ind-1) + fibRecursive(ind-2)
}

//O(n)
func fibIterative(k int) int {
	s := make([]int, k)
	s[0] = 0
	s[1] = 1
	sum := 1
	for j := 2; j < k+1; j++ {
		sum = s[0] + s[1]
		s[0] = s[1]
		s[1] = sum
	}
	return s[1]
}

//O(n)
//dynamic programming

func f(n int) int {
	count++
	if v, ok := cache[n]; ok {
		return v
	}
	if n < 2 {
		cache[n] = n
		return n
	} else {
		cache[n] = f(n-1) + f(n-2)
		return cache[n]
	}
}

func MemoizationFib() func(n int) int {
	cache = make(map[int]int)
	return f
}

func TestMemoizationFib(t *testing.T) {

	for _, test := range fibTestcases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			count = 0
			cm := MemoizationFib()
			runtimeroutput := cm(test.input)
			fmt.Println("called fb dynamic :", count, " times")
			if test.output != runtimeroutput {
				t.Errorf(" Expected: %d, Computed: %d", test.output, runtimeroutput)
			}
		})
	}
}

func TestRecursiveFib(t *testing.T) {
	for _, test := range fibTestcases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			count = 0
			runtimeroutput := fibRecursive(test.input)
			fmt.Println("called fb recursive  :", count, " times")
			if test.output != runtimeroutput {
				t.Errorf(" Expected: %d, Computed: %d", test.output, runtimeroutput)
			}
		})
	}
}
