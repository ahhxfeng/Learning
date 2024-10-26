package search

/*
二分搜索，基于分治，递归实现
*/

func dfs(nums []int, target, i, j int) int {
	// 区间为空 返回 -1
	if i > j {
		return -1
	}

	m := i + ((i + j) >> 1)
	if target < nums[m] {
		return dfs(nums, target, i, m-1)
	} else if target > nums[m] {
		return dfs(nums, target, m+1, j)
	} else {
		return m
	}

}

func binarySearch(nums []int, target int) int {
	length := len(nums)
	return dfs(nums, target, 0, length-1)
}
