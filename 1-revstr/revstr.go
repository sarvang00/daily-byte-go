package revstr

import (
	"fmt"
	"strings"
)

// func main() {
// 	testReverseString()
// }

func ReverseString(ipString string) string {
	alph := strings.Split(ipString, "")
	str2 := ""
	if len(ipString) > 0 {
		for a := len(alph) - 1; a >= 0; a-- {
			str2 += alph[a]
		}
		return str2
	} else {
		return "eyy... Vediya"
	}
}

func testReverseString() {
	testStr := []string{"Cat", "The Daily Byte", "civic", ""}
	for i := range testStr {
		fmt.Printf("Input: \"%s\" Output: \"%s\"\n", testStr[i], ReverseString(testStr[i]))
	}
}
