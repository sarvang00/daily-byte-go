package contains_cycle

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}

		curr.next = newNode
	}
}

func (l *List) print() string {
	listVal := ""
	curr := l.head
	for {
		listVal += fmt.Sprintf("%d ", curr.data)
		if curr.next != nil {
			curr = curr.next
		} else {
			break
		}
	}
	return listVal[:len(listVal)-1]
}

func (l *List) detectCycle() bool {
	detectionMap := make(map[int]bool)
	curr := l.head
	if curr != nil {
		for {
			if detectionMap[curr.data] {
				return true
			} else {
				detectionMap[curr.data] = true
			}

			if curr.next != nil {
				curr = curr.next
			} else {
				break
			}
		}
	}
	return false
}

func testCaseExecution(lListVals []int) bool {
	LL := &List{}
	for elItr := range lListVals {
		LL.add(lListVals[elItr])
	}
	return LL.detectCycle()
}

func RunTestCases() {
	testCase1 := []int{1, 2, 3, 1}
	fmt.Printf("%v\n", testCaseExecution(testCase1) == true)
	testCase2 := []int{1, 2, 3}
	fmt.Printf("%v\n", testCaseExecution(testCase2) == false)
	testCase3 := []int{1, 1}
	fmt.Printf("%v\n", testCaseExecution(testCase3) == true)
}
