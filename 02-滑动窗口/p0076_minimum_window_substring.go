//go:build p0041

// LeetCode 第76题 - 最小覆盖子串 (Minimum Window Substring)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】滑动窗口、哈希表、字符串
//
// 【题目描述】
// 给你一个字符串 s、一个字符串 t。返回 s 中涵盖 t 所有字符（含重复字符）的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""。
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 【示例 1】
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 【示例 2】
// 输入：s = "a", t = "a"
// 输出："a"
//
// 【示例 3】
// 输入：s = "a", t = "aa"
// 输出：""
// 解释：t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子串。
//
// 【数据范围】
// 1 <= s.length, t.length <= 10^5
// s 和 t 由英文字母组成
//
// 【思路提示】滑动窗口经典题：need 记录 t 的字符需求，window 记录当前窗口计数，
// 右扩时满足一个需求就 matched++；全部满足时收缩左边并更新最短答案。
//
// 题目链接：https://leetcode.cn/problems/minimum-window-substring/
package main

// minWindow 返回 s 中覆盖 t 所有字符的最小子串。
func minWindow(s string, t string) string {
	// TODO: 实现
	return ""
}

func main() {
	minWindow("ADOBECODEBANC", "ABC")
	minWindow("a", "a")
	minWindow("a", "aa")
}
