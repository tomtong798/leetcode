//go:build p0041

// LeetCode 第22题 - 括号生成 (Generate Parentheses)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】字符串、动态规划、回溯
//
// 【题目描述】
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的括号组合。
//
// 【示例 1】
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
// 【示例 2】
// 输入：n = 1
// 输出：["()"]
//
// 【数据范围】
// 1 <= n <= 8
//
// 【思路提示】回溯 + 剪枝：维护已用的左括号数 left、右括号数 right；
// left < n 可加 '('，right < left 可加 ')'；长度达 2n 时收集。
//
// 题目链接：https://leetcode.cn/problems/generate-parentheses/
package main

import "fmt"

// generateParenthesis 生成所有合法括号组合。
func generateParenthesis(n int) []string {
	var kuohao = make([]string, 2)
	kuohao[0], kuohao[1] = "(", ")"

	var result []string
	var dfs func(left, right int, cur string)
	dfs = func(left, right int, cur string) {
		index := left + right - 1 // 🔴 BUG：第一次调用 left=0, right=0，index=-1
		if index == n*2 {         // 🔴 BUG：应该是 len(cur) == n*2 或 left+right == n*2
			result = append(result, cur)
			return
		}

		for i := 0; i < 2; i++ {
			if kuohao[i] == "(" {
				if left >= n {
					continue
				}
				left += 1
			}

			if kuohao[i] == ")" {
				if left <= right {
					continue
				}
				right += 1
			}

			cur = cur + kuohao[i]
			dfs(left, right, cur)

			if string(cur[index-1]) == "(" { // 🔴 BUG：index 可能是 -1 或 0，index-1 越界
				left--
			} else {
				right--
			}
			cur = cur[:index]
		}
	}

	dfs(0, 0, "")

	return result
}

// 【V0 点评】
// 编译通过，但运行会 panic：
//   🔴 第 37 行：index = left + right - 1，第一次调用时 index = -1
//      应该用 len(cur) 或 left + right 表示当前长度
//   🔴 第 38 行：if index == n*2 判断错误
//      应该是 if len(cur) == n*2 或 if left+right == n*2
//   🔴 第 61 行：cur[index-1] 在 index <= 0 时越界 panic
//   🔴 回溯逻辑混乱：应该根据当前添加的字符恢复 left/right，
//      而不是根据 cur[index-1] 猜测
//   🟡 用 kuohao 数组 + for 循环遍历 "(" 和 ")" 过于复杂，
//      直接写两个 if 分支更清晰
//   🟡 字符串拼接 cur = cur + kuohao[i] 每次分配新内存
//   ✅ 剪枝条件正确：left >= n 不加左括号，left <= right 不加右括号

func generateParenthesis11111(n int) []string {
	var kuohao = make([]string, 2)
	kuohao[0], kuohao[1] = "(", ")"

	var result []string
	var dfs func(left, right int, cur string)
	dfs = func(left, right int, cur string) {
		index := left + right
		if index == n*2 {
			result = append(result, cur)
			return
		}

		for i := 0; i < 2; i++ {
			if kuohao[i] == "(" {
				if left >= n {
					continue
				}
				left += 1
				cur = cur + kuohao[i]
				dfs(left, right, cur)
				left--
				cur = cur[:index]
			}

			if kuohao[i] == ")" {
				if left <= right {
					continue
				}
				right += 1
				cur = cur + kuohao[i]
				dfs(left, right, cur)
				right--
				cur = cur[:index]
			}
		}
	}

	dfs(0, 0, "")

	return result
}

func main() {
	// fmt.Println("V0:", generateParenthesis(3))
	// fmt.Println("V1:", generateParenthesisV1(3))
	fmt.Println("V2:", generateParenthesisV2(3))
	fmt.Println("generateParenthesis11111", generateParenthesis11111(3))
}

// generateParenthesisV1 简洁版：直接写两个 if 分支
func generateParenthesisV1(n int) []string {
	var result []string

	var dfs func(left, right int, cur string)
	dfs = func(left, right int, cur string) {
		if len(cur) == n*2 { // ✅ 终止条件：长度达到 2n
			result = append(result, cur)
			return
		}

		if left < n { // ✅ 可以加左括号
			dfs(left+1, right, cur+"(")
		}
		if right < left { // ✅ 可以加右括号（剪枝：右括号数必须小于左括号数）
			dfs(left, right+1, cur+")")
		}
	}

	dfs(0, 0, "")
	return result
}

// 【V1 点评 - 简洁版】
// ✅ 终止条件正确：len(cur) == n*2
// ✅ 剪枝条件正确：left < n 加左括号，right < left 加右括号
// ✅ 逻辑清晰：直接写两个 if 分支，不用 for 循环
// ✅ 无需手动回溯：Go 字符串不可变，递归调用时传新字符串，自动回溯
// 🟡 字符串拼接 cur+"(" 每次分配新内存，性能一般
// 面试推荐写法（思路清晰，容易理解）

// generateParenthesisV2 优化版：用 byte slice 减少内存分配
func generateParenthesisV2(n int) []string {
	var result []string
	path := make([]byte, 0, 2*n) // ✅ 预分配容量

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == n && right == n { // ✅ 终止条件：左右括号都用完
			result = append(result, string(path))
			return
		}

		if left < n {
			path = append(path, '(')
			dfs(left+1, right)
			path = path[:len(path)-1] // ✅ 回溯：弹出最后一个字符
		}
		if right < left {
			path = append(path, ')')
			dfs(left, right+1)
			path = path[:len(path)-1] // ✅ 回溯：弹出最后一个字符
		}
	}

	dfs(0, 0)
	return result
}

// 【V2 点评 - 优化版】
// ✅ 用 byte slice 代替字符串拼接，避免重复分配内存
// ✅ 预分配容量 2*n，减少扩容开销
// ✅ 回溯逻辑清晰：path = path[:len(path)-1] 弹出最后一个字符
// ✅ 终止条件：left == n && right == n（左右括号都用完）
// 性能最优写法，适合对性能有要求的场景

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法              | 时间        | 空间      | 面试期望        |
// |-------------------|-------------|-----------|-----------------|
// | 回溯（V0 有 bug） | O(4^n/√n)   | O(n)      | ❌ 会 panic      |
// | 回溯简洁版（V1）   | O(4^n/√n)   | O(n)      | ✅ 推荐写法      |
// | 回溯优化版（V2）   | O(4^n/√n)   | O(n)      | ✅ 性能最优      |
// | 动态规划          | O(4^n/√n)   | O(4^n)    | ⚠️ 可以但不如回溯 |
//
// 注：卡特兰数 C(n) = C(2n, n) / (n+1) ≈ 4^n / (n^(3/2) * √π)
//
// 【剪枝策略】
// 这道题**需要剪枝**，因为有约束条件：
//   1. 左括号数不能超过 n
//   2. 右括号数不能大于左括号数（否则无法匹配）
//
// 对比不需要剪枝的题目（如电话号码组合）：
//   - 电话号码组合：无约束，所有路径都有效
//   - 括号生成：有约束，必须满足匹配条件
//
// 【面试讲解建议】
// 1. 先说思路："用回溯枚举所有可能的括号组合，加剪枝条件"
// 2. 画出递归树，标注剪枝的位置
// 3. 解释为什么 right < left 才能加右括号
// 4. 分析复杂度：时间 O(4^n/√n)，空间 O(n)
// ──────────────────────────────────────────────────────────
