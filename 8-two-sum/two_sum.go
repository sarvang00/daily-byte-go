package main

import (
	"fmt"
)

func main() {
	testTwoSum()
}

func TwoSum(ipVals []int, k int) bool {
	for i := 0; i < len(ipVals); i++ {
		for j := i + 1; j < len(ipVals); j++ {
			if ipVals[i]+ipVals[j] == k {
				return true
			}
		}
	}
	return false
}

func testTwoSum() {
	testVals := [][]int{{1, 3, 8, 2}, {3, 9, 13, 7}, {4, 2, 6, 5, 2}}
	testKs := []int{10, 8, 4}
	expectedResults := []bool{true, false, true}
	testResults := []bool{}
	for testCaseIdx := range testKs {
		testResults = append(testResults, TwoSum(testVals[testCaseIdx], testKs[testCaseIdx]))
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
