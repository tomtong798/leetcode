//go:build p0041

// LeetCode 第561题 - 数组拆分 (Array Partition)
//
// 【难度】简单 | 【面试频率】★★★ | 【标签】贪心/数组/排序
//
// 【题目描述】
// 给定长度为 2n 的整数数组，将这些数分成 n 对 (a1,b1)...(an,bn)，使得每对取 min 后的总和最大。
//
// 【示例 1】
// 输入：nums = [1,4,3,2]
// 输出：4（最优分法 (1,2)+(3,4)，min 和为 1+3=4）
// 【数据范围】
// 1 <= n <= 10^4；nums.length == 2n；-10^4 <= nums[i] <= 10^4
//
// 【思路提示】排序后取偶数下标元素求和：让每对损失最小即配对相邻元素。
//
// 题目链接：https://leetcode.cn/problems/array-partition/
package main

// arrayPairSum 返回分成 n 对后每对取 min 的最大总和。
func arrayPairSum(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(arrayPairSum([]int{1, 4, 3, 2})) // 4
}
