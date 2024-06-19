package sort

import "testing"

func TestBubble(t *testing.T) {
	t.Log("Test bubble sort start !!")
	list1 := IntSlice{1, 3, 5, 8, 4, 6, 9, 7, 10}
	list2 := IntSlice{5, 4, 4, 6, 3, 7, 9, 11, 43, 21, 66, 100}
	list3 := IntSlice{0, 1, 5, 8, 11, 9, 9, 9, 8, 9, 11, 77}
	list1.bubbleSort()
	list2.bubbleSort()
	list3.bubbleSort()
	t.Logf("sorted slice list1: %d", list1)
	t.Logf("sorted slice list2: %d", list2)
	t.Logf("sorted slice list3 %d", list3)

}
