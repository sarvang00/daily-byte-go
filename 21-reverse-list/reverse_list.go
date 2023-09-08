package reverse_list

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

func (l *List) reverseList() {
	var prev, next *Node
	curr := l.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func testCaseExecution(lListVals []int) string {
	LL := &List{}
	for elItr := range lListVals {
		LL.add(lListVals[elItr])
	}
	LL.reverseList()
	return LL.print()
}

func RunTests() {
	testCase1 := []int{1, 2, 3}
	fmt.Printf("%v\n", testCaseExecution(testCase1) == "3 2 1")
	testCase2 := []int{7, 15, 9, 2}
	fmt.Printf("%v\n", testCaseExecution(testCase2) == "2 9 15 7")
	testCase3 := []int{1}
	fmt.Printf("%v\n", testCaseExecution(testCase3) == "1")
}
