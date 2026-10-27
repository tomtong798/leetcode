//go:build p0041

// LeetCode 第42题 - 接雨水 (Trapping Rain Water) —— 双指针解法
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】双指针、单调栈、动态规划
//
// 【题目描述】
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，
// 计算按此排列的柱子，下雨之后能接多少雨水。
//
// 【示例 1】
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
//
// 【示例 2】
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 【数据范围】
// 1 <= height.length <= 2 * 10^4
// 0 <= height[i] <= 10^5
//
// 【思路提示】见 00-数组/p0042。双指针版本：左右指针向中间移动，
// 维护左侧最大值 leftMax 与右侧最大值 rightMax，较小的一侧当前位置的水量已可确定。
//
// 题目链接：https://leetcode.cn/problems/trapping-rain-water/
package main

// trapTwoPointers 双指针 O(1) 空间解法。
func trapTwoPointers(height []int) int {
	// TODO: 实现
	return 0
}

func main() {
	trapTwoPointers([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	trapTwoPointers([]int{4, 2, 0, 3, 2, 5})
}
