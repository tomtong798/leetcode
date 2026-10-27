//go:build p0041

// LeetCode 第494题 - 目标和 (Target Sum)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组/动态规划/回溯
//
// 【题目描述】
// 给数组 nums 的每个数字前添加 '+' 或 '-'，使运算结果等于 target，返回方案数。
//
// 【示例 1】
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 【数据范围】
// 1 <= nums.length <= 20；0 <= nums[i] <= 1000；nums 所有元素之和 <= 1000
//
// 【思路提示】转化为 0-1 背包：sum(P) = (target+total)/2，求装满背包的方案数；(target+total) 为奇数或绝对值超界直接 0。
//
// 题目链接：https://leetcode.cn/problems/target-sum/
package main

// findTargetSumWays 返回添加正负号后结果等于 target 的方案数。
func findTargetSumWays(nums []int, target int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)) // 5
}
