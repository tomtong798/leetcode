//go:build p0041

// LeetCode 第45题 - 跳跃游戏 II (Jump Game II)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】贪心/数组/动态规划
//
// 【题目描述】
// 给定非负整数数组 nums，你最初位于第一个下标，返回到达最后一个下标的最少跳跃次数（题目保证可以到达）。
//
// 【示例 1】
// 输入：nums = [2,3,1,1,4]
// 输出：2（先跳 1 步到下标 1，再跳 3 步到终点）
// 【数据范围】
// 1 <= nums.length <= 10^4；0 <= nums[i] <= 1000；题目保证可达
//
// 【思路提示】贪心分层（隐式 BFS）：维护当前层边界 curEnd 与下层最远 nextFar，i 到达 curEnd 时步数 +1。
//
// 题目链接：https://leetcode.cn/problems/jump-game-ii/
package main

// jump 返回到达最后一个下标的最少跳跃次数。
func jump(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(jump([]int{2, 3, 1, 1, 4})) // 2
}
