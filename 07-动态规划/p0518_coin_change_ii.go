//go:build p0041

// LeetCode 第518题 - 零钱兑换 II (Coin Change II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组/动态规划
//
// 【题目描述】
// 给你一个整数数组 coins 表示不同面额的硬币和一个整数 amount，每种硬币数量无限，计算凑成总金额的硬币组合数。
//
// 【示例 1】
// 输入：amount = 5, coins = [1,2,5]
// 输出：4（5=5、2+2+1、2+1+1+1、1+1+1+1+1）
// 【数据范围】
// 1 <= coins.length <= 300；1 <= coins[i] <= 5000；amount <= 5000；题目保证答案符合 32 位整数
//
// 【思路提示】完全背包求组合数：外层遍历硬币、内层遍历金额（顺序不能反，否则数成排列数）。
//
// 题目链接：https://leetcode.cn/problems/coin-change-ii/
package main

// change 返回凑成总金额的硬币组合数。
func change(amount int, coins []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(change(5, []int{1, 2, 5})) // 4
}
