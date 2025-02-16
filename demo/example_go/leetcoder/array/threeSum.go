/*
* 三数之和
* mid
 */

package array

import (
	"fmt"
	"sort"
)

// 外层for循环,内部调用twoSum(nums.pop(num[i]), 0-nums[i])
// X 这样索引比较难处理
func threeSum(nums []int) []int {
	res := []int{}

	// 外层循环,固定一个数,消减两数成两数之和
	for i, v := range nums {
		fmt.Println(i)
		res = twoSum(nums, 0-v)
	}

	return res
}

// 先排序, 然后双指针逼近

func threeSumD(nums []int) [][]int {
	res := [][]int{}

	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 双指针遍历
	// for k := 0; k < len(nums)-1; k++ {
	// 	i := k + 1
	// 	j := len(nums) - 1
	// 	for i != j {
	// 		tmp := nums[k] + nums[i] + nums[j]
	// 		if tmp == 0 {
	// 			res = append(res, []int{nums[k], nums[i], nums[j]})
	// 			break
	// 		}

	// 		if tmp < 0 {
	// 			// 移向更大的数
	// 			i++
	// 		} else {
	// 			// 移向更小的数
	// 			j--
	// 		}
	// 	}
	// }

	// 双指针遍历

	for i := 0; i < len(nums)-2; i++ {
		// 跳过相同的数
		if i > 0 && nums[i] == nums[i-1] {
			// i++
			continue
		}

		left := i + 1
		right := len(nums) - 1
		//双指针 逼近
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]

			switch {
			case tmp < 0:
				left++
			case tmp > 0:
				right--
			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 跳过相同的数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}

		// 跳过,如果nums[i] >0 ,后续就不可能在=0了
		if nums[i] > 0 {
			break
		}
	}

	return res
}
