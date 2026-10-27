//go:build p0041

// LeetCode 第213题 - 打家劫舍 II (House Robber II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组/动态规划
//
// 【题目描述】
// 你是一个专业的小偷，计划偷窃沿街的房屋。所有房屋围成一圈（首尾相邻），相邻两户不能同时偷窃，求能偷窃到的最高金额。
//
// 【示例 1】
// 输入：nums = [2,3,2]
// 输出：3
//
// 【示例 2】
// 输入：nums = [1,2,3,1]
// 输出：4
// 【数据范围】
// 1 <= nums.length <= 100；0 <= nums[i] <= 1000
//
// 【思路提示】环拆成两条链：不偷首户（0..n-2）与不偷尾户（1..n-1），各做一次线形打家劫舍取 max。
//
// 题目链接：https://leetcode.cn/problems/house-robber-ii/
package main

// rob213 返回环形房屋能偷窃到的最高金额。
func rob213(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(rob213([]int{2, 3, 2})) // 3
	println(rob213([]int{1, 2, 3, 1})) // 4
}
