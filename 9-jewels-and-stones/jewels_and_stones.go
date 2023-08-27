package jewels_and_stones

import "fmt"

// func main() {
// 	testJewelsAndStones()
// }

func JewelsAndStones(testJewels string, testStones string) int {
	jewelCount := 0
	jewels := make(map[string]int)
	for iter := range testJewels {
		jewels[testJewels[iter:iter+1]] = 0
	}
	for iter := range testStones {
		_, check := jewels[testStones[iter:iter+1]]
		if check {
			jewels[testStones[iter:iter+1]] += 1
		}
	}
	for _, val := range jewels {
		jewelCount += val
	}
	// fmt.Println(jewels)
	return jewelCount
}

func testJewelsAndStones() {
	testJewels := []string{"abc", "Af", "AYOPD"}
	testStones := []string{"ac", "AaaddfFf", "ayopd"}
	expectedResults := []int{2, 3, 0}
	testResults := []int{}
	for testCaseIdx := range testJewels {
		testResults = append(testResults, JewelsAndStones(testJewels[testCaseIdx], testStones[testCaseIdx]))
	}
	for i := range testResults {
		if testResults[i] == expectedResults[i] {
			fmt.Printf("Passed %d!!\n", i+1)
			continue
		} else {
			fmt.Printf("Jewels:%s, Stones:%s, Output:%d, Expected:%d\n", testJewels[i], testStones[i], testResults[i], expectedResults[i])
			fmt.Printf("Failing %d!!\n", i+1)
			break
		}
	}
}
