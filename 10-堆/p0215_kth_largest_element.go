//go:build p0041

// LeetCode 第215题 - 数组中的第K个最大元素 (Kth Largest Element in an Array)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、分治、快速选择、排序、堆
//
// 【题目描述】
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 【示例 1】
// 输入：[3,2,1,5,6,4], k = 2
// 输出：5
//
// 【示例 2】
// 输入：[3,2,3,1,2,4,5,5,6], k = 4
// 输出：4
//
// 【数据范围】
// 1 <= k <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
//
// 【思路提示】两种主流解法：
// 1) 快速选择（推荐掌握）：随机 pivot 分区，只递归包含第 k 大位置的一侧，平均 O(n)；
// 2) 小顶堆维护 k 个元素，遍历结束堆顶即答案，O(n log k)。
//
// 题目链接：https://leetcode.cn/problems/kth-largest-element-in-an-array/
package main

// findKthLargest 返回数组中第 k 个最大元素。
func findKthLargest(nums []int, k int) int {
	// TODO: 实现（建议练快速选择 + 堆两种写法）
	return 0
}

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
}
