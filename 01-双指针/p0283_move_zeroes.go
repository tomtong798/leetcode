//go:build p0041

// LeetCode 第283题 - 移动零 (Move Zeroes)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】数组、双指针
//
// 【题目描述】
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意，必须在不复制数组的情况下原地对数组进行操作。
//
// 【示例 1】
// 输入：nums = [0,1,0,3,12]
// 输出：[1,3,12,0,0]
//
// 【示例 2】
// 输入：nums = [0]
// 输出：[0]
//
// 【数据范围】
// 1 <= nums.length <= 10^4
// -2^31 <= nums[i] <= 2^31 - 1
//
// 【思路提示】快慢指针：慢指针指向下一个非零元素应放的位置，
// 快指针遇到非零就与慢指针交换并前移慢指针，一遍扫完。
//
// 题目链接：https://leetcode.cn/problems/move-zeroes/
package main

// moveZeroes 原地将所有 0 移到末尾，保持非零元素相对顺序。
func moveZeroes(nums []int) {
	// TODO: 实现
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
}
