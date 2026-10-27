//go:build p0041

// LeetCode 第394题 - 字符串解码 (Decode String)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、字符串、递归
//
// 【题目描述】
// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为：k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k，
// 例如不会出现像 3a 或 2[4] 的输入。
//
// 【示例 1】
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
//
// 【示例 2】
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
//
// 【示例 3】
// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
//
// 【数据范围】
// 1 <= s.length <= 30
// s 由小写英文字母、数字和方括号 '[]' 组成
// 保证输入是合法的，k 的范围为 [1, 300]
//
// 【思路提示】双栈：数字栈存重复次数，字符串栈存"当前已解码的字符串"；
// 遇 '[' 把当前字符串与数字压栈并清空；遇 ']' 弹出数字 n 与 prev，
// 当前字符串 = prev + 当前字符串重复 n 次。也可用递归。
//
// 题目链接：https://leetcode.cn/problems/decode-string/
package main

// decodeString 解码字符串。
func decodeString(s string) string {
	// TODO: 实现
	return ""
}

func main() {
	decodeString("3[a]2[bc]")
	decodeString("3[a2[c]]")
	decodeString("2[abc]3[cd]ef")
}
