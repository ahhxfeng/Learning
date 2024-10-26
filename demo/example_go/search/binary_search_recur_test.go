package search

import "testing"

func TestBinarySerach(T *testing.T) {
	T.Log("二分搜索，测试")
	nums := []int{1, 3, 4, 5, 2, 6, 11, 9, 9, 13, 0, 40, 93}
	res := binarySearch(nums, 3)
	T.Logf("搜索结果%v", res)

}
