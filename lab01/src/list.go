package src

import (
	"bytes"
)

type DoublyLinkedList struct {
	head *Node
}

func (l *DoublyLinkedList) String() string {
	buffer := &bytes.Buffer{}

	curr := l.head
	idx := 0
	for curr != nil {
		if idx > 0 {
			buffer.WriteString(" <--> ")
		}
		buffer.WriteString(curr.String())

		curr = curr.next
		idx++
	}

	return buffer.String()
}

func (l *DoublyLinkedList) AddNode(node *Node) {
	if l.head == nil {
		l.head = node
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = node
	node.prev = curr
}

func (l *DoublyLinkedList) Sum() int {
	curr := l.head
	res := 0
	for curr != nil {
		res += curr.value
		curr = curr.next
	}
	return res
}

func (l *DoublyLinkedList) Average() int {
	curr := l.head
	res := 0
	i := 0
	for curr != nil {
		res += curr.value
		i++
		curr = curr.next
	}
	if i == 0 {
		return 0
	}

	return res / i
}

func (l *DoublyLinkedList) sort() {
	if l.head == nil || l.head.next == nil {
		return
	}

	curr := l.head
	for curr.next != nil {
		index := curr.next
		for index != nil {
			if curr.value > index.value {
				curr.value, index.value = index.value, curr.value
			}
			index = index.next
		}
		curr = curr.next
	}
}

func (l *DoublyLinkedList) GetSmallest(n int) []int {
	l.sort()

	curr := l.head
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if curr == nil {
			return res
		}

		res[i] = curr.value
		curr = curr.next
	}

	return res
}

func (l *DoublyLinkedList) GetLargest(n int) []int {
	l.sort()

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if curr == nil {
			return res
		}

		res[i] = curr.value
		curr = curr.prev
	}

	return res
}

func (l *DoublyLinkedList) GetNthElement(n int) *Node {
	if l.head == nil {
		return nil
	}

	curr := l.head
	idx := 0
	for idx < n {
		if idx == n {
			return curr
		}
		curr = curr.next
		idx++
	}

	return curr
}
