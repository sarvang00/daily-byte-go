package spot_the_diff

import (
	"fmt"
)

// func main() {
// 	testSpotDiff()
// }

func SpotDiff(ipString1 string, ipString2 string) string {
	if len(ipString1) != len(ipString2) {
		letterMap := make(map[string]int)
		for itr := range ipString1 {
			letterMap[ipString1[itr:itr+1]] += 1
		}
		for itr := range ipString2 {
			if letterMap[ipString2[itr:itr+1]] == 0 {
				return ipString2[itr : itr+1]
			}
		}
	}
	return ""
}

func testSpotDiff() {
	testInputs1 := []string{"foobar", "ide", "coding"}
	testInputs2 := []string{"barfoot", "idea", "ingcod"}
	expectedResults := []string{"t", "a", ""}
	testResults := []string{}
	for testCaseIdx := range expectedResults {
		testResults = append(testResults, SpotDiff(testInputs1[testCaseIdx], testInputs2[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Input 1:%s, Input 1:%s, Output:%s, Expected:%s\n", testInputs1[i], testInputs2[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
