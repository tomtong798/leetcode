//go:build p0041

// LeetCode 第152题 - 乘积最大子数组 (Maximum Product Subarray)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、动态规划
//
// 【题目描述】
// 给你一个整数数组 nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），
// 并返回该子数组所对应的乘积。测试用例的答案是一个 32 位整数。
//
// 【示例 1】
// 输入：nums = [2,3,-2,4]
// 输出：6
// 解释：子数组 [2,3] 有最大乘积 6。
//
// 【示例 2】
// 输入：nums = [-2,0,-1]
// 输出：0
// 解释：结果不能为 2，因为 [-2,-1] 不是子数组。
//
// 【数据范围】
// 1 <= nums.length <= 2 * 10^4
// -10 <= nums[i] <= 10
// 任何子数组的乘积都保证是 32 位整数
//
// 【思路提示】同时维护以 i 结尾的最大乘积 imax 和最小乘积 imin（负数乘负数会变最大）；
// 遇到负数交换 imax/imin 再转移。与"最大子数组和"对照着记。
//
// 题目链接：https://leetcode.cn/problems/maximum-product-subarray/
package main

// maxProduct 返回乘积最大的非空连续子数组乘积。
func maxProduct(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	maxProduct([]int{2, 3, -2, 4})
	maxProduct([]int{-2, 0, -1})
}
