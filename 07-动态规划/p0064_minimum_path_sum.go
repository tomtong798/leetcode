//go:build p0041

// LeetCode 第64题 - 最小路径和 (Minimum Path Sum)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组/动态规划/矩阵
//
// 【题目描述】
// 给定一个包含非负整数的 m x n 网格 grid，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。每次只能向下或者向右移动一步。
//
// 【示例 1】
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7（路径 1→3→1→1→1 总和最小）
// 【数据范围】
// m == grid.length；n == grid[i].length；1 <= m, n <= 200；0 <= grid[i][j] <= 200
//
// 【思路提示】dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])，首行首列只累加，可压缩成一维数组。
//
// 题目链接：https://leetcode.cn/problems/minimum-path-sum/
package main

// minPathSum 返回网格左上角到右下角的最小路径和。
func minPathSum(grid [][]int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7
}
