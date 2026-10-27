//go:build p0041

// LeetCode 第125题 - 验证回文串 (Valid Palindrome)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】双指针、字符串
//
// 【题目描述】
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
// 短语正着读和反着读都一样，则可以认为该短语是一个回文串。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是回文串，返回 true；否则返回 false。
//
// 【示例 1】
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
//
// 【示例 2】
// 输入：s = "race a car"
// 输出：false
//
// 【示例 3】
// 输入：s = " "
// 输出：true
//
// 【数据范围】
// 1 <= s.length <= 2 * 10^5
// s 仅由可打印 ASCII 字符组成
//
// 【思路提示】双指针从两端向中间收拢，跳过非字母数字字符，比较时统一转小写。
//
// 题目链接：https://leetcode.cn/problems/valid-palindrome/
package main

// isPalindrome 判断字符串是否为回文串（忽略大小写和非字母数字）。
func isPalindrome(s string) bool {
	// TODO: 实现
	return false
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
	isPalindrome("race a car")
	isPalindrome(" ")
}
