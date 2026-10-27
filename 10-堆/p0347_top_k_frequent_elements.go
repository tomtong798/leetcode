//go:build p0041

// LeetCode 第347题 - 前 K 个高频元素 (Top K Frequent Elements)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、哈希表、分治、桶排序、计数、快速选择、堆
//
// 【题目描述】
// 给你一个整数数组 nums 和一个整数 k，请你返回其中出现频率前 k 高的元素。你可以按任意顺序返回答案。
// 进阶：你所设计算法的时间复杂度必须优于 O(n log n)，其中 n 是数组中元素的个数。
//
// 【示例 1】
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
//
// 【示例 2】
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 【数据范围】
// 1 <= nums.length <= 10^5
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一
//
// 【思路提示】哈希表统计频次后三种解法：
// 1) 小顶堆维护 k 个，O(n log k)；
// 2) 桶排序：按频次分桶，从高到低收集，O(n)；
// 3) 快速选择按频次分区。
//
// 题目链接：https://leetcode.cn/problems/top-k-frequent-elements/
package main

// topKFrequent 返回出现频率前 k 高的元素。
func topKFrequent(nums []int, k int) []int {
	// TODO: 实现
	return nil
}

func main() {
	topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	topKFrequent([]int{1}, 1)
}
