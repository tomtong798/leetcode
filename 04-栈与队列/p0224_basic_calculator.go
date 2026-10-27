//go:build p0041

// LeetCode 第224题 - 基本计算器 (Basic Calculator)
//
// 【难度】困难 | 【面试频率】★★★★ | 【标签】栈、递归、数学、字符串
//
// 【题目描述】
// 给你一个字符串表达式 s，请你实现一个基本计算器来计算并返回它的值。
// '+' 和 '-' 为运算符，'(' 和 ')' 用于分组，数字可能有多位。
// 你可以假设给定的表达式总是有效的。
//
// 【示例 1】
// 输入：s = "1 + 1"
// 输出：2
//
// 【示例 2】
// 输入：s = " 2-1 + 2 "
// 输出：3
//
// 【示例 3】
// 输入：s = "(1+(4+5+2)-3)+(6+8)"
// 输出：23
//
// 【数据范围】
// 1 <= s.length <= 3 * 10^5
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
//
// 【思路提示】用栈处理括号：
//   - 遇到数字：解析完整数字
//   - 遇到 '+'：sign = 1
//   - 遇到 '-'：sign = -1
//   - 遇到 '('：把当前 result 和 sign 压栈，重置 result = 0, sign = 1
//   - 遇到 ')'：弹出栈顶 sign，再弹出栈顶 result，result = 栈顶result + sign * 当前result
//
// 核心思想：括号内的结果乘以外层的符号，再加上外层之前累积的值。
//
// 题目链接：https://leetcode.cn/problems/basic-calculator/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate 计算表达式的值（支持加减和括号）。
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	isSymbol := func(s string) bool {
		return s == "+" || s == "-" || s == "*" || s == "/" || s == "(" || s == ")"
	}

	s = strings.Trim(s, " ")
	fmt.Println("s", s)

	// 提取 tokens, 处理多位数字
	var tokens []string
	var curNum string
	for i, ele := range s {
		if string(ele) == " " {
			continue
		}

		if isSymbol(string(ele)) {
			if curNum != "" {
				tokens = append(tokens, curNum)
				curNum = ""
			}
			tokens = append(tokens, string(ele))
			continue
		}
		curNum += string(ele)

		if i == len(s)-1 && curNum != "" {
			tokens = append(tokens, curNum)
		}
	}
	// ✅ 修复了多位数解析问题
	// 🟡 可以简化：循环结束后统一处理 curNum，不需要第 78-80 行的判断

	fmt.Println("tokens", tokens)
	fmt.Println("len tokens", len(tokens))

	priority := func(s string) int {
		if s == "+" || s == "-" {
			return 1
		}
		if s == "*" || s == "/" {
			return 2
		}
		return 0
	}

	// 中序转后序
	var result []string
	var stack []string
	for i, t := range tokens {
		if !isSymbol(t) {
			fmt.Println("1111111111  1")
			result = append(result, t)
			if i == len(tokens)-1 {
				fmt.Println("1111111111  1.1", stack)
				var n int = len(stack)
				for i := n - 1; i >= 0; i-- {
					result = append(result, stack[i])
				}
			}
			continue
		}

		if len(stack) == 0 {
			stack = append(stack, t)
			fmt.Println("1111111111  2", stack)
			continue
		}

		if t == "(" || stack[len(stack)-1] == "(" || priority(stack[len(stack)-1]) < priority(t) { // 🟡 条件太复杂，可以简化
			fmt.Println("1111111111  3-1", stack)
			stack = append(stack, t)
			fmt.Println("1111111111  3-2", stack)
			continue
		}

		if t == ")" || stack[len(stack)-1] != "(" && priority(stack[len(stack)-1]) >= priority(t) { // 🔴 逻辑有问题：运算符优先级判断不完整
			fmt.Println("1111111111  5-1", stack)

			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) != 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1] // 弹出 "("
			}

			if t != ")" {
				stack = append(stack, t) // 当前运算符入栈
			}

			fmt.Println("1111111111  5-2", stack)
		}
	}
	// ✅ 括号处理逻辑正确，但代码太复杂
	// 🔴 第 126 行条件逻辑容易出错，建议简化
	// 🔴 大量调试代码应删除

	fmt.Println("result", result)

	// 后序求值
	var nums []int
	for _, r := range result {
		if !isSymbol(r) {
			num, _ := strconv.Atoi(r)
			nums = append(nums, num)
			continue
		}

		var s int
		if r == "+" {
			s = nums[len(nums)-1] + nums[len(nums)-2]
		}
		if r == "-" {
			s = nums[len(nums)-2] - nums[len(nums)-1]
		}
		if r == "*" {
			s = nums[len(nums)-1] * nums[len(nums)-2]
		}
		if r == "/" {
			s = nums[len(nums)-2] / nums[len(nums)-1]
		}

		nums = nums[:len(nums)-2]
		nums = append(nums, s)
	}

	return nums[0]
}

// 【V0 点评 - 中缀转后缀法（支持括号）】
// ✅ 修复了多位数解析问题（第 60-81 行）
// ✅ 添加了括号处理逻辑（第 119-143 行）
// ✅ 中缀转后缀 + 后缀求值的整体思路正确
// ✅ 后缀求值逻辑正确
// ✅ 测试用例覆盖了多位数、括号、嵌套括号
// 🔴 大量调试代码：fmt.Println("1111111111 ...") 应删除
// 🟡 括号处理逻辑太复杂，第 126 行条件容易出错
// 🟡 第 78-80 行可以简化：循环结束后统一处理 curNum
// 🟡 思路是"中缀转后缀 + 后缀求值"，但 p0224 只有加减，不需要处理 * /
// 🟡 这个实现同时支持 + - * / 和括号，更适合 p0772（完整版计算器）

func main() {
	fmt.Println(calculate("1 +   12   "))                      // 13
	fmt.Println(calculate(" 2-1 + 2 "))                        // 3
	fmt.Println(calculate(" 21-1 * 2 "))                       // 19
	fmt.Println(calculate(" (21-1) * 2 +((3+1)*2+(40-20)*2)")) // 40+8+40=88
	// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) // 23
}
