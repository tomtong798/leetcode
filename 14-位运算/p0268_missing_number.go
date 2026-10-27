//go:build p0041

// LeetCode 第268题 - 丢失的数字 (Missing Number)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】位运算/数组/哈希表/数学
//
// 【题目描述】
// 给定一个包含 [0, n] 中 n 个数字的数组 nums，找出 [0, n] 范围内没有出现在数组中的那个数字。要求 O(n) 时间、O(1) 空间。
//
// 【示例 1】
// 输入：nums = [3,0,1]
// 输出：2
// 【数据范围】
// n == nums.length；1 <= n <= 10^4；0 <= nums[i] <= n；nums 中的数字互不相同
//
// 【思路提示】异或所有下标与所有值，缺失数字只出现一次被保留；或用等差求和减数组总和。
//
// 题目链接：https://leetcode.cn/problems/missing-number/
package main

// missingNumber 返回 [0, n] 中缺失的那个数字。
func missingNumber(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(missingNumber([]int{3, 0, 1})) // 2
}
