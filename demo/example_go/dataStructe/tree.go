package main

import "fmt"

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

type BST struct {
	root *TreeNode
}

func (bst *BST) Insert(data string) {
	bst.InsertRec(bst.root, data)
}

func (bst *BST) InsertRec(node *TreeNode, data string) *TreeNode {
	if bst.root == nil {
		bst.root = &TreeNode{data, nil, nil}
		return bst.root
	}

	if node == nil {
		node = &TreeNode{data, nil, nil}
	}

	if data < node.data {
		node.left = bst.InsertRec(node.left, data)
	}

	if data > node.data {
		node.right = bst.InsertRec(node.right, data)
	}

	return node
}

func (bst *BST) Search(data string) bool {
	found := bst.SearchRec(bst.root, data)
	return found
}

func (bst *BST) SearchRec(node *TreeNode, data string) bool {
	if node.data == data {
		return true
	}

	if node == nil {
		return false
	}

	if data < node.data {
		return bst.SearchRec(node.left, data)
	}

	if data > node.data {
		return bst.SearchRec(node.right, data)
	}

	return false
}

// 中序遍历
func (bst *BST) InOrder(node *TreeNode) {
	if node == nil {
		return
	} else {
		bst.InOrder(node.left)
		fmt.Printf("  %s", node.data)
		bst.InOrder(node.right)
	}

}

func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}

	nodeList := make([](*TreeNode), 0)
	nodeList = append(nodeList, bst.root)

	for !(len(nodeList) == 0) {
		current := nodeList[0]
		fmt.Print(current.data, " ")
		if current.left != nil {
			nodeList = append(nodeList, current.left)
		}
		if current.right != nil {
			nodeList = append(nodeList, current.right)
		}
		nodeList = nodeList[1:]
	}
}
