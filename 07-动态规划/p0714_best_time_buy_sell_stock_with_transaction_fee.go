//go:build p0041

// LeetCode 第714题 - 买卖股票的最佳时机含手续费 (Best Time to Buy and Sell Stock with Transaction Fee)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】数组/动态规划/贪心
//
// 【题目描述】
// 给定股票每日价格和整数手续费，可无限次交易，每笔交易（买入+卖出）只需付一次手续费，求最大利润。
//
// 【示例 1】
// 输入：prices = [1,3,2,8,4,9], fee = 2
// 输出：8
// 【数据范围】
// 1 <= prices.length <= 5*10^4；1 <= prices[i], fee < 5*10^4
//
// 【思路提示】两状态：hold（持有）与 free（未持有），卖出时扣一次 fee，滚动转移即可。
//
// 题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
package main

// maxProfitFee 返回含手续费的多笔交易最大利润。
func maxProfitFee(prices []int, fee int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(maxProfitFee([]int{1, 3, 2, 8, 4, 9}, 2)) // 8
}
