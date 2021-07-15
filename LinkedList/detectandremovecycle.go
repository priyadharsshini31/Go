package main

import (
	"fmt"
)

type Node struct {
	value int 
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func(l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.head == nil {
		l.head = &n
		l.length++
		return
	}
	ptr := l.head
	for ptr.next != nil {
		ptr = ptr.next 
	}
	ptr.next = &n
	l.length++
}

func(l *LinkedList) Print(){
	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
}


func(l *LinkedList) FormCycle(){
	ptr := l.head
	for i := 0 ; i < 3; i++ {
		ptr = ptr.next
		
	}
	tmp := ptr
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = tmp
}

func (l *LinkedList)DetectCycle() {
	slow := l.head
	fast := l.head
		
		for fast != nil && slow != nil && fast.next != nil {
		slow = slow.next 
		fast = fast.next.next
		if slow == fast {
			fmt.Println("There is cycle")
			return
		} 
	}
	
	fmt.Println("There is no cycle")
}

func (l *LinkedList) RemoveCycle() {
	cycle := make(map[*Node]int)
	temp := l.head
	for temp  != nil {
		if _, ok := cycle[temp.next]; ok{
			temp.next = nil
			break
		} else {
			cycle[temp.next] = temp.value
			temp = temp.next
	}
	}
}
	
		
	
func main() {
	l := LinkedList{}
	l.Insert(10)
	l.Insert(20)
	l.Insert(30)
	l.Insert(40)
	l.Insert(50)
	l.Insert(60)
	l.Print()
	l.FormCycle()
	l.DetectCycle()
	l.RemoveCycle()
	l.DetectCycle()
	
}
