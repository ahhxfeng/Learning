package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkList struct {
	head *Node
}

func (list *LinkList) append(data string) {
	// insert data at linklist end

	newNode := &Node{
		data: data,
		next: nil,
	}

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

func (list *LinkList) prepend(data string) {
	// insert data before first data

	newNode := &Node{
		data: data,
		next: list.head,
	}

	list.head = newNode

}

func (list *LinkList) deleteWithValue(data string) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		return
	}

	current := list.head
	// current.next != nil error: invaild memory address
	for current != nil {
		if current.next.data == data {
			current.next = current.next.next
		}
		current = current.next
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
