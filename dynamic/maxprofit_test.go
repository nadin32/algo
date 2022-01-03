package dynamic

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future
to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

*/

import (
	"fmt"
	"testing"
)

type mp struct {
	name   string
	input  []int
	output int
}

var mptTestcases = []mp{
	mp{"first", []int{7, 1, 5, 3, 6, 4}, 5},
	mp{"second", []int{7, 6, 4, 3, 1}, 0},
}

// maxprofit[i] reflects  max profit  achieved so far  by i-day.
// it is max off {maxprpfit[i-1], (price[i] -minprice)}
// price : {7, 1, 5, 3, 6, 4}
//
// i=0  minprice =7 maxprpofit{0,0,0,0,0,0}
// i=1  minprice =1 maxprofit{0,0,0,0,0,0}
// i=2  minprice =1 maxprofit{0,0,4,0,0,0}
// i=3  minprice =1 maxprofit{0,0,4,4,0,0}
// i=4  minprice =1 maxprofit{0,0,4,4,5,0}
// i=5  minprice =1 maxprofit{0,0,4,4,5,5}
//
func maxprofitDP(price []int) int {
	if len(price) < 2 {
		return 0
	}
	maxprofit := make([]int, len(price))
	minprice := price[0]
	maxprofit[0] = 0
	for i := 1; i < len(price); i++ {
		if price[i] < minprice {
			minprice = price[i]
		}
		t := price[i] - minprice
		if t > maxprofit[i-1] {
			maxprofit[i] = t
		} else {
			maxprofit[i] = maxprofit[i-1]
		}

	}
	return maxprofit[len(maxprofit)-1]
}

func TestMaxpropfitDP(t *testing.T) {
	for _, test := range mptTestcases {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
			ns := maxprofitDP(test.input)
			if test.output != ns {
				t.Errorf(" Expected: %d, Computed: %d", test.output, ns)
			}
		})

	}
}
