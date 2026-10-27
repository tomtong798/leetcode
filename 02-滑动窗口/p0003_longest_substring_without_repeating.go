//go:build p0041

// LeetCode 第3题 - 无重复字符的最长子串 (Longest Substring Without Repeating Characters)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】滑动窗口、哈希表、字符串
//
// 【题目描述】
// 给定一个字符串 s，请你找出其中不含有重复字符的最长子串的长度。
//
// 【示例 1】
// 输入：s = "abcabcbb"
// 输出：3
// 解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 【示例 2】
// 输入：s = "bbbbb"
// 输出：1
// 解释：无重复字符的最长子串是 "b"，长度为 1。
//
// 【示例 3】
// 输入：s = "pwwkew"
// 输出：3
// 解释：无重复字符的最长子串是 "wke"，注意答案必须是子串的长度，"pwke" 是子序列而非子串。
//
// 【数据范围】
// 0 <= s.length <= 5 * 10^4
// s 由英文字母、数字、符号和空格组成
//
// 【思路提示】滑动窗口 + 哈希表记录字符最近出现位置；
// 右指针扩张时若字符重复，把左指针跳到"重复位置+1"，全程更新最大长度。
//
// 题目链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package main

// lengthOfLongestSubstring 返回无重复字符最长子串的长度。
func lengthOfLongestSubstring(s string) int {
	// TODO: 实现
	return 0
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("bbbbb")
	lengthOfLongestSubstring("pwwkew")
}
