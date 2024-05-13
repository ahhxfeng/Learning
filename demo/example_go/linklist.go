package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkList struct {
	head *Node
}

func (list *LinkList) insert(data string) {
	newNode := &Node{data: data}

	if list.head == nil {
		// LinkList is empty
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkList) display() {
	current := list.head

	if current.next == nil {
		fmt.Println("LinkList is empty")
		return
	}

	fmt.Print("LinkList: ")
	for current != nil {
		fmt.Printf("%s ", current.data)
		current = current.next
	}
	fmt.Println("")
}

func main() {
	list := LinkList{}

	list.insert("1")
	list.insert("20")
	list.insert("33")
	list.insert("99")
	list.insert("77")

	list.display()
}
