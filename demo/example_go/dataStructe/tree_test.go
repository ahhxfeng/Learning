package main

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Log("Test Tree Start ")

	bst := BST{}
	// bst.Insert("8")
	// bst.Insert("3")
	// bst.Insert("6")

	bst.Insert("10")
	bst.Insert("5")
	bst.Insert("15")
	bst.Insert("20")
	bst.Insert("17")
	bst.Insert("4")
	bst.Insert("6")

	bst.InOrder(bst.root)
	fmt.Println()
	bst.LevelOrder()
	fmt.Println()
	fmt.Println(bst.Search("5"))

}
