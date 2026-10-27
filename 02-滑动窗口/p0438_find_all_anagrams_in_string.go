//go:build p0041

// LeetCode 第438题 - 找到字符串中所有字母异位词 (Find All Anagrams in a String)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】哈希表、字符串、滑动窗口
//
// 【题目描述】
// 给定两个字符串 s 和 p，找到 s 中所有是 p 的异位词的子串，返回这些子串的起始索引。
//
// 【示例 1】
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释：起始索引等于 0 的子串是 "cba"，是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac"，是 "abc" 的异位词。
//
// 【示例 2】
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
//
// 【数据范围】
// 1 <= s.length, p.length <= 3 * 10^4
// s 和 p 仅包含小写英文字母
//
// 【思路提示】定长滑动窗口 + 26 位计数数组，
// 维护窗口与 p 的差异字符数 diff，diff == 0 时记录起始下标。
//
// 题目链接：https://leetcode.cn/problems/find-all-anagrams-in-a-string/
package main

// findAnagrams 返回所有异位词子串的起始索引。
func findAnagrams(s string, p string) []int {
	// TODO: 实现
	return nil
}

func main() {
	findAnagrams("cbaebabacd", "abc")
	findAnagrams("abab", "ab")
}
