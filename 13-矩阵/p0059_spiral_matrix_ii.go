//go:build p0041

// LeetCode 第59题 - 螺旋矩阵 II (Spiral Matrix II)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】数组、矩阵、模拟
//
// 【题目描述】
// 给你一个正整数 n，生成一个包含 1 到 n^2 所有元素，并按顺时针顺序排列的 n x n 正方形矩阵 matrix。
//
// 【示例 1】
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 【示例 2】
// 输入：n = 1
// 输出：[[1]]
//
// 【数据范围】
// 1 <= n <= 20
//
// 【思路提示】与第 54 题同套路：四边界收缩，按 右→下→左→上 循环填数。
//
// 题目链接：https://leetcode.cn/problems/spiral-matrix-ii/
package main

// generateMatrix 生成 n x n 的螺旋矩阵。
func generateMatrix(n int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	generateMatrix(3)
	generateMatrix(1)
}
