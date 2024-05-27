package main

import "testing"

func TestStack(t *testing.T) {
	Sstack := Stack{}

	Sstack.Pop()
	Sstack.Push("This")
	Sstack.Push("is")
	Sstack.Push("just")
	Sstack.Push("a")
	Sstack.Push("Test")
	Sstack.Push("and")
	Sstack.Push("today")
	Sstack.Push("is")
	Sstack.Push("very")
	Sstack.Push("nice")
	Sstack.Push("day")
	t.Log("Test start !!!")
	t.Log(Sstack.Pop())
	t.Log(Sstack.Pop())
	t.Log(Sstack.Pop())
	t.Log(Sstack.Pop())
	t.Log(Sstack.Pop())
	// t.Log(Sstack.Pop())
}
