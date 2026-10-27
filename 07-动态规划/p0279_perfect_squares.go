//go:build p0041

// LeetCode 第279题 - 完全平方数 (Perfect Squares)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】广度优先搜索、动态规划、数学
//
// 【题目描述】
// 给你一个整数 n，返回和为 n 的完全平方数的最少数量。
// 完全平方数是一个整数，其值等于另一个整数的平方；
// 换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数。
//
// 【示例 1】
// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4。
//
// 【示例 2】
// 输入：n = 13
// 输出：2
// 解释：13 = 4 + 9。
//
// 【数据范围】
// 1 <= n <= 10^4
//
// 【思路提示】dp[i] = min(dp[i-j*j]) + 1，j 从 1 到 sqrt(i)；
// 也可以抽象成从 n 出发减平方数的最短路，用 BFS 解。
//
// 题目链接：https://leetcode.cn/problems/perfect-squares/
package main

// numSquares 返回和为 n 的完全平方数的最少数量。
func numSquares(n int) int {
	// TODO: 实现
	return 0
}

func main() {
	numSquares(12)
	numSquares(13)
}
