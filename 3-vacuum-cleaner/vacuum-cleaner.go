package vaccum_cleaner

import (
	"fmt"
	"strings"
)

// func main() {
// 	testVacuumCleaner()
// }

func VaccumCleaner(ipPath string) bool {
	charSet := strings.Split(ipPath, "")
	x, y := 0, 0
	for char := range charSet {
		switch charSet[char] {
		case "U":
			y += 1
		case "D":
			y -= 1
		case "L":
			x -= 1
		case "R":
			x += 1
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}

func testVacuumCleaner() {
	testStrgs := []string{"LR", "URURD", "RUULLDRD"}
	expectedResults := []bool{true, false, true}
	testResults := []bool{}
	for testStr := range testStrgs {
		testResults = append(testResults, VaccumCleaner(testStrgs[testStr]))
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
