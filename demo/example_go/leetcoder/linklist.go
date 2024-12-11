package leetcoder

// linked list

// Leetcoder 707
// 设计链表

type MyLinkedList struct {
	dummy *ListNode
}

func Constructor() MyLinkedList {
	rear := &ListNode{
		Val:  -1,
		Next: nil,
	}
	rear.Next = rear
	return MyLinkedList{rear}
}

func (*MyLinkedList) AddAtHead(val int) {

}
