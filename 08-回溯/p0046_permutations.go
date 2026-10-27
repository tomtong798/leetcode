//go:build p0041

// LeetCode 第46题 - 全排列 (Permutations)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、回溯
//
// 【题目描述】
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列。你可以按任意顺序返回答案。
//
// 【示例 1】
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 【示例 2】
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
//
// 【示例 3】
// 输入：nums = [1]
// 输出：[[1]]
//
// 【数据范围】
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数互不相同
//
// 【思路提示】回溯模板题：用 used 数组标记已选元素；
// 路径长度等于 nums 长度时收集答案。务必默写熟练——回溯的"地基"。
//
// 题目链接：https://leetcode.cn/problems/permutations/
package main

// permute 返回所有全排列。
func permute(nums []int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	permute([]int{1, 2, 3})
	permute([]int{1})
}
