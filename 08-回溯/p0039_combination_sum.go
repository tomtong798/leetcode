//go:build p0041

// LeetCode 第39题 - 组合总和 (Combination Sum)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、回溯
//
// 【题目描述】
// 给你一个无重复元素的整数数组 candidates 和一个目标整数 target，
// 找出 candidates 中可以使数字和为目标数 target 的所有不同组合，并以列表形式返回。
// 你可以按任意顺序返回这些组合。candidates 中的同一个数字可以无限制重复被选取。
// 如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 【示例 1】
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：2 和 3 可以组成一组候选，2 + 2 + 3 = 7。注意 2 可以使用多次。
//
// 【示例 2】
// 输入：candidates = [2,3,5], target = 8
// 输出：[[2,2,2,2],[2,3,3],[3,5]]
//
// 【示例 3】
// 输入：candidates = [2], target = 1
// 输出：[]
//
// 【数据范围】
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// candidates 的所有元素互不相同
// 1 <= target <= 40
//
// 【思路提示】回溯 + 剪枝：递归时传入 start 下标（允许重复选自己，所以下一层仍从 i 开始）；
// 剩余 target < 0 剪枝，== 0 收集答案。先排序可提前 break。
//
// 题目链接：https://leetcode.cn/problems/combination-sum/
package main

// combinationSum 返回所有和为 target 的组合。
func combinationSum(candidates []int, target int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
	combinationSum([]int{2, 3, 5}, 8)
	combinationSum([]int{2}, 1)
}
