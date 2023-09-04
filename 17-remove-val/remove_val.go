package remove_val

import "fmt"

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

func (l *List) remove(val int) {
	curr := l.head
	if l.head.data == val {
		l.head = curr.next
		return
	}

	for curr.next != nil {
		if curr.next.data == val {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func RunTestCase() {
	testCase1 := []int{1, 2, 3}
	LL1 := &List{}
	for tCaseItr := range testCase1 {
		LL1.add(testCase1[tCaseItr])
	}
	LL1.remove(3)
	fmt.Printf("Test Case 1: %t\n", LL1.print() == "1 2")
	testCase2 := []int{8, 1, 1, 4, 12}
	LL2 := &List{}
	for tCaseItr := range testCase2 {
		LL2.add(testCase2[tCaseItr])
	}
	LL2.remove(1)
	fmt.Printf("Test Case 2: %t\n", LL2.print() == "8 4 12")
	testCase3 := []int{7, 12, 2, 9}
	LL3 := &List{}
	for tCaseItr := range testCase3 {
		LL3.add(testCase3[tCaseItr])
	}
	LL3.remove(7)
	fmt.Printf("Test Case 3: %t\n", LL3.print() == "12 2 9")
	LL3.remove(9)
	fmt.Printf("Test Case 4: %t\n", LL3.print() == "12 2")
}
