//go:build p0041

// LeetCode 第567题 - 字符串的排列 (Permutation in String)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】滑动窗口、哈希表、字符串
//
// 【题目描述】
// 给你两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true；
// 否则，返回 false。换句话说，s1 的排列之一是 s2 的子串。
//
// 【示例 1】
// 输入：s1 = "ab", s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba")。
//
// 【示例 2】
// 输入：s1 = "ab", s2 = "eidboaoo"
// 输出：false
//
// 【数据范围】
// 1 <= s1.length, s2.length <= 10^4
// s1 和 s2 仅包含小写字母
//
// 【思路提示】固定窗口大小 = len(s1) 的滑动窗口，
// 用两个 26 长度计数数组（或一个差分数组）比较窗口内字符频次是否一致。
//
// 题目链接：https://leetcode.cn/problems/permutation-in-string/
package main

// checkInclusion 判断 s2 是否包含 s1 的排列。
func checkInclusion(s1 string, s2 string) bool {
	// TODO: 实现
	return false
}

func main() {
	checkInclusion("ab", "eidbaooo")
	checkInclusion("ab", "eidboaoo")
}
