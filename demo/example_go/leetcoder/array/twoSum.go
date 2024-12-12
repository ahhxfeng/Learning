/*
* 001. 两数之和 two sum
* easy
 */
package array

func twoSum(nums []int, target int) []int {
	return twoSumV(nums, target)
}

// 暴力循环 求解
// T O(n^2)
// S O(1)
func twoSumV(nums []int, target int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}

	return res
}

// hash查找
// T O(n)
// S O(n)
// 0ms 100%

func twoSumH(nums []int, target int) []int {
	res := make([]int, 0)
	numsMap := make(map[int]int)

	// 生成map
	for i, v := range nums {
		numsMap[target-v] = i
	}

	for i, v := range nums {
		if _, exsits := numsMap[v]; exsits {
			res = append(res, i, numsMap[target-v])
		}
	}

	return res
}

// hash 查找 from ChatGPT
// T 0(n)
// S O(n)
// 3ms

func twoSumHG(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return nil
}
