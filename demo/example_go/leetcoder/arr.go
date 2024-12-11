/*
* 数组
* diffcultiy: easy, mid ,hard
* T: time 时间复杂度
* S: storage 空间复杂度
* author: ahhxfeng
* 2024@copyright
 */

package leetcoder

/*
* .011 盛水最多的容器
 */
func maxArea(height []int) int {
	return maxAreaV(height)
}

// 暴力解法，双层循环，找最大值
// T: O(n^2)
// S: O(1)
func maxAreaV(height []int) int {
	// left, right
	// area min(height[left], height[right]) * (right -left)
	max := 0
	for l := 0; l < len(height); l++ {
		for r := l + 1; r < len(height); r++ {
			tmp := min(height[l], height[r]) * (r - l)
			if tmp > max {
				max = tmp
			}
		}
	}

	return max

}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

// 双指针 贪心
// T: O(n)
// S: O(1)
// Greedy
func maxAreaG(height []int) int {
	// left, right
	// 最大面积短板(height[left], height[right] 小的值) * 索引差值
	// area =  min(height[left], height[right]) * (right -left)
	// 向内移动长板时 height 不变, 但width = (right - left) -1 所以面积不会比当前大
	// 向内移动长板时 height 可能变大, width = (right - left) -1 所以面积可能变大
	l := 0
	r := len(height) - 1
	maxArea := 0

	// 直到指针相遇
	for l != r {
		area := min(height[l], height[r]) * (r - l)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
