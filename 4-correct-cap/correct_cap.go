package correct_cap

import (
	"fmt"
	"strings"
)

// func main() {
// 	testCorrectCap()
// }

func CorrectCap(ipString string) bool {
	// Fully Upper Case String
	if ipString == strings.ToUpper(ipString) {
		return true
	}
	// Fully Lower Case String
	if ipString == strings.ToLower(ipString) {
		return true
	}
	// If First is upper and rest is lower
	if ipString[0:1] == strings.ToUpper(ipString[0:1]) && ipString[1:] == strings.ToLower(ipString[1:]) {
		return true
	}
	return false
}

func testCorrectCap() {
	testStrgs := []string{"USA", "Calvin", "compUter", "coding"}
	expectedResults := []bool{true, true, false, true}
	testResults := []bool{}
	for testStr := range testStrgs {
		testResults = append(testResults, CorrectCap(testStrgs[testStr]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
