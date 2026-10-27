//go:build p0041

// LeetCode 第240题 - 搜索二维矩阵 II (Search a 2D Matrix II)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】二分查找、分治、矩阵
//
// 【题目描述】
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。
// 该矩阵具有以下特性：
//   - 每行的元素从左到右升序排列。
//   - 每列的元素从上到下升序排列。
//
// 【示例 1】
// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// 输出：true
//
// 【示例 2】
// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// 输出：false
//
// 【数据范围】
// m == matrix.length, n == matrix[i].length
// 1 <= n, m <= 300
// -10^9 <= matrix[i][j] <= 10^9
//
// 【思路提示】Z 字形：从右上角出发，当前值 > target 左移，< target 下移；
// 每步淘汰一行或一列，O(m+n)。
//
// 题目链接：https://leetcode.cn/problems/search-a-2d-matrix-ii/
package main

// searchMatrix240 在行列均升序的矩阵中查找 target。
func searchMatrix240(matrix [][]int, target int) bool {
	// TODO: 实现
	return false
}

func main() {
	searchMatrix240([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5)
}
