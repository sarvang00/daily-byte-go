package comparekeystrokes

import "fmt"

type keyPad struct {
	strokes string
}

func (k *keyPad) enterKeyStroke(ipStr string) {
	if ipStr == "#" {
		if k.strokes != "" {
			k.strokes = k.strokes[:len(k.strokes)-1]
		}
	} else {
		k.strokes += ipStr
	}
}

func (k *keyPad) printString() string {
	return k.strokes
}

func runTestCase(s string, t string) bool {
	sPad := keyPad{strokes: ""}
	for i := range s {
		sPad.enterKeyStroke(s[i : i+1])
	}
	tPad := keyPad{strokes: ""}
	for i := range t {
		tPad.enterKeyStroke(t[i : i+1])
	}
	return sPad.printString() == tPad.printString()
}

func RunTestCases() {
	testCase1 := []string{"ABC#", "CD##AB"}
	fmt.Printf("s:%v t:%v pass:%v\n", testCase1[0], testCase1[1], runTestCase(testCase1[0], testCase1[1]) == true)
	testCase2 := []string{"como#pur#ter", "computer"}
	fmt.Printf("s:%v t:%v pass:%v\n", testCase2[0], testCase2[1], runTestCase(testCase2[0], testCase2[1]) == true)
	testCase3 := []string{"cof#dim#ng", "code"}
	fmt.Printf("s:%v t:%v pass:%v\n", testCase3[0], testCase3[1], runTestCase(testCase3[0], testCase3[1]) == false)
}
