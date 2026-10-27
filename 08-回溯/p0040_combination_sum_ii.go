//go:build p0041

// LeetCode 第40题 - 组合总和 II (Combination Sum II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、回溯
//
// 【题目描述】
// 给定一个候选人编号的集合 candidates 和一个目标数 target，
// 找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。解集不能包含重复的组合。
//
// 【示例 1】
// 输入: candidates = [10,1,2,7,6,1,5], target = 8
// 输出: [[1,1,6],[1,2,5],[1,7],[2,6]]
//
// 【示例 2】
// 输入: candidates = [2,5,2,1,2], target = 5
// 输出: [[1,2,2],[5]]
//
// 【数据范围】
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// 【思路提示】先排序，回溯时同层跳过重复 (i > start && c[i] == c[i-1])，每数只能用一次所以下层从 i+1。
//
// 题目链接：https://leetcode.cn/problems/combination-sum-ii/
package main

// combinationSum2 返回和为 target 的不重复组合。
func combinationSum2(candidates []int, target int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	combinationSum2([]int{2, 5, 2, 1, 2}, 5)
}
