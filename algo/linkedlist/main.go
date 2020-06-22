package main

import (
	"fmt"
)

type Node struct {
	Value string
	Next  *Node
}

func insert(node *Node, val string) *Node {

	if node == nil {
		return &Node{
			Value: val,
			Next:  nil,
		}
	}

	if node.Next == nil {
		node.Next = &Node{
			Value: val,
			Next:  nil,
		}
		return node
	}

	current := node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{
		Value: val,
		Next:  nil,
	}
	return current
}

func lookup(node *Node, v string) *Node {
	if node.Value == v {
		return node
	}
	current := node
	for {
		if current.Value == v {
			return current
		}
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}
}

func remove(list *Node, val string) *Node {
	if list == nil {
		return nil
	}
	if list.Value == val {
		return list.Next
	}
	current := list
	prev := list
	for {
		if current.Value == val {
			prev.Next = current.Next
			return list
		}

		if current.Next == nil {
			return list
		}
		prev = current
		current = current.Next
	}
}

func main() {

	var linkedList *Node = nil
	linkedList = insert(linkedList, "start")
	insert(linkedList, "hello")
	insert(linkedList, "bye")
	insert(linkedList, "cat")
	n := lookup(linkedList, "bye")
	fmt.Println(n)

	n2 := remove(linkedList, "bye")

	fmt.Println(n2)
}
