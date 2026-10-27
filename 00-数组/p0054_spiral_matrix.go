//go:build p0041

// LeetCode 第54题 - 螺旋矩阵 (Spiral Matrix)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、矩阵、模拟
//
// 【题目描述】
// 给你一个 m 行 n 列的矩阵 matrix，请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
// 【示例 1】
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
//
// 【示例 2】
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 【数据范围】
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// 【思路提示】四个边界 left/right/top/bottom，循环按 右→下→左→上 收缩边界；
// 每走完一边要判断边界是否交叉，防止重复遍历。
//
// 题目链接：https://leetcode.cn/problems/spiral-matrix/
package main

// spiralOrder 按螺旋顺序返回矩阵所有元素。
func spiralOrder(matrix [][]int) []int {
	// TODO: 实现
	return nil
}

func main() {
	spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
}
