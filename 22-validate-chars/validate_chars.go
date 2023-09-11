package validateChars

import (
	"fmt"
	"strings"
)

type Element struct {
	data string
	prev *Element
}

type Stack struct {
	top *Element
}

func isOpen(ipStr string) bool {
	return ipStr == "(" || ipStr == "{" || ipStr == "["
}

func getRev(ipStr string) string {
	switch ipStr {
	case "}":
		return "{"
	case ")":
		return "("
	case "]":
		return "["
	default:
		// return "no rev of " + ipStr
		return ""
	}
}

func (s *Stack) push(ipStr string) {
	newElement := &Element{data: ipStr, prev: s.top}
	if s != nil && !isOpen(ipStr) {
		if s.top.data == getRev(ipStr) {
			// fmt.Println("Popping valid")
			s.pop()
		} else {
			// fmt.Println("Entered: ", ipStr)
			// fmt.Println("Not reverse of: ", getRev(s.top.data))
			// fmt.Println("Adding invalid")
			s.top = newElement
		}
	} else {
		// fmt.Println("Pushing")
		s.top = newElement
	}
}

func (s *Stack) isComplete() bool {
	return s.printTop() == ""
}

func (s *Stack) pop() string {
	if s.top != nil {
		el := s.top.data
		s.top = s.top.prev
		return el
	}
	return ""
}

func (s *Stack) printTop() string {
	if s.top != nil {
		return s.top.data
	}
	return ""
}

func (s *Stack) printStack() string {
	stackTrace := ""
	for s.top != nil {
		stackTrace += s.top.data
		if s.top.prev != nil {
			s.top = s.top.prev
		} else {
			break
		}
	}
	retStr := ""
	for i := len(stackTrace); i > 0; i-- {
		retStr += stackTrace[i-1 : i]
	}
	return retStr
}

func runTestCase(ipStr string) bool {
	newStack := &Stack{}
	strs := strings.Split(ipStr, "")
	for i := range strs {
		newStack.push(strs[i])
	}
	return newStack.isComplete()
}

func RunTestCases() {
	testCase1 := "(){}[]"
	fmt.Printf("%v Result: %v\n", testCase1, runTestCase(testCase1) == true)
	testCase2 := "(({[]}))"
	fmt.Printf("%v Result: %v\n", testCase2, runTestCase(testCase2) == true)
	testCase3 := "{(})"
	fmt.Printf("%v Result: %v\n", testCase3, runTestCase(testCase3) == false)
}
