//go:build p0041

// LeetCode 第322题 - 零钱兑换 (Coin Change)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】广度优先搜索、数组、动态规划
//
// 【题目描述】
// 给你一个整数数组 coins，表示不同面额的硬币；以及一个整数 amount，表示总金额。
// 计算并返回可以凑成总金额所需的最少硬币个数。如果没有任何一种硬币组合能组成总金额，
// 返回 -1。你可以认为每种硬币的数量是无限的。
//
// 【示例 1】
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1。
//
// 【示例 2】
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 【示例 3】
// 输入：coins = [1], amount = 0
// 输出：0
//
// 【数据范围】
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2^31 - 1
// 0 <= amount <= 10^4
//
// 【思路提示】完全背包：dp[i] = min(dp[i], dp[i-coin]+1)，外层遍历金额、内层遍历硬币；
// 初始化 dp[0]=0，其余为无穷大。
//
// 题目链接：https://leetcode.cn/problems/coin-change/
package main

// coinChange 返回凑出 amount 的最少硬币数。
func coinChange(coins []int, amount int) int {
	// TODO: 实现
	return -1
}

func main() {
	coinChange([]int{1, 2, 5}, 11)
	coinChange([]int{2}, 3)
	coinChange([]int{1}, 0)
}
