//go:build p0041

// LeetCode 第875题 - 爱吃香蕉的珂珂 (Koko Eating Bananas)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 珂珂喜欢香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。
// 警卫已经离开了，将在 h 小时后回来。珂珂可以决定吃香蕉的速度 k（根/小时）。
// 每小时选一堆香蕉吃掉 k 根；若堆中少于 k 根则吃掉这堆且这一小时不再吃。
// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k。
//
// 【示例 1】
// 输入：piles = [3,6,7,11], h = 8
// 输出：4
//
// 【示例 2】
// 输入：piles = [30,11,23,4,20], h = 5
// 输出：30
//
// 【数据范围】
// 1 <= piles.length <= 10^4
// piles.length <= h <= 10^9
// 1 <= piles[i] <= 10^9
//
// 【思路提示】二分答案速度 k ∈ [1, max(piles)]，判定 ceil(pile/k) 总和 <= h。
//
// 题目链接：https://leetcode.cn/problems/koko-eating-bananas/
package main

// minEatingSpeed 返回最小吃香蕉速度。
func minEatingSpeed(piles []int, h int) int {
	// TODO: 实现
	return 0
}

func main() {
	minEatingSpeed([]int{3, 6, 7, 11}, 8)
	minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)
}
