//go:build p0041

// LeetCode 第772题 - 基本计算器 III (Basic Calculator III)
//
// 【难度】困难 | 【面试频率】★★★ | 【标签】栈、递归、数学、字符串
//
// 【题目描述】
// 实现一个基本的计算器来计算简单的表达式字符串。
// 表达式字符串只包含非负整数，算符 +、-、*、/ ，左括号 ( 和右括号 )。
// 整数除法需要向下截断。
//
// 【示例 1】
// 输入：s = "1+1"
// 输出：2
//
// 【示例 2】
// 输入：s = "6-4/2"
// 输出：4
//
// 【示例 3】
// 输入：s = "2*(5+5*2)/3+(6/2+8)"
// 输出：21
//
// 【数据范围】
// 1 <= s.length <= 10^4
// s 由数字、'+'、'-'、'*'、'/'、'('、')' 和 ' ' 组成
// s 是一个有效的表达式
//
// 【思路提示】综合 224 和 227：
//   - 用两个栈：nums 存数字，ops 存运算符
//   - 优先级：( 最低，+ - 中等，* / 最高
//   - 遇到数字：解析完整数字入栈
//   - 遇到 (：入栈
//   - 遇到 )：计算直到 (
//   - 遇到运算符：按优先级计算后入栈
//
// 题目链接：https://leetcode.cn/problems/basic-calculator-iii/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate 计算表达式的值（加减乘除 + 括号）。
// calculate 计算表达式的值（支持加减和括号）。
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	isSymbol := func(s string) bool { // 🟡 可以简化，不需要 isSymbol
		return s == "+" || s == "-" || s == "*" || s == "/" || s == "(" || s == ")"
	}

	s = strings.Trim(s, " ")
	fmt.Println("s", s) // 🔴 调试代码，应删除

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

		if i == len(s)-1 && curNum != "" { // 🟡 可以简化：循环结束后统一处理
			tokens = append(tokens, curNum)
		}
	}
	// ✅ 修复了多位数解析问题
	// 🟡 可以简化：循环结束后统一处理 curNum，不需要第 78-80 行的判断

	fmt.Println("tokens", tokens) // 🔴 调试代码，应删除
	fmt.Println("len tokens", len(tokens)) // 🔴 调试代码，应删除

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
			fmt.Println("1111111111  1") // 🔴 调试代码，应删除
			result = append(result, t)
			if i == len(tokens)-1 {
				fmt.Println("1111111111  1.1", stack) // 🔴 调试代码，应删除
				var n int = len(stack)
				for i := n - 1; i >= 0; i-- {
					result = append(result, stack[i])
				}
			}
			continue
		}

		if len(stack) == 0 {
			stack = append(stack, t)
			fmt.Println("1111111111  2", stack) // 🔴 调试代码，应删除
			continue
		}

		if t == "(" || stack[len(stack)-1] == "(" || priority(stack[len(stack)-1]) < priority(t) { // 🟡 条件太复杂，可以简化
			fmt.Println("1111111111  3-1", stack) // 🔴 调试代码，应删除
			stack = append(stack, t)
			fmt.Println("1111111111  3-2", stack) // 🔴 调试代码，应删除
			continue
		}

		if t == ")" || stack[len(stack)-1] != "(" && priority(stack[len(stack)-1]) >= priority(t) { // 🔴 逻辑有问题：运算符优先级判断不完整
			fmt.Println("1111111111  5-1", stack) // 🔴 调试代码，应删除

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

			fmt.Println("1111111111  5-2", stack) // 🔴 调试代码，应删除
		}
	}
	// ✅ 括号处理逻辑正确，但代码太复杂
	// 🔴 第 126 行条件逻辑容易出错，建议简化
	// 🔴 大量调试代码应删除

	fmt.Println("result", result) // 🔴 调试代码，应删除

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

// 【V0 点评 - 中缀转后缀法】
// ✅ 多位数解析正确
// ✅ 括号处理逻辑正确
// ✅ 中缀转后缀 + 后缀求值的思路正确
// 🔴 大量调试代码应删除
// 🔴 第 128 行条件逻辑有问题：运算符优先级判断不完整
// 🟡 代码太复杂，中缀转后缀容易出错
// 🟡 第 78-80 行可以简化：循环结束后统一处理 curNum
// 🟡 第 121 行条件太复杂，可以简化

// calculateV2 优化版：中缀转后缀法（修复所有问题）
func calculateV2(s string) int {
	// 提取 tokens, 处理多位数字
	var tokens []string
	var curNum string
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if c >= '0' && c <= '9' {
			curNum += string(c)
		} else {
			if curNum != "" {
				tokens = append(tokens, curNum)
				curNum = ""
			}
			tokens = append(tokens, string(c))
		}
	}
	if curNum != "" {
		tokens = append(tokens, curNum)
	}

	// 优先级表
	priority := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2,
		"(": 0, // 左括号优先级最低
	}

	// 中缀转后缀
	var result []string
	var stack []string
	for _, t := range tokens {
		switch t {
		case "(":
			stack = append(stack, t)
		case ")":
			// 弹出直到 (
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 弹出 (
			}
		case "+", "-", "*", "/":
			// 弹出优先级 >= 当前运算符的栈顶
			for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[t] {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, t)
		default:
			// 数字直接入结果
			result = append(result, t)
		}
	}
	// 弹出栈中剩余运算符
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	// 后缀求值
	var nums []int
	for _, r := range result {
		switch r {
		case "+", "-", "*", "/":
			a := nums[len(nums)-2]
			b := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			switch r {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		default:
			num, _ := strconv.Atoi(r)
			nums = append(nums, num)
		}
	}

	return nums[0]
}

// 【V2 点评 - 优化版（中缀转后缀法）】
// ✅ 删除了所有调试代码
// ✅ 修复了多位数解析：简化了逻辑，循环结束后统一处理 curNum
// ✅ 修复了括号处理：用 switch-case 简化了条件判断
// ✅ 修复了优先级判断：用 map 和 switch-case，逻辑清晰
// ✅ 用 switch-case 代替多个 if，代码更简洁
// ✅ 时间 O(n)，空间 O(n)
// 中缀转后缀法的标准写法

// calculateV1 标准答案：递归法（更简洁）
func calculateV1(s string) int {
	i := 0
	return parseExpr(s, &i)
}

// parseExpr 解析表达式（递归处理括号）
func parseExpr(s string, i *int) int {
	nums := []int{} // 数字栈
	num := 0        // 当前数字
	sign := byte('+') // 前一个运算符

	for *i < len(s) {
		c := s[*i]

		if c == ' ' {
			*i++
			continue
		}

		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '(' {
			*i++ // 跳过 '('
			num = parseExpr(s, i) // 递归计算括号内的表达式
		} else {
			// 遇到运算符或 ')'，处理前一个运算符
			applyOp(&nums, &num, &sign)

			if c == ')' {
				break // 遇到 ')' 返回
			}
			sign = c
			num = 0
		}
		*i++
	}

	// 处理最后一个数字
	applyOp(&nums, &num, &sign)

	// 栈中所有数字之和
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

// applyOp 根据运算符处理数字
func applyOp(nums *[]int, num *int, sign *byte) {
	switch *sign {
	case '+':
		*nums = append(*nums, *num)
	case '-':
		*nums = append(*nums, -*num)
	case '*':
		(*nums)[len(*nums)-1] *= *num
	case '/':
		(*nums)[len(*nums)-1] /= *num
	}
}

// 【V1 点评 - 递归法（标准答案）】
// ✅ 代码简洁，用递归处理括号
// ✅ 不需要中缀转后缀，直接计算
// ✅ 时间 O(n)，空间 O(n)（递归栈）
// ✅ 核心思想：遇到 '(' 递归计算括号内的表达式
// ✅ 用栈处理优先级：遇到 * / 立即计算，遇到 + - 入栈最后求和
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【标准答案对比】
//
// | 方法 | 思路 | 优点 | 缺点 |
// |------|------|------|------|
// | 中缀转后缀（V0） | 先转后缀表达式，再求值 | 通用，可处理任意运算符 | 代码复杂，容易出错 |
// | 中缀转后缀优化（V2） | 修复 V0 的问题，用 switch-case 简化 | 代码简洁，逻辑清晰 | 仍需两次遍历 |
// | 递归法（V1） | 递归处理括号，栈处理优先级 | 代码简洁，一次遍历 | 递归有栈开销 |
// | 双栈法 | 数字栈 + 运算符栈 | 通用，可处理任意运算符 | 代码复杂 |
//
// 面试推荐：
//   - 如果只用 + - 和括号：递归法（V1）最简洁
//   - 如果需要 + - * / 和括号：中缀转后缀优化版（V2）或递归法（V1）
// ──────────────────────────────────────────────────────────

func main() {
	fmt.Println(calculate("1+1"))                 // 2
	fmt.Println(calculate("6-4/2"))               // 4
	fmt.Println(calculate("2*(5+5*2)/3+(6/2+8)")) // 21
}
