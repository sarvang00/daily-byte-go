package add_binary

import "fmt"

// func main() {
// 	addBinaryTest()
// }

func AddBinary(principal string, adder string) string {
	if adder == "1" {
		if len(principal) > 1 {
			if principal[len(principal)-1:] == "0" {
				// If multi-valued string with last char as 0
				return principal[:len(principal)-1] + "1"
			} else {
				// If multi-valued string with last char as 1
				return AddBinary(principal[:len(principal)-1], "1") + "0"
			}
		} else {
			// adder is 1 but principal is single digit
			if principal == "1" {
				return "10"
			} else {
				return "1"
			}
		}
	} else {
		return principal
	}
}

func addBinaryTest() {
	testIps := [][2]string{{"100", "1"}, {"11", "1"}, {"1", "0"}}
	expectedResults := []string{"101", "100", "1"}
	testResults := []string{}
	for testCaseIdx := range testIps {
		testResults = append(testResults, AddBinary(testIps[testCaseIdx][0], testIps[testCaseIdx][1]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input:%s Output:%s\n", testIps[i], testResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
