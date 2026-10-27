//go:build p0041

// LeetCode 第53题 - 最大子数组和 (Maximum Subarray)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、动态规划、分治
//
// 【题目描述】
// 给你一个整数数组 nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），
// 返回其最大和。子数组是数组中的一个连续部分。
//
// 【示例 1】
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6。
//
// 【示例 2】
// 输入：nums = [1]
// 输出：1
//
// 【示例 3】
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
// 【数据范围】
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
//
// 【思路提示】Kadane 算法：dp[i] 表示以 i 结尾的最大子数组和，
// dp[i] = max(nums[i], dp[i-1]+nums[i])，答案为所有 dp[i] 的最大值，可压缩成一个变量。
//
// 题目链接：https://leetcode.cn/problems/maximum-subarray/
package main

// maxSubArray 返回最大连续子数组和。
func maxSubArray(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	maxSubArray([]int{1})
	maxSubArray([]int{5, 4, -1, 7, 8})
}
