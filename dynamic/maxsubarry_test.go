package dynamic

/* The maximum subarray problem is the task of finding the contiguous subarray within a one-dimensional array of numbers which has the largest sum.

Example:

int [] A = {−2, 1, −3, 4, −1, 2, 1, −5, 4};
Output: contiguous subarray with the largest sum is 4, −1, 2, 1, with sum 6. 

Bottom-Up approach
“MS(i)  is maximum sum ending at index i”
Maximum Subarray ProblemTo calculate the solution for any element at index “i” has two options

    EITHER added to the solution found till “i-1“th index
    OR start a new sum from the index “i“.

Recursive Solution:

MS(i) = Max[MS(i-1) + A[i] , A[i]]
*/
import (
	"fmt"
	"testing"
)

type sr struct {
	name string
	m    []int
	ms   int
}

var srTestcases = []sr{
	sr{"first", []int{4, -1, 2, 1}, 6},
	sr{"second", []int{2, -7, 9, 3, -1}, 12},
	sr{"third", []int{0, -7, 0, 0, -1}, 0},
}

// sum[i] = max sum ending at index i
// sum[1] = max(sum[0]+ a[1]),a[1])
func maxSubArray(a []int) int {
	sum := make([]int, len(a))
	sum[0] = a[0]
	for j := 1; j < len(sum); j++ {
		s1 := sum[j-1] + a[j]
		if s1 > a[j] {
			sum[j] = s1
		} else {
			sum[j] = a[j]
		}
	}
    return max(sum)
}

//max of the array
func max(a []int) int {
    m:= a[0]
    for _,j:= range a{
		if j > m {
			m =j
		}
	}
	return m
}
func TestMaxSubArray(t *testing.T) {
	for _, test := range srTestcases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			fmt.Println(test.name)
			sm := maxSubArray(test.m)
			if sm != test.ms {
				t.Errorf(" Expected: %d, Computed: %d", test.ms, sm)
			}
		})

	}

}
