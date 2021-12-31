//https://www.geeksforgeeks.org/closures-in-golang/
// example of using closure for data isolation in dynmaic programming
// for memoization function (  caching intermediate results )
//
package dynamic

import (
	"fmt"
	"testing"
)

type calcTest struct {
	name   string
	input  int
	output int
}

var calcTestcases = []calcTest{
	calcTest{
		"first",
		1,
		2,
	},
    calcTest{
		"second",
		1,
		2,
	},
}

func TestMemoizationCalc(t *testing.T) {
	cm := MemoizationCalc()
	for _, test := range calcTestcases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			runtimeroutput := cm(test.input)
            if (test.output != runtimeroutput) {
				t.Errorf(" Expected: %d, Computed: %d", test.output, runtimeroutput)
			}
		})
	}
}
