//go:build p0041

// LeetCode 第209题 - 长度最小的子数组 (Minimum Size Subarray Sum)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、二分查找、前缀和、滑动窗口
//
// 【题目描述】
// 给定一个含有 n 个正整数的数组和一个正整数 target。
// 找出该数组中满足其和 ≥ target 的长度最小的连续子数组，并返回其长度。
// 如果不存在符合条件的子数组，返回 0。
//
// 【示例 1】
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
// 【示例 2】
// 输入：target = 4, nums = [1,4,4]
// 输出：1
//
// 【示例 3】
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 【数据范围】
// 1 <= target <= 10^9
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^4
//
// 【思路提示】滑动窗口：右指针扩张累加，和 ≥ target 时收缩左指针并更新最小长度。
// 进阶 O(n log n)：前缀和 + 二分。
//
// 题目链接：https://leetcode.cn/problems/minimum-size-subarray-sum/
package main

// minSubArrayLen 返回和 ≥ target 的最短连续子数组长度。
func minSubArrayLen(target int, nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	minSubArrayLen(4, []int{1, 4, 4})
	minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
}
