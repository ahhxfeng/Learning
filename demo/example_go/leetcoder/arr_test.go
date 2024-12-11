/*
* 数组 测试

* author: ahhxfeng
* 2024@copyright
 */

package leetcoder

import "testing"

// 测试 暴力循环 求最大面积
func TestMaxAreaV(t *testing.T) {
	t.Log(" maxArea Test start \n")
	heightData := []int{8, 7, 2, 1}
	res := maxAreaV(heightData)
	t.Logf("暴力循环求解:%v", res)
}

// 测试 双指针 贪心
func TestMaxAreaG(t *testing.T) {
	t.Log(" maxArea Test start \n")
	heightData := []int{8, 7, 2, 1}
	res := maxAreaG(heightData)
	t.Logf("双指针 贪心:%v", res)
}
