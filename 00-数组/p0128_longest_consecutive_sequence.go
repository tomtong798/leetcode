//go:build p0041

// LeetCode 第128题 - 最长连续序列 (Longest Consecutive Sequence)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、哈希表、并查集
//
// 【题目描述】
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 【示例 1】
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 【示例 2】
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 【数据范围】
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
//
// 【思路提示】哈希集合存所有数；只从"没有前驱 x-1 的数"开始向右延伸计数，
// 保证每个数最多被访问两次，整体 O(n)。
//
// 题目链接：https://leetcode.cn/problems/longest-consecutive-sequence/
package main

// longestConsecutive 返回最长连续序列长度，O(n)。
func longestConsecutive(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	longestConsecutive(nil)
}
