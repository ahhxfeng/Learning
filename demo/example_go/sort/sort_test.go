package sort

import "testing"

type UnsortSlice struct {
	Slices []IntSlice
}

var unsortslice UnsortSlice

func (u *UnsortSlice) buildData() {
	u.Slices = append(u.Slices, IntSlice{1, 5, 3, 8, 4, 6, 9, 7, 10})
	u.Slices = append(u.Slices, IntSlice{5, 4, 4, 6, 3, 7, 9, 11, 43, 21, 66, 100})
	u.Slices = append(u.Slices, IntSlice{0, 1, 5, 8, 11, 9, 9, 9, 8, 9, 11, 77})
	u.Slices = append(u.Slices, IntSlice{0, 6, 5, 8, 11, 9, 9, 9, 8, 9, 11, 77, 122, 1024, 333})
}

func TestBubble(t *testing.T) {
	t.Log("Test bubble sort start !!")
	// list1 := IntSlice{1, 3, 5, 8, 4, 6, 9, 7, 10}
	// list2 := IntSlice{5, 4, 4, 6, 3, 7, 9, 11, 43, 21, 66, 100}
	// list3 := IntSlice{0, 1, 5, 8, 11, 9, 9, 9, 8, 9, 11, 77}
	// list4 := IntSlice{0, 1, 5, 8, 11, 9, 9, 9, 8, 9, 11, 77}
	// list1.bubbleSort()
	// list2.bubbleSort()
	// list3.bubbleSort()
	// list4.bubbleSort()
	// t.Logf("sorted slice list1: %d", list1)
	// t.Logf("sorted slice list2: %d", list2)
	// t.Logf("sorted slice list3 %d", list3)
	// t.Logf("sorted slice list4 %d", list4)

	unsortslice.buildData()
	for i := 0; i < len(unsortslice.Slices); i++ {
		t.Logf("Origin Slice %d ", unsortslice.Slices[i])
		unsortslice.Slices[i].bubbleSort()
		t.Logf("sorted Slice %d ", unsortslice.Slices[i])
		t.Log("\n")

	}

}

func TestSelectSort(t *testing.T) {
	t.Log("Test simple select sort start !!")
	unsortslice.buildData()
	for i := 0; i < len(unsortslice.Slices); i++ {
		t.Logf("Origin Slice %d ", unsortslice.Slices[i])
		// newSlice := unsortslice.Slices[i].selectSort()
		// t.Logf("sorted Slice %d ", newSlice)
		unsortslice.Slices[i].selectSort()
		t.Logf("sorted Slice %d ", unsortslice.Slices[i])
		t.Log("\n")

	}

}

func TestStraightInsert(t *testing.T) {
	t.Log("Test straight insert sort start !!")
	unsortslice.buildData()
	for i := 0; i < len(unsortslice.Slices); i++ {
		t.Logf("Origin Slice %d ", unsortslice.Slices[i])
		// newSlice := unsortslice.Slices[i].selectSort()
		// t.Logf("sorted Slice %d ", newSlice)
		unsortslice.Slices[i].straightInsert()
		t.Logf("sorted Slice %d ", unsortslice.Slices[i])
		t.Log("\n")

	}

}

func TestMergeSort(t *testing.T) {
	t.Log("归并 排序 测试")
	unsortslice.buildData()
	for i := 0; i < len(unsortslice.Slices); i++ {
		t.Logf("Origin Slice %v", unsortslice.Slices[i])
		mergeSort(unsortslice.Slices[i], 0, len(unsortslice.Slices[i]))
		t.Logf("Sorted Slice %v\n", unsortslice.Slices[i])
	}

}
