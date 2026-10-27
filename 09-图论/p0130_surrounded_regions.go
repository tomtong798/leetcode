//go:build p0041

// LeetCode 第130题 - 被围绕的区域 (Surrounded Regions)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】深度优先搜索、广度优先搜索、并查集、数组、矩阵
//
// 【题目描述】
// 给你一个 m x n 的矩阵 board，由若干字符 'X' 和 'O' 组成，
// 找出所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 被围绕的区间不会存在于边界上。
//
// 【示例 1】
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// 解释：边界上的 'O' 不会被围绕，其余被围绕的 'O' 被填充为 'X'。
//
// 【数据范围】
// m == board.length, n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
// 【思路提示】从四边界 DFS 把与边界连通的 'O' 临时标记为 '#',
// 再遍历：'#' 还原为 'O'，剩余 'O' 变 'X'。
//
// 题目链接：https://leetcode.cn/problems/surrounded-regions/
package main

// solve 被围绕的区域填充。
func solve(board [][]byte) {
	// TODO: 实现
}

func main() {
	solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}
