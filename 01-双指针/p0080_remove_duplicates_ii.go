//go:build p0041

// LeetCode 第80题 - 删除有序数组中的重复项 II (Remove Duplicates from Sorted Array II)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】数组、双指针
//
// 【题目描述】
// 给你一个有序数组 nums，请你原地删除重复出现的元素，使得出现次数超过两次的元素
// 只出现两次，返回删除后数组的新长度。不要使用额外的数组空间，
// 你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
// 【示例 1】
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
//
// 【示例 2】
// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
//
// 【数据范围】
// 1 <= nums.length <= 10^4
// -10^4 <= nums[i] <= 10^4
// nums 已按非严格递增排列
//
// 【思路提示】快慢指针：慢指针指向写入位置，当 nums[fast] != nums[slow-2] 时才写入
// （与"保留 k 位"通用解法：比较 nums[fast] 与 nums[slow-k]）。
//
// 题目链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
package main

// removeDuplicatesII 原地使每个元素最多出现两次并返回新长度。
func removeDuplicatesII(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	removeDuplicatesII([]int{1, 1, 1, 2, 2, 3})
	removeDuplicatesII([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
}
