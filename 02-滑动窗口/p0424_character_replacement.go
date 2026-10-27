//go:build p0041

// LeetCode 第424题 - 替换后的最长重复字符 (Longest Repeating Character Replacement)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】滑动窗口、哈希表、字符串
//
// 【题目描述】
// 给你一个字符串 s 和一个整数 k。你可以选择字符串中的任一字符，
// 并将其更改为任何其他大写英文字符。该过程最多执行 k 次。
// 请返回执行上述操作后，得到的仅包含同一字符的最长子串的长度。
//
// 【示例 1】
// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两次操作将两个 'A' 替换为 'B'（或相反），得到 "BBBB" 或 "AAAA"。
//
// 【示例 2】
// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：将中间的 'A' 替换为 'B'，字符串变为 "AABBBBA"，子串 "BBBB" 长度为 4。
//
// 【数据范围】
// 1 <= s.length <= 10^5
// s 仅由大写英文字母组成
// 0 <= k <= s.length
//
// 【思路提示】滑动窗口：维护窗口内出现最多的字符次数 maxCount；
// 当窗口长度 - maxCount > k 时收缩左边界。maxCount 不需要随左移而减小。
//
// 题目链接：https://leetcode.cn/problems/longest-repeating-character-replacement/
package main

// characterReplacement 返回替换后最长同字符子串长度。
func characterReplacement(s string, k int) int {
	// TODO: 实现
	return 0
}

func main() {
	characterReplacement("ABAB", 2)
	characterReplacement("AABABBA", 1)
}
