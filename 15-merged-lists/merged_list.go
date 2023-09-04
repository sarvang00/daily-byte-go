package merged_list

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

// func (l *List) remove(val int) {
// 	curr := l.head
// 	if l.head.data == val {
// 		l.head = curr.next
// 		return
// 	}

// 	for curr.next != nil {
// 		if curr.next.data == val {
// 			curr.next = curr.next.next
// 			return
// 		} else {
// 			curr = curr.next
// 		}
// 	}
// }

func mergeLists(l1 *List, l2 *List) *List {
	resList := &List{}
	l1.add(-1)
	l2.add(-1)
	curr1 := l1.head
	curr2 := l2.head

	for curr1.next != nil || curr2.next != nil {
		if curr1.next != nil && curr2.next != nil {
			if curr1.data < curr2.data {
				resList.add(curr1.data)
				if curr1.next != nil {
					curr1 = curr1.next
				}
			} else if curr1.data > curr2.data {
				resList.add(curr2.data)
				if curr2.next != nil {
					curr2 = curr2.next
				}
			}
		} else if curr1.next == nil {
			if curr2.next != nil {
				resList.add(curr2.data)
				curr2 = curr2.next
			} else {
				break
			}
		} else if curr2.next == nil {
			if curr1.next != nil {
				resList.add(curr1.data)
				curr1 = curr1.next
			} else {
				break
			}
		} else {
			break
		}
	}
	return resList
}

func testCaseExecution(lVals1 []int, lVals2 []int) string {
	LL1 := &List{}
	LL2 := &List{}
	for lValItr := range lVals1 {
		LL1.add(lVals1[lValItr])
	}
	// fmt.Println(LL1.print())
	for lValItr := range lVals2 {
		LL2.add(lVals2[lValItr])
	}
	// fmt.Println(LL2.print())
	LL3 := mergeLists(LL1, LL2)
	return LL3.print()
}

func RunTestCases() {
	testCase1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("Test Case 1: %d %d; result: %t\n", testCase1[0][:], testCase1[1][:], testCaseExecution(testCase1[0][:], testCase1[1][:]) == "1 2 3 4 5 6")
	testCase2 := [2][3]int{{1, 3, 5}, {2, 4, 6}}
	fmt.Printf("Test Case 2: %d %d; result: %t\n", testCase2[0][:], testCase2[1][:], testCaseExecution(testCase2[0][:], testCase2[1][:]) == "1 2 3 4 5 6")
	testCase3 := [2][3]int{{4, 4, 7}, {1, 5, 6}}
	fmt.Printf("Test Case 1: %d %d; result: %t\n", testCase3[0][:], testCase3[1][:], testCaseExecution(testCase3[0][:], testCase3[1][:]) == "1 4 4 5 6 7")
}
