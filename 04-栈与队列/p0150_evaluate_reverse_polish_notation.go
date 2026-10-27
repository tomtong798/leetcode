//go:build p0041

// LeetCode 第150题 - 逆波兰表达式求值 (Evaluate Reverse Polish Notation)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、数组、数学
//
// 【题目描述】
// 给你一个字符串数组 tokens，表示一个逆波兰表达式。求该逆波兰表达式的值。
// 有效的算符为 +、-、*、/。每个操作数都是整数。
// 两个整数之间的除法总是向零截断。
//
// 【示例 1】
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：((2 + 1) * 3) = 9
//
// 【示例 2】
// 输入：tokens = ["4","13","5","/","+"]
// 输出：6
// 解释：4 + (13 / 5) = 4 + 2 = 6
//
// 【数据范围】
// 1 <= tokens.length <= 10^4
// tokens[i] 是算符或范围 [-200, 200] 的整数
//
// 【思路提示】栈：遇数入栈，遇运算符弹两个（注意先弹的是右操作数），结果入栈。
//
// 题目链接：https://leetcode.cn/problems/evaluate-reverse-polish-notation/
package main

import (
	"fmt"
	"strconv"
)

// evalRPN 求逆波兰表达式值。
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	var stack []string // 🟡 用字符串栈效率低，应该用 []int
	var s string

	for _, ele := range tokens {
		if ele != "+" && ele != "-" && ele != "*" && ele != "/" {
			stack = append(stack, ele)
			continue
		}

		a, _ := strconv.Atoi(stack[len(stack)-1]) // 先弹的是右操作数
		b, _ := strconv.Atoi(stack[len(stack)-2]) // 后弹的是左操作数

		if ele == "+" {
			s = strconv.Itoa(a + b)
		} else if ele == "-" {
			s = strconv.Itoa(b - a) // ✅ 减法顺序正确：b - a
		} else if ele == "*" {
			s = strconv.Itoa(a * b)
		} else {
			s = strconv.Itoa(b / a) // ✅ 除法顺序正确：b / a
		}

		stack = stack[:len(stack)-2]
		stack = append(stack, s)
	}

	result, _ := strconv.Atoi(s) // 🔴 BUG：如果 tokens 只有一个数字，s 是空字符串
	return result                //    应该返回 strconv.Atoi(stack[0])
}

// 【V0 点评】
// 编译通过，示例正确，但有 bug：
//   🔴 第 68 行：返回 strconv.Atoi(s)，但如果 tokens 只有一个数字（如 ["5"]），
//      s 是空字符串，Atoi("") 返回 0，应该返回 5
//      修复：return strconv.Atoi(stack[0])
//   🟡 用字符串栈效率低：每次都要 Atoi/Itoa 转换
//      应该用 []int 栈，直接存整数
//   🟡 第 38-40 行：空数组判断可以简化
//   ✅ 运算符处理逻辑正确：先弹右操作数 a，后弹左操作数 b
//   ✅ 减法和除法顺序正确：b - a 和 b / a

func main() {
	fmt.Println(evalRPNV1([]string{"2", "1", "+", "3", "*"}))  // 9
	fmt.Println(evalRPNV1([]string{"4", "13", "5", "/", "+"})) // 6
	fmt.Println(evalRPNV1([]string{"5"}))                      // 5（测试单数字）
}

// evalRPNV1 修复版：用整数栈，时间 O(n)，空间 O(n)
func evalRPNV1(tokens []string) int {
	stack := []int{} // ✅ 用整数栈，避免字符串转换

	for _, token := range tokens {
		switch token {
		case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b+a)
		case "-":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a) // ✅ 减法：b - a
		case "*":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b*a)
		case "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a) // ✅ 除法：b / a
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0] // ✅ 直接返回栈顶，不会出错
}

// 【V1 点评 - 修复版】
// ✅ 修复了返回值：直接返回 stack[0]，不会出错
// ✅ 用整数栈：避免 Atoi/Itoa 转换，效率更高
// ✅ 用 switch 简化判断：代码更清晰
// ✅ 时间 O(n)，空间 O(n)
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【基础知识】
//
// 三种表达式对比（以 3 + 4 * 2 为例）：
//
// 1. 中缀表达式（人类习惯）
//    3 + 4 * 2
//    - 运算符在操作数中间
//    - 需要优先级规则（先乘除后加减）
//    - 需要括号改变优先级：(3 + 4) * 2
//    - 人类易读，但计算机难解析
//
// 2. 前缀表达式（波兰式）
//    + 3 * 4 2
//    - 运算符在操作数之前
//    - 不需要括号和优先级
//    - 从右向左扫描计算
//    - 例如：+ 3 * 4 2 → 先算 * 4 2 = 8，再算 + 3 8 = 11
//
// 3. 后缀表达式（逆波兰式）← 本题
//    3 4 2 * +
//    - 运算符在操作数之后
//    - 不需要括号和优先级
//    - 从左向右扫描，用栈计算
//    - 例如：3 4 2 * + → 先算 4 2 * = 8，再算 3 8 + = 11
//
// 后缀表达式计算规则：
//   - 遇到数字：入栈
//   - 遇到运算符：弹出两个数字计算，结果入栈
//   - 最后栈中只剩一个数字，就是结果
//
// Go 字符串/整数转换函数：
//   strconv.Atoi(s)    // string → int，如 Atoi("123") = 123
//   strconv.Itoa(i)    // int → string，如 Itoa(123) = "123"
//
//   记忆技巧：
//   - Atoi = ASCII to Integer（ASCII 转整数）
//   - Itoa = Integer to ASCII（整数转 ASCII）
//
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法              | 时间   | 空间   | 面试期望        |
// |-------------------|--------|--------|-----------------|
// | 字符串栈（V0）    | O(n)   | O(n)   | ❌ 有 bug        |
// | 整数栈（V1）      | O(n)   | O(n)   | ✅ 推荐写法      |
//
// 【核心易错点】
// 1. 先弹出的是右操作数，后弹出的是左操作数
// 2. 减法和除法顺序：b - a 和 b / a（不是 a - b）
// 3. 用整数栈而不是字符串栈，避免频繁转换
// 4. 最后返回 stack[0]，不是某个中间变量
//
// 【面试讲解建议】
// 1. 先说思路："遇数入栈，遇运算符弹两个数计算，结果入栈"
// 2. 强调操作数顺序：先弹右操作数，后弹左操作数
// 3. 画图说明栈的变化过程
// 4. 分析复杂度：时间 O(n)，空间 O(n)
// ──────────────────────────────────────────────────────────
