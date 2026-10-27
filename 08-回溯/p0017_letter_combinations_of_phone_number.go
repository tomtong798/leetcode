//go:build p0041

// LeetCode 第17题 - 电话号码的字母组合 (Letter Combinations of a Phone Number)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】哈希表、字符串、回溯
//
// 【题目描述】
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
// 数字到字母的映射与电话按键相同（2→abc, 3→def, 4→ghi, 5→jkl, 6→mno, 7→pqrs, 8→tuv, 9→wxyz）。
//
// 【示例 1】
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 【示例 2】
// 输入：digits = ""
// 输出：[]
//
// 【示例 3】
// 输入：digits = "2"
// 输出：["a","b","c"]
//
// 【数据范围】
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字
//
// 【思路提示】回溯：每个数字选一个字母，递归到末尾收集。空串返回空列表。
//
// 题目链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
package main

import "fmt"

// letterCombinations 返回电话号码所有字母组合。
func letterCombinations(digits string) []string {
	var n = len(digits)
	if n == 0 {
		return nil
	}

	var digitToKey = make(map[string]string) // 🟡 每次调用都重建，应提取为包级变量
	digitToKey["2"] = "abc"
	digitToKey["3"] = "def"
	digitToKey["4"] = "ghi"
	digitToKey["5"] = "jkl"
	digitToKey["6"] = "mno"
	digitToKey["7"] = "pqrs"
	digitToKey["8"] = "tuv"
	digitToKey["9"] = "wxyz"
	// 🟡 用 map[string]string 查找效率低，建议用数组：
	// var digitToLetters = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var result []string
	var dfs func(index int, curS string)

	// index 当前第几个按键， 也表示按键完成之后有几个字母

	dfs = func(index int, curS string) {
		var digit = string(digits[index])
		var keys = digitToKey[digit]
		var keyCount = len(keys)
		for i := 0; i < keyCount; i++ {
			curS = curS + string(keys[i])
			if index == n-1 {
				result = append(result, curS)
			} else {
				dfs(n+1, curS) // 🔴 BUG：应该是 dfs(index+1, curS)
				// n 是外层变量（digits 长度），永远是 2
				// 所以调用 dfs(3, "a")，导致 digits[3] 越界 panic
			}

			curS = curS[:index] // ✅ 回溯逻辑正确：截断到 index 位置
		}
	}

	dfs(0, "")

	return result
}

// 【V0 点评】
// 编译通过，但运行会 panic：
//   🔴 第 65 行：dfs(n+1, curS) 应该是 dfs(index+1, curS)
//      n 是外层变量（digits 长度），导致递归到错误位置，数组越界
//   🟡 digitToKey 每次调用都重建，应提取为包级变量
//   🟡 用 map[string]string 查找效率低，建议用数组索引
//   🟡 变量名 digitToKey 不准确，应叫 digitToLetters
//   ✅ 回溯逻辑 curS = curS[:index] 是正确的
//   ✅ 空串返回 nil 的处理正确

func main() {
	fmt.Println("V0:", letterCombinations("23"))
	fmt.Println("V1:", letterCombinationsV1("23"))
}

// digitToLetters 包级变量，避免每次调用重建
var digitToLetters = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// letterCombinationsV1 优化版：用数组代替 map，用 byte slice 减少字符串拼接
func letterCombinationsV1(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var result []string
	path := make([]byte, len(digits)) // 用 byte slice 代替字符串拼接

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			result = append(result, string(path))
			return
		}

		digit := digits[index] - '0'     // 字符转数字
		letters := digitToLetters[digit] // 数组索引 O(1)
		for i := 0; i < len(letters); i++ {
			path[index] = letters[i] // 直接写入对应位置
			dfs(index + 1)           // ✅ 正确的递归参数
		}
	}

	dfs(0)
	return result
}

// 【V1 点评 - 优化版】
// ✅ 修复了递归参数：dfs(index + 1)
// ✅ 用数组代替 map：digitToLetters[digit] 直接索引，O(1)
// ✅ 用 byte slice 代替字符串拼接：path[index] = letter，避免重复分配
// ✅ 映射表提取为包级变量，避免每次调用重建
// ✅ 终止条件更清晰：index == len(digits) 时收集结果
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法              | 时间          | 空间        | 面试期望    |
// |-------------------|---------------|-------------|-------------|
// | 回溯（V0 有 bug） | O(4^n * n)    | O(n)        | ❌ 会 panic  |
// | 回溯优化（V1）     | O(4^n * n)    | O(n)        | ✅ 推荐写法  |
// | 迭代/BFS          | O(4^n * n)    | O(4^n)      | ⚠️ 可以但不如回溯 |
//
// 注：n 是 digits 长度，4 是单个数字最多对应字母数（7/9 对应 4 个字母）
//
// 【面试讲解建议】
// 1. 先说思路："每个数字选一个字母，递归枚举所有组合"
// 2. 画出递归树帮助理解
// 3. 提到用数组代替 map 提升效率
// 4. 分析复杂度：时间 O(4^n * n)，空间 O(n)
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【关于剪枝的说明】
//
// 这道题**不需要剪枝**，因为题目要求是**枚举所有可能的组合**，每个分支都必须探索到底。
//
// 剪枝的前提：存在约束条件，某些分支不可能产生有效解，可以提前终止。
// 这道题：每个数字都必须选一个字母，没有约束条件，所有组合都是有效的。
//
// 对比需要剪枝的回溯题：
//
// | 题目              | 约束条件                  | 剪枝时机                      |
// |-------------------|---------------------------|-------------------------------|
// | 17. 电话号码组合   | 无约束，全部有效           | ❌ 不需要剪枝                  |
// | 39. 组合总和       | 和必须等于 target          | ✅ 当前和 > target 时剪枝      |
// | 51. N 皇后         | 不能同行同列同对角线       | ✅ 放置冲突时剪枝              |
// | 46. 全排列         | 元素不能重复使用           | ✅ 已使用的元素跳过            |
//
// 举例：组合总和（target=7, candidates=[2,3,6,7]）
//   搜索路径：[2, 3, 6]，当前和 = 11 > 7 → 剪枝！直接回溯
//   搜索路径：[2, 2, 3]，当前和 = 7 == target → 收集结果 ✅
//
// 这道题是**无条件枚举**，所有路径都有效，所以没有可剪的枝。
// ──────────────────────────────────────────────────────────
