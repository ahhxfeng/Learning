package sort

type IntSlice []int

// 从最后最大的开始，逐渐有序
func (l IntSlice) bubbleSort() IntSlice {
	for i := 0; i < len(l); i++ {
		// j < len(l)-1 O(n^2)
		// better j < len(1)-1-i O(nlogn)
		// assert Slice is ordered ?
		var order_flag = true
		for j := 0; j < len(l)-1-i; j++ {
			// l[i] > l[j] is not right ,because is not bubble thought
			//
			// think bubble pop up this bigest one to the end
			// in the first for range
			if l[j] > l[j+1] {
				l.Swap(j, j+1)
				order_flag = false
			}
		}

		if order_flag {
			return l
		}
	}
	return l
}

func (l IntSlice) Swap(i int, j int) {
	// swap the value of index i and j
	tmp := l[j]
	l[j] = l[i]
	l[i] = tmp
}

// 简单选择，从index 0 开始，每次选最小的，放在当前位置
func (l IntSlice) selectSort() IntSlice {
	for i := 0; i < len(l); i++ {
		// declear the min value index
		min := i
		for j := i + 1; j < len(l); j++ {
			if l[min] > l[j] {
				min = j
			}

		}
		if i != min {

			l.Swap(i, min)
		}

	}
	return l
}
