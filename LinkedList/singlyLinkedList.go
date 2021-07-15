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

func (l *LinkedList) search(value int) {
	ptr := l.head
	for ptr != nil && ptr.next != nil {
		if ptr.value == value {
			fmt.Println("The value is present")
			return
		}
		ptr = ptr.next
	}
	fmt.Println("The value is not present")
}

func (l *LinkedList) delete(value int) {
	ptr := l.head
	var prev *Node
	for ptr.next != nil && ptr.value != value {
		prev = ptr
		ptr = ptr.next
	}
	if ptr != l.head {
		prev.next = ptr.next
		l.length--
		return
	} else {
		l.head = ptr.next
		l.length--
		return
	}
}

func (l *LinkedList) findMiddle() {
	slow := l.head
	fast := l.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	fmt.Println(slow.value)
}

func (l *LinkedList) reverse() {
	current := l.head
	var prev *Node
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
		l.head = prev
	}
}

// Maintain two pointers – reference pointer and main pointer. 
// Initialize both reference and main pointers to head. 
// First, move the reference pointer to n nodes from head. 
// Now move both pointers one by one until the reference pointer reaches the end. 
// Now the main pointer will point to nth node from the end. 
// Return the main pointer.

func (l *LinkedList) findNthElementFromEnd(position int) {
	temp := l.head
	main := l.head
	for i := 0; i < position; i++ {
		temp = temp.next
	}
	for temp != nil {
		main = main.next
		temp = temp.next

	}
	fmt.Println("the node at the nth position is", main.value)
}

func (l *LinkedList) rotate(rotation int) {
	temp := l.head
	ptr := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = l.head
	for i := 1; i < rotation; i++ {
		ptr = ptr.next
	}
	l.head = ptr.next
	ptr.next = nil
}

func (l *LinkedList) formCycle() {
	temp := l.head
	for i := 0; i < 2; i++ {
		temp = temp.next
	}
	cyclestart := temp

	temp = l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = cyclestart
}

func (l *LinkedList) detectCycle() {
	slow := l.head
	fast := l.head
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			fmt.Println("There is a cycle")
			return
		}
	}
	fmt.Println("There is no cycle")
}

func (l *LinkedList) removeCycleWithMap() {
	cycle := make(map[*Node]int)
	temp := l.head

	

		for temp != nil {
			if _, ok := cycle[temp.next]; ok {
				temp.next = nil
				return
			} else {
				cycle[temp.next] = temp.value
				temp = temp.next

			}
		}
	
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
	singlyLinkedList.print()
	fmt.Println("Printing the middle element")
	singlyLinkedList.findMiddle()
	fmt.Println("Find the nth element from the end")
	singlyLinkedList.findNthElementFromEnd(2)
	fmt.Println("Reversing the linked list")
	singlyLinkedList.reverse()
	singlyLinkedList.print()
	fmt.Println("Rotating the linked list")
	singlyLinkedList.rotate(3)
	singlyLinkedList.print()
	singlyLinkedList.search(10)
	singlyLinkedList.search(100)	
	singlyLinkedList.delete(10)
	singlyLinkedList.print()
	fmt.Println("Now deleting")
	singlyLinkedList.delete(40)
	singlyLinkedList.print()
	singlyLinkedList.detectCycle()
	singlyLinkedList.formCycle()
	singlyLinkedList.detectCycle()
	singlyLinkedList.removeCycleWithMap()
	singlyLinkedList.print()
	singlyLinkedList.detectCycle()

}
