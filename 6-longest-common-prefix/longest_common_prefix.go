package longest_common_prefix

import (
	"fmt"
	"sort"
)

// func main() {
// 	longestCommonPrefixTest()
// }

func LongestCommonPrefix(stringList []string) string {
	if len(stringList) > 1 {
		// sort string
		sort.Strings(stringList)

		// iterate from the largest possible substr and reduce the size iteratively
		for substrLenIter := len(stringList[0]); substrLenIter >= 0; substrLenIter-- {
			checker := true
			for valIter := 1; valIter < len(stringList); valIter++ {
				checker = checker && (stringList[0][:substrLenIter] == stringList[valIter][:substrLenIter])
			}
			if checker {
				return stringList[0][:substrLenIter]
			} else {
				continue
			}
		}

		return ""
	} else {
		return ""
	}
}

func longestCommonPrefixTest() {
	loloString := [][]string{{"colorado", "color", "cold"}, {"a", "b", "c"}, {"spot", "spotty", "spotted"}}
	expectedResults := []string{"col", "", "spot"}
	testResults := []string{}
	for testCaseIdx := range loloString {
		testResults = append(testResults, LongestCommonPrefix(loloString[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input:%s Output:\"%s\"\n", loloString[i], testResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
