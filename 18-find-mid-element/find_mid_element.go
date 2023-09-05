package find_mid_element

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

func (l *List) findMidElement() int {
	length := 0
	curr := l.head
	if curr == nil {
		return 0
	} else {
		length = 1
	}

	for {
		if curr.next != nil {
			curr = curr.next
			length++
		} else {
			break
		}
	}

	curr = l.head

	mid := length/2 + 1

	for itr := 1; itr <= mid; itr++ {
		if itr == mid {
			return curr.data
		}
		curr = curr.next
	}

	return -1
}

func testCaseExecution(lListVals []int) int {
	LL := &List{}
	for elItr := range lListVals {
		LL.add(lListVals[elItr])
	}
	return LL.findMidElement()
}

func RunTestCases() {
	testCase1 := []int{1, 2, 3}
	fmt.Printf("%v\n", testCaseExecution(testCase1) == 2)
	testCase2 := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", testCaseExecution(testCase2) == 3)
	testCase3 := []int{1}
	fmt.Printf("%v\n", testCaseExecution(testCase3) == 1)
}
