package palin

import (
	"fmt"
	"regexp"
	"strings"
)

// func main() {
// 	testIsPalindrome()
// }

func cleanString(ipString string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return nonAlphanumericRegex.ReplaceAllString(ipString, "")
}

func IsPalindrome(ipString string) bool {
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

func testIsPalindrome() {
	testStr := []string{"level", "algorithm", "A man, a plan, a canal: Panama.", ""}
	for i := range testStr {
		fmt.Printf("Input: \"%s\" Output: %t\n", testStr[i], IsPalindrome(testStr[i]))
	}
}
