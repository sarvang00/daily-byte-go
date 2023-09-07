package return_cycle_start

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

func (l *List) detectCycleStart() *Node {
	detectionMap := make(map[int]bool)
	curr := l.head
	if curr != nil {
		for {
			if detectionMap[curr.data] {
				return curr
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
	return nil
}

func testCaseExecution(lListVals []int) *Node {
	LL := &List{}
	for elItr := range lListVals {
		LL.add(lListVals[elItr])
	}
	return LL.detectCycleStart()
}

func RunTestCase() {
	testCase1 := []int{1, 2, 3}
	fmt.Printf("%v\n", testCaseExecution(testCase1) == nil)
	testCase2 := []int{1, 2, 3, 4, 5, 2}
	fmt.Printf("%v\n", testCaseExecution(testCase2).data == 2)
	testCase3 := []int{1, 9, 3, 7, 7}
	fmt.Printf("%v\n", testCaseExecution(testCase3).data == 7)
}
