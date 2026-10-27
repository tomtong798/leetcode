//go:build p0041

// LeetCode 第647题 - 回文子串 (Palindromic Substrings)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】字符串/动态规划
//
// 【题目描述】
// 给你一个字符串 s，请你统计并返回其中回文子串的数目（子串连续，内容相同下标不同算不同）。
//
// 【示例 1】
// 输入：s = "abc"
// 输出：3（"a"、"b"、"c"）
//
// 【示例 2】
// 输入：s = "aaa"
// 输出：6
// 【数据范围】
// 1 <= s.length <= 1000；s 由小写英文字母组成
//
// 【思路提示】中心扩展法：枚举 2n-1 个中心（单字符与字符间隙），向两侧扩展计数。
//
// 题目链接：https://leetcode.cn/problems/palindromic-substrings/
package main

// countSubstrings 返回字符串中回文子串的数目。
func countSubstrings(s string) int {
	// TODO: 实现
	return 0
}

func main() {
	println(countSubstrings("abc")) // 3
	println(countSubstrings("aaa")) // 6
}
