package removeadjacentduplicates

import "fmt"

type strMan struct {
	strokes string
}

func (k *strMan) enterKeyStroke(ipStr string) {
	cLen := len(k.strokes)
	if cLen > 0 && ipStr == k.strokes[cLen-1:cLen] {
		k.strokes = k.strokes[:cLen-1]
		return
	}
	k.strokes += ipStr
}

func (k *strMan) printString() string {
	return k.strokes
}

func runTestCase(s string) string {
	strStream := &strMan{}
	for i := range s {
		strStream.enterKeyStroke(s[i : i+1])
	}
	return strStream.printString()
}

func RunTestCases() {
	testCase1 := "abccba"
	fmt.Printf("s:%v pass:%v\n", testCase1, runTestCase(testCase1) == "")
	testCase2 := "foobar"
	fmt.Printf("s:%v pass:%v\n", testCase2, runTestCase(testCase2) == "fbar")
	testCase3 := "abccbefggfe"
	fmt.Printf("s:%v pass:%v\n", testCase3, runTestCase(testCase3) == "a")
}
