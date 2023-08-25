package palin_on_removal

import (
	"fmt"
	"regexp"
	"strings"
)

// func main() {
// 	testIsPalindromeOnRemoval()
// }

func cleanString(ipString string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return nonAlphanumericRegex.ReplaceAllString(ipString, "")
}

func IsPalindromeOnRemoval(ipString string) bool {
	if isPalindrome(ipString) {
		return true
	} else {
		versionedResults := false
		for strItr := range ipString {
			versionedResults = versionedResults || isPalindrome(ipString[0:strItr]+ipString[strItr+1:])
		}
		return versionedResults
	}
}

func isPalindrome(ipString string) bool {
	if len(ipString) > 1 {
		cleanString := strings.ToLower(cleanString(ipString))
		for left, right := 0, len(cleanString)-1; left < right; left, right = left+1, right-1 {
			if cleanString[left] == cleanString[right] {
				continue
			} else {
				return false
			}
		}
		return true
	} else {
		return true
	}
}

func testIsPalindromeOnRemoval() {
	loloString := []string{"abcba", "foobof", "abccab"}
	expectedResults := []bool{true, true, false}
	testResults := []bool{}
	for testCaseIdx := range loloString {
		testResults = append(testResults, IsPalindromeOnRemoval(loloString[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input:%s Output:\"%t\"\n", loloString[i], testResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
