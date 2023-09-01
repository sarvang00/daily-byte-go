package uncommon_words

// package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	testUncommonWords()
// }

func UncommonWords(ipStrs1 string, ipStrs2 string) []string {
	resMap := []string{}
	wordCount := make(map[string]int)

	strList := strings.Split(ipStrs1, " ")
	for strItr := range strList {
		wordCount[strList[strItr]] += 1
	}
	strList = strings.Split(ipStrs2, " ")
	for strItr := range strList {
		wordCount[strList[strItr]] += 1
	}

	for listItr := range wordCount {
		if wordCount[listItr] == 1 {
			resMap = append(resMap, listItr)
		}
	}

	return resMap
}

func testUncommonWords() {
	testInputs1 := []string{"the quick", "the tortoise beat the haire", "copper coffee pot"}
	testInputs2 := []string{"brown fox", "the tortoise lost to the haire", "hot coffee pot"}
	expectedResults := [][]string{{"the", "quick", "brown", "fox"}, {"beat", "to", "lost"}, {"copper", "hot"}}
	testResults := [][]string{}
	for testCaseIdx := range testInputs1 {
		// fmt.Println("IP1: ", testInputs1[testCaseIdx])
		// fmt.Println("IP2: ", testInputs2[testCaseIdx])
		// result := UncommonWords(testInputs1[testCaseIdx], testInputs2[testCaseIdx])
		// fmt.Println("Result: ", result)
		testResults = append(testResults, UncommonWords(testInputs1[testCaseIdx], testInputs2[testCaseIdx]))
	}
	for i := range testResults {
		if len(testResults[i]) == len(expectedResults[i]) {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input 1:%s, Input 2:%s, Output:%s, Expected:%s\n", testInputs1[i], testInputs2[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
