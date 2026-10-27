//go:build p0041

// LeetCode 第118题 - 杨辉三角 (Pascal's Triangle)
//
// 【难度】简单 | 【面试频率】★★★ | 【标签】数组、动态规划
//
// 【题目描述】
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
// 在「杨辉三角」中，每个数是它左上方和上方的数的和。
//
// 【示例 1】
// 输入：numRows = 5
// 输出：[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
// 【示例 2】
// 输入：numRows = 1
// 输出：[[1]]
//
// 【数据范围】
// 1 <= numRows <= 30
//
// 【思路提示】每行首尾为 1，中间元素 = 上一行相邻两元素之和。
//
// 题目链接：https://leetcode.cn/problems/pascals-triangle/
package main

// generate 生成杨辉三角前 numRows 行。
func generate(numRows int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	generate(5)
	generate(1)
}
