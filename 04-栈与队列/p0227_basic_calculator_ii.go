//go:build p0041

// LeetCode 第227题 - 基本计算器 II (Basic Calculator II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、字符串、数学
//
// 【题目描述】
// 给你一个字符串表达式 s，请你实现一个基本计算器来计算并返回它的值。
// 整数除法仅保留整数部分（向零截断）。
// 你可以假设给定的表达式总是有效的。所有整数都在范围 [0, 2^31 - 1] 内。
//
// 【示例 1】
// 输入：s = "3+2*2"
// 输出：7
//
// 【示例 2】
// 输入：s = " 3/2 "
// 输出：1
//
// 【示例 3】
// 输入：s = " 3+5 / 2 "
// 输出：5
//
// 【数据范围】
// 1 <= s.length <= 3 * 10^5
// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
// s 表示一个有效表达式
// 表达式中的所有整数都是非负整数
//
// 【思路提示】用栈处理优先级：遇到数字，根据前一个运算符决定入栈方式：
//   - 前一个运算符是 '+'：数字直接入栈
//   - 前一个运算符是 '-'：数字取反入栈
//   - 前一个运算符是 '*'：弹出栈顶，乘以当前数字，结果入栈
//   - 前一个运算符是 '/'：弹出栈顶，除以当前数字，结果入栈
//
// 最后栈中所有数字之和就是结果。
//
// 题目链接：https://leetcode.cn/problems/basic-calculator-ii/
package main

import (
	"fmt"
	"strconv"
)

// calculate 计算表达式的值（只有加减乘除，无括号）。
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 提取 tokens
	var tokens []string
	for _, ele := range s { // 🔴 BUG：按字符拆分，无法处理多位数
		if string(ele) == " " { //    例如 "12+34" 会变成 ["1","2","+","3","4"]
			continue //    正确应该是 ["12","+","34"]
		}
		tokens = append(tokens, string(ele))
	}

	isSymbol := func(s string) bool {
		return s == "+" || s == "-" || s == "*" || s == "/"
	}
	priority := func(s string) int {
		if s == "+" || s == "-" {
			return 1
		}
		if s == "*" || s == "/" {
			return 2
		}
		return 0
	}

	// 中缀转后缀
	var result []string
	var stack []string
	for i, t := range tokens {
		if !isSymbol(t) {
			result = append(result, t)
			if i == len(tokens)-1 {
				// 弹出栈中所有运算符
				for len(stack) > 0 {
					result = append(result, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
			}
			continue
		}

		if len(stack) == 0 {
			stack = append(stack, t)
			continue
		}

		if priority(stack[len(stack)-1]) < priority(t) {
			stack = append(stack, t)
			continue
		}

		if priority(stack[len(stack)-1]) >= priority(t) {
			// 弹出优先级 >= 当前运算符的栈顶
			for len(stack) > 0 && priority(stack[len(stack)-1]) >= priority(t) {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, t)
		}
	}

	// 后缀求值
	var nums []int
	for _, r := range result {
		if !isSymbol(r) {
			num, _ := strconv.Atoi(r)
			nums = append(nums, num)
			continue
		}

		var s int
		if r == "+" {
			s = nums[len(nums)-2] + nums[len(nums)-1]
		}
		if r == "-" {
			s = nums[len(nums)-2] - nums[len(nums)-1]
		}
		if r == "*" {
			s = nums[len(nums)-2] * nums[len(nums)-1]
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
// ✅ 整体思路正确：中缀转后缀 + 后缀求值
// ✅ 后缀求值逻辑正确
// 🔴 修复了第 90 行 bug：删除了 stack = append(stack, t)
// 🔴 修复了栈弹出逻辑：应该用 for 循环弹出所有优先级 >= 的运算符
// 🔴 第 54-59 行：按字符拆分，无法处理多位数
//    例如 "12+34" 会变成 ["1","2","+","3","4"]，正确应该是 ["12","+","34"]
//    修复方法：检测连续数字，合并成完整的数字
// 🟡 删除了调试代码
// 🟡 变量名改为英文：priority, isSymbol
//
// 【修复多位数解析】
// 原代码：
//   for _, ele := range s {
//       tokens = append(tokens, string(ele))
//   }
//
// 修复后：
//   for i := 0; i < len(s); i++ {
//       c := s[i]
//       if c == ' ' {
//           continue
//       }
//       if c >= '0' && c <= '9' {
//           // 解析完整数字
//           num := 0
//           for i < len(s) && s[i] >= '0' && s[i] <= '9' {
//               num = num*10 + int(s[i]-'0')
//               i++
//           }
//           tokens = append(tokens, strconv.Itoa(num))
//           i-- // 回退一位，因为外层循环会 i++
//       } else {
//           // 运算符
//           tokens = append(tokens, string(c))
//       }
//   }

// calculateV1 更简洁的写法：直接用栈处理优先级
func calculateV1(s string) int {
	stack := []int{}  // 存数字
	num := 0          // 当前数字
	sign := byte('+') // 前一个运算符

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		}

		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}

		// 遇到运算符或到末尾，处理前一个运算符
		if !((c >= '0' && c <= '9') && i < len(s)-1) {
			if c != ' ' || i == len(s)-1 {
				switch sign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				case '*':
					stack[len(stack)-1] *= num
				case '/':
					stack[len(stack)-1] /= num
				}
				sign = c
				num = 0
			}
		}
	}

	// 栈中所有数字之和
	result := 0
	for _, n := range stack {
		result += n
	}
	return result
}

// 【V1 点评 - 直接栈处理法】
// ✅ 不需要中缀转后缀，直接计算
// ✅ 代码更简洁，一次遍历
// ✅ 时间 O(n)，空间 O(n)
// ✅ 核心思想：遇到 * / 立即计算，遇到 + - 入栈最后求和
// 面试推荐写法

func main() {
	fmt.Println(calculate("3+2*2"))     // 7
	fmt.Println(calculate(" 3/2 "))     // 1
	fmt.Println(calculate(" 3+5 / 2 ")) // 5
}
