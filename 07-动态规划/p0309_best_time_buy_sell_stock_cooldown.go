//go:build p0041

// LeetCode 第309题 - 最佳买卖股票时机含冷冻期 (Best Time to Buy and Sell Stock with Cooldown)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组/动态规划
//
// 【题目描述】
// 给定股票每日价格，可完成多笔交易，但卖出后有一天冷冻期（第二天才能买入），求最大利润。
//
// 【示例 1】
// 输入：prices = [1,2,3,0,2]
// 输出：3（[买, 卖, 冷冻, 买, 卖]）
// 【数据范围】
// 1 <= prices.length <= 5000；0 <= prices[i] <= 1000
//
// 【思路提示】三状态机：持有 / 未持有且冷冻 / 未持有且自由，每天按转移方程滚动更新。
//
// 题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
package main

// maxProfitCooldown 返回含冷冻期的多笔交易最大利润。
func maxProfitCooldown(prices []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(maxProfitCooldown([]int{1, 2, 3, 0, 2})) // 3
}
