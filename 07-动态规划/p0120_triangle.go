//go:build p0041

// LeetCode 第120题 - 三角形最小路径和 (Triangle)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组/动态规划
//
// 【题目描述】
// 给定一个三角形 triangle，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的节点上（即下一行的下标 i 或 i+1）。
//
// 【示例 1】
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11（路径 2+3+5+1=11）
// 【数据范围】
// 1 <= triangle.length <= 200；triangle[0].length == 1；triangle[i].length == triangle[i-1].length+1；-10^4 <= triangle[i][j] <= 10^4
//
// 【思路提示】自底向上滚动数组：dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]，最终 dp[0] 即答案。
//
// 题目链接：https://leetcode.cn/problems/triangle/
package main

// minimumTotal 返回三角形自顶向下的最小路径和。
func minimumTotal(triangle [][]int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) // 11
}
