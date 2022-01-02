package dynamic

// https://leetcode.com/problems/house-robber/
/* You are a professional robber planning to rob houses along a street.
 Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is
  that adjacent houses have security systems connected and it will automatically contact the police if two adjacent
  houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount
of money you can rob tonight without alerting the police. */
/* Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:

Input: nums = [2,7,9,1,1,3]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 6 (money = 3).
Total amount you can rob = 2 + 9 + 3 = 12.

https://www.geeksforgeeks.org/find-maximum-possible-stolen-value-houses/
*/
import (
	"fmt"
	"testing"
)

type hr struct {
	name string
	m    []int
	ms   int
}

var hrTestcases = []hr{
	hr{"first", []int{1, 2, 3, 1}, 4},
	hr{"second", []int{2, 7, 9, 3, 1}, 12},
	hr{"third", []int{5, 3, 4, 11, 2}, 16},
}

//to cover base cases
func houseRob(h []int) int {
	if len(h) == 0 {
		return 0
	}
	if len(h) == 1 {
		return h[0]
	}

	return houseRobber(h)
}

// sum [i] - is the max ammount the robber can steal if stopping and stealing  at the house i
func houseRobber(a []int) int {
	sum := make([]int, len(a))
	sum[0] = a[0]
	sum[1] = maxa(a[:2])
	for j := 2; j < len(sum); j++ {
		s1 := maxa([]int{sum[j-2] + a[j], sum[j-1]})
		sum[j] = s1
	}
	return maxa(sum)
}

//max of the array
func maxa(a []int) int {
	m := a[0]
	for _, j := range a {
		if j > m {
			m = j
		}
	}
	return m
}

func TestHouseRobber(t *testing.T) {

	for _, test := range hrTestcases {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
			rms := houseRob(test.m)
			if test.ms != rms {
				t.Errorf(" Expected: %d, Computed: %d", test.ms, rms)
			}
		})

	}
}
