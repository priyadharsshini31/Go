package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}


func (t *Tree) insert(n *Node, val int) {
	newNode := Node{}
	newNode.value = val

	if t.root == nil {
		t.root = &newNode
		return
	}
		
	if val <= n.value{
		if n.left == nil {
			n.left = &newNode
		} else {
			t.insert(n.left,val)
		}
	}
	
	if val > n.value{
		if n.right == nil {
			n.right = &newNode
		} else {
			t.insert(n.right,val)
		}
	}
		
}

func (t *Tree) Inorder(temp *Node) {
    if temp == nil {
        return
	}
    fmt.Println(temp.value);
    t.Inorder(temp.left);
    t.Inorder(temp.right);
}


func main() {
	t := Tree{}
	t.insert(t.root,100)
	t.insert(t.root,-20)
	t.insert(t.root,-50)
	t.insert(t.root,-15)
	t.insert(t.root,-60)
	t.insert(t.root,50)
	t.insert(t.root,60)
	t.insert(t.root,55)
	t.insert(t.root,85)
	t.insert(t.root,15)
	t.insert(t.root,5)
	t.insert(t.root,-10)
	t.Inorder(t.root)
}
