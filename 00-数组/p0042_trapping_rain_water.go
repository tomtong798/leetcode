//go:build p0041

// LeetCode 第42题 - 接雨水 (Trapping Rain Water)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】数组、双指针、单调栈、动态规划
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
// 【思路提示】每个位置能接的水 = min(左边最高, 右边最高) - 自身高度。
// 双指针解法：左右指针向中间走，哪侧最高柱更矮就先算哪侧，空间 O(1)。
//
// 题目链接：https://leetcode.cn/problems/trapping-rain-water/
package main

// trap 返回能接的雨水总量。
func trap(height []int) int {
	// TODO: 实现
	return 0
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	trap([]int{4, 2, 0, 3, 2, 5})
}
