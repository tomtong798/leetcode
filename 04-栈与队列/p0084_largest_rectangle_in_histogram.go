//go:build p0041

// LeetCode 第84题 - 柱状图中最大的矩形 (Largest Rectangle in Histogram)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】栈、数组、单调栈
//
// 【题目描述】
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。
// 每个柱子彼此相邻，且宽度为 1。求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 【示例 1】
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10。
//
// 【示例 2】
// 输入：heights = [2,4]
// 输出：4
//
// 【数据范围】
// 1 <= heights.length <= 10^5
// 0 <= heights[i] <= 10^4
//
// 【思路提示】单调递增栈：弹出时以弹出柱高为高度、
// 当前下标与栈内新栈顶为左右边界，宽度 = i - stack[top] - 1。
//
// 题目链接：https://leetcode.cn/problems/largest-rectangle-in-histogram/
package main

// largestRectangleArea 返回柱状图最大矩形面积。
func largestRectangleArea(heights []int) int {
	// TODO: 实现
	return 0
}

func main() {
	largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	largestRectangleArea([]int{2, 4})
}
