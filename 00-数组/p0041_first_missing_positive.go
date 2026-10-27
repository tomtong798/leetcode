//go:build p0041

// LeetCode 第41题 - 缺失的第一个正数 (First Missing Positive)
//
// 【难度】困难 | 【面试频率】★★★★ | 【标签】数组、哈希表（原地哈希）
//
// 【题目描述】
// 给你一个未排序的整数数组 nums，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
// 【示例 1】
// 输入：nums = [1,2,0]
// 输出：3
//
// 【示例 2】
// 输入：nums = [3,4,-1,1]
// 输出：1
//
// 【示例 3】
// 输入：nums = [7,8,9,11,12]
// 输出：1
//
// 【数据范围】
// 1 <= nums.length <= 5 * 10^5
// -2^31 <= nums[i] <= 2^31 - 1
//
// 【思路提示】原地哈希：把值 v（1<=v<=n）交换到下标 v-1；
// 再扫一遍，第一个 nums[i] != i+1 的位置，答案即 i+1；都归位则答案为 n+1。
//
// 题目链接：https://leetcode.cn/problems/first-missing-positive/
package main

// firstMissingPositive 返回缺失的第一个正数，O(n) 时间 O(1) 空间。
func firstMissingPositive(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	firstMissingPositive([]int{1, 2, 0})
	firstMissingPositive([]int{3, 4, -1, 1})
	firstMissingPositive([]int{7, 8, 9, 11, 12})
}
