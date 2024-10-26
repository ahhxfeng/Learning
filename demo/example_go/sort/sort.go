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
		// declear the index of min value
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

// TODO：
// 直接插入排序,
// func (l IntSlice) straightInsert() IntSlice {
// 	for i := 1; i < len(l); i++ {
// 		for j := i - 1; j >= 0; j-- {
// 			temp := l[i]
// 			if temp < l[j] {
// 				// 选中的 比当前小，当前的后移一位
// 				l[j+1] = l[j]
// 				continue
// 			}
// 			l[j] = temp

// 		}
// 	}

// 	return l
// }

// 直接插入排序，

func (l IntSlice) straightInsert() IntSlice {
	for i := 1; i < len(l); i++ {
		temp := l[i]
		j := i - 1
		for j >= 0 && l[j] > temp {
			l[j+1] = l[j]
			j--
		}
		l[j+1] = temp
	}

	return l
}

// 快速排序
// func (l IntSlice) quickSort() {
// 	Qsort(l, 0, len(l))

// }

// func Qsort(nums []int) {

// }
