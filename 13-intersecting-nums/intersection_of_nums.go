// package spot_the_diff
package main

import (
	"fmt"
)

func main() {
	testIntersectionOfNums()
}

func IntersectionOfNums(ipArr1 []int, ipArr2 []int) []int {
	arrMap := []int{}
	testMap := make(map[int]int)
	interMap := make(map[int]bool)
	for itr := range ipArr1 {
		testMap[ipArr1[itr]] += 1
	}
	// fmt.Println(ipArr2)
	for itr := range ipArr2 {
		if testMap[ipArr2[itr]] != 0 {
			interMap[ipArr2[itr]] = true
		}
	}
	for k := range interMap {
		arrMap = append(arrMap, k)
	}
	return arrMap
}

func testIntersectionOfNums() {
	testInputs1 := [][]int{{2, 4, 4, 2}, {1, 2, 3, 3}, {2, 4, 6, 8}}
	testInputs2 := [][]int{{2, 4}, {3, 3}, {1, 3, 5, 7}}
	expectedResults := [][]int{{2, 4}, {3}, {}}
	testResults := [][]int{}
	for testCaseIdx := range testInputs1 {
		// fmt.Println("IP1: ", testInputs1[testCaseIdx])
		// fmt.Println("IP2: ", testInputs2[testCaseIdx])
		// result := IntersectionOfNums(testInputs1[testCaseIdx], testInputs2[testCaseIdx])
		// fmt.Println("Result: ", result)
		testResults = append(testResults, IntersectionOfNums(testInputs1[testCaseIdx], testInputs2[testCaseIdx]))
	}
	for i := range testResults {
		if len(testResults[i]) == len(expectedResults[i]) {
			res := true
			for itr := range testResults[i] {
				if testResults[i][itr] == expectedResults[i][itr] {
					continue
				} else {
					res = false
				}
			}
			if res == false {
				fmt.Printf("Input 1:%d, Input 2:%d, Output:%d, Expected:%d\n", testInputs1[i], testInputs2[i], testResults[i], expectedResults[i])
				fmt.Printf("Failing %d!!\n", i+1)
				break
			}
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input 1:%d, Input 2:%d, Output:%d, Expected:%d\n", testInputs1[i], testInputs2[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
