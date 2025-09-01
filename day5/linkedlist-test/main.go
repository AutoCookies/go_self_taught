package main

import "fmt"

type Node struct {
	element interface{}
	next    *Node
}

type LinkedList struct {
	head *Node
}

// Find searches for a node with the given element
func (l *LinkedList) Find(element interface{}) *Node {
	current := l.head
	for current != nil {
		if current.element == element {
			return current
		}
		current = current.next
	}
	return nil // not found
}

// Insert inserts new node at head
func (l *LinkedList) Insert(element interface{}) {
	newNode := &Node{element: element, next: l.head}
	l.head = newNode
}

// Display all elements
func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Print(current.element, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	ll.Display()

	node := ll.Find(2)
	if node != nil {
		fmt.Println("Found:", node.element)
	} else {
		fmt.Println("Not found")
	}
}
