//go:build p0041

// LeetCode 第128题 - 最长连续序列 (Longest Consecutive Sequence)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、哈希表
//
// 【题目描述】
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列的长度。O(n)。
//
// 【示例 1】
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
//
// 【示例 2】
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 【思路提示】哈希集合，只从"无前驱"的数开始向右延伸。
//
// 题目链接：https://leetcode.cn/problems/longest-consecutive-sequence/
package main

// longestConsecutiveHash 哈希集合解法。
func longestConsecutiveHash(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	longestConsecutiveHash([]int{100, 4, 200, 1, 3, 2})
	longestConsecutiveHash([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}
