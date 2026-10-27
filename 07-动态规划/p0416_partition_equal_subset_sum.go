//go:build p0041

// LeetCode 第416题 - 分割等和子集 (Partition Equal Subset Sum)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、动态规划（0-1 背包）
//
// 【题目描述】
// 给你一个只包含正整数的非空数组 nums。请你判断是否可以将这个数组分割成两个子集，
// 使得两个子集的元素和相等。
//
// 【示例 1】
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11]。
//
// 【示例 2】
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
//
// 【数据范围】
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
// 【思路提示】总和为奇数直接 false；问题转化为"能否选出子集和为 sum/2"——0-1 背包可行性：
// 一维布尔数组 dp，倒序遍历容量：dp[j] |= dp[j-nums[i]]。
//
// 题目链接：https://leetcode.cn/problems/partition-equal-subset-sum/
package main

// canPartition 判断能否分割成两个和相等的子集。
func canPartition(nums []int) bool {
	// TODO: 实现
	return false
}

func main() {
	canPartition([]int{1, 5, 11, 5})
	canPartition([]int{1, 2, 3, 5})
}
