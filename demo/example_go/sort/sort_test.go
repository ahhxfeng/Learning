package sort

import "testing"

func TestBubble(t *testing.T) {
	t.Log("Test bubble sort start !!")
	list1 := IntSlice{1, 3, 5, 8, 4, 6, 9, 7, 10}
	list1.bubbleSort()
	t.Logf("sorted slice %d", list1)
}
