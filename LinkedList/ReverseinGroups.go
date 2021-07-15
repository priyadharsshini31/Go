package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insert(val int) {
	n := Node{}
	n.value = val
	ptr := l.head
	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}

	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &n
	l.length++
}

func (l *LinkedList) print() {
	ptr := l.head
	for ptr != nil {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}
}

func (l *LinkedList)ReverseInGroups(head *Node, val int) (prev *Node) {
	var next *Node
	current := head
	count := 0
	for current != nil && count < val {
		next = current.next
		current.next = prev
		prev = current
		current = next
		count++
	}
	if next != nil {
		head.next = l.ReverseInGroups(next,val)
	}
	return prev
	
}

func main() {
	singlyLinkedList := LinkedList{}
	singlyLinkedList.insert(10)
	singlyLinkedList.insert(20)
	singlyLinkedList.insert(30)
	singlyLinkedList.insert(40)
	singlyLinkedList.insert(50)
	singlyLinkedList.insert(60)
	singlyLinkedList.insert(70)
	singlyLinkedList.head = singlyLinkedList.ReverseInGroups(singlyLinkedList.head,3)
	singlyLinkedList.print()
	
	

}
