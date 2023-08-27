package two_sum

import "fmt"

// import (
// 	"fmt"
// )

func main() {
	testTwoSumOptimal()
}

func TwoSumOptimal(ipVals []int, k int) bool {
	var diffMap = make(map[int]int)
	for iter := range ipVals {
		diff := k - ipVals[iter]
		_, diffCheck := diffMap[diff]
		if diffCheck {
			return true
		} else {
			diffMap[ipVals[iter]] = ipVals[iter]
		}
	}
	return false
}

func testTwoSumOptimal() {
	testVals := [][]int{{1, 3, 8, 2}, {3, 9, 13, 7}, {4, 2, 6, 5, 2}}
	testKs := []int{10, 8, 4}
	expectedResults := []bool{true, false, true}
	testResults := []bool{}
	for testCaseIdx := range testKs {
		testResults = append(testResults, TwoSumOptimal(testVals[testCaseIdx], testKs[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input:%d, K:%d, Output:\"%t\"\n", testVals[i], testKs[i], testResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
