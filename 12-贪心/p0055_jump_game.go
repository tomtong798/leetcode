//go:build p0041

// LeetCode 第55题 - 跳跃游戏 (Jump Game)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】贪心/数组/动态规划
//
// 【题目描述】
// 给定非负整数数组 nums，你最初位于第一个下标，数组中每个元素代表该位置可跳跃的最大长度，判断你是否能够到达最后一个下标。
//
// 【示例 1】
// 输入：nums = [2,3,1,1,4]
// 输出：true
//
// 【示例 2】
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 【数据范围】
// 1 <= nums.length <= 3*10^4；0 <= nums[i] <= 10^5
//
// 【思路提示】贪心维护最远可达 maxReach = max(maxReach, i+nums[i])，若 i > maxReach 则返回 false。
//
// 题目链接：https://leetcode.cn/problems/jump-game/
package main

// canJump 判断能否到达数组最后一个下标。
func canJump(nums []int) bool {
	// TODO: 实现
	return false
}

func main() {
	println(canJump([]int{2, 3, 1, 1, 4})) // true
	println(canJump([]int{3, 2, 1, 0, 4})) // false
}
