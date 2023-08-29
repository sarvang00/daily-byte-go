package main

import (
	"fmt"
)

func main() {
	testFirstUniqueChar()
}

func FirstUniqueChar(ipString string) int {
	if len(ipString) == 0 {
		return -1
	} else if len(ipString) == 1 {
		return 0
	} else {
		posMap := make(map[string]int)
		countMap := make(map[string]int)

		for charItr := range ipString {
			if countMap[ipString[charItr:charItr+1]] == 0 {
				countMap[ipString[charItr:charItr+1]] += 1
				posMap[ipString[charItr:charItr+1]] = charItr
			} else {
				countMap[ipString[charItr:charItr+1]] += 1
			}
		}

		var valuableResults []string
		for k, v := range countMap {
			if v == 1 {
				valuableResults = append(valuableResults, k)
			}
		}

		if len(valuableResults) == 0 {
			return -1
		} else {
			minPos := posMap[valuableResults[0]]
			for resultItr := range valuableResults {
				if posMap[valuableResults[resultItr]] < minPos {
					minPos = posMap[valuableResults[resultItr]]
				} else {
					continue
				}
			}
			return minPos
		}
	}
}

func testFirstUniqueChar() {
	testInputs := []string{"abcabd", "thedailybyte", "developer", "aa", "", "z"}
	expectedResults := []int{2, 1, 0, -1, -1, 0}
	testResults := []int{}
	for testCaseIdx := range testInputs {
		testResults = append(testResults, FirstUniqueChar(testInputs[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Inputs:%s, Output:%d, Expected:%d\n", testInputs[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
