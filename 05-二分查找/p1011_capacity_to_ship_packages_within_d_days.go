//go:build p0041

// LeetCode 第1011题 - 在 D 天内送达包裹的能力 (Capacity To Ship Packages Within D Days)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
// 传送带上的第 i 个包裹的重量为 weights[i]。每一天按顺序运送包裹，每天运载重量总和不超过船的载重。
// 返回能在 days 天内将所有包裹送达的最低载重。
//
// 【示例 1】
// 输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
// 输出：15
// 解释：载重 15 时：[1,2,3,4,5]、[6,7]、[8]、[9]、[10]，共 5 天。
//
// 【示例 2】
// 输入：weights = [3,2,2,4,1,4], days = 3
// 输出：6
//
// 【数据范围】
// 1 <= days <= weights.length <= 5 * 10^4
// 1 <= weights[i] <= 500
//
// 【思路提示】二分答案载重 [max(weights), sum(weights)]，
// 贪心判定：每天尽量装包裹，超载则换天。
//
// 题目链接：https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/
package main

// shipWithinDays 返回最低载重。
func shipWithinDays(weights []int, days int) int {
	// TODO: 实现
	return 0
}

func main() {
	shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3)
}
