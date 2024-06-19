package sort

type IntSlice []int

// 从最后最大的开始，逐渐有序
func (l IntSlice) bubbleSort() IntSlice {
	for i := 0; i < len(l); i++ {
		// j < len(l)-1 O(n^2)
		// better j < len(1)-1-i O(nlogn)
		for j := 0; j < len(l)-1-i; j++ {
			// l[i] > l[j] is not right ,because is not bubble thought
			//
			// think bubble pop up this bigest one to the end
			// in the first for range
			if l[j] > l[j+1] {
				l.Swap(j, j+1)
			}
		}
	}
	return l
}

func (l IntSlice) Swap(i int, j int) {
	// swap the index i and j value
	tmp := l[j]
	l[j] = l[i]
	l[i] = tmp
}
