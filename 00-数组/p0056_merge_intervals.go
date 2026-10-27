//go:build p0041

// LeetCode 第56题 - 合并区间 (Merge Intervals)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、排序
//
// 【题目描述】
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [start_i, end_i]。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
// 【示例 1】
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠，将它们合并为 [1,6]。
//
// 【示例 2】
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 【数据范围】
// 1 <= intervals.length <= 10^4
// intervals[i].length == 2
// 0 <= start_i <= end_i <= 10^4
//
// 【思路提示】按左端点排序，依次遍历：若当前区间左端点 <= 上一合并区间的右端点则合并（取右端点较大值），否则新开一段。
//
// 题目链接：https://leetcode.cn/problems/merge-intervals/
package main

// mergeIntervals 合并所有重叠区间。
func mergeIntervals(intervals [][]int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	mergeIntervals([][]int{{1, 4}, {4, 5}})
}
