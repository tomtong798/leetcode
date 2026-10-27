//go:build p0041

// LeetCode 第121题 - 买卖股票的最佳时机 (Best Time to Buy and Sell Stock)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】数组、动态规划、贪心
//
// 【题目描述】
// 给定一个数组 prices，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。返回你可以从这笔交易中获取的最大利润。
// 如果你不能获取任何利润，返回 0。
//
// 【示例 1】
// 输入：prices = [7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（价格=1）买入，在第 5 天（价格=6）卖出，最大利润 = 6-1 = 5。
//
// 【示例 2】
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下不进行交易，所以最大利润为 0。
//
// 【数据范围】
// 1 <= prices.length <= 10^5
// 0 <= prices[i] <= 10^4
//
// 【思路提示】一次遍历：维护到当前天为止的最低买入价，每天尝试以当天价格卖出，更新最大利润。
//
// 题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
package main

// maxProfit 返回一次交易的最大利润。
func maxProfit(prices []int) int {
	// TODO: 实现
	return 0
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
	maxProfit([]int{7, 6, 4, 3, 1})
}
