package remove_nth_last_node

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

func (l *List) remove_node(n int) {
	length := 0
	if l.head != nil {
		length = 1
	} else {
		length = 0
		return
	}

	curr := l.head
	for curr.next != nil {
		length++
		curr = curr.next
	}

	curr = l.head
	currLoc := length

	for {
		if n == length {
			l.head = l.head.next
			break
		}

		currLoc--
		if n == currLoc {
			if curr.next == nil {
				break
			}
			curr.next = curr.next.next
			break
		}
		curr = curr.next
	}
}

// func main() {
// 	RunTestCases()
// }

func RunTestCases() {
	testCaseLL := []int{1, 2, 3}
	for i := range testCaseLL {
		LL1 := &List{}
		LL1.add(1)
		LL1.add(2)
		LL1.add(3)
		LL1.remove_node(i + 1)
		fmt.Printf("Test Case %d: %d; n: %d; Result: %s\n", i+1, testCaseLL[:], i+1, LL1.print())
	}
}
