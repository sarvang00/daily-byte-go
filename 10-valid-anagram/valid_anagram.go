package valid_anagram

import (
	"fmt"
)

// func main() {
// 	testValidAnagram()
// }

func ValidAnagram(originalInput string, anagramInput string) bool {
	wordMap := make(map[string]int)
	// put all original inputs in a hash map
	for itr := range originalInput {
		wordMap[originalInput[itr:itr+1]] += 1
	}
	// check all anagrams for count in hash map
	for itr := range anagramInput {
		if wordMap[anagramInput[itr:itr+1]] <= 0 {
			return false
		} else {
			wordMap[anagramInput[itr:itr+1]] -= 1
		}
	}
	// if match==0 return false
	return true
}

func testValidAnagram() {
	testInputs := []string{"cat", "listen", "program"}
	testGrams := []string{"tac", "silent", "function"}
	expectedResults := []bool{true, true, false}
	testResults := []bool{}
	for testCaseIdx := range testInputs {
		testResults = append(testResults, ValidAnagram(testInputs[testCaseIdx], testGrams[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Inputs:%s, Grams:%s, Output:%t, Expected:%t\n", testInputs[i], testGrams[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
