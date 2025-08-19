package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d ", n.value)
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(value int) {

	node := new(Node)
	node.value = value

	if l.head == nil {
		l.head = node
	} else {
		for iterator := l.head; iterator != nil; iterator = iterator.next {
			if iterator.next == nil {
				iterator.next = node
				break
			}
		}
	}

}

func (l *LinkedList) AddFront(value int) {
	node := Node{value, l.head}
	node.next = l.head
	l.head = &node
}

func (l LinkedList) Print() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(iterator.String())
	}
	return sb.String()
}

func main() {
	list := LinkedList{}
	list.Add(5)
	list.Add(6)

	fmt.Println(list.Print())

	list1 := LinkedList{}
	list1.AddFront(5)
	list1.AddFront(6)
	list1.AddFront(7)
	fmt.Println(list1.Print())

}
