//go:build p0041

// LeetCode 第49题 - 字母异位词分组 (Group Anagrams)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、哈希表、字符串、排序
//
// 【题目描述】
// 给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。
// 字母异位词是由重新排列源单词的所有字母得到的一个新单词。
//
// 【示例 1】
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["eat","tea","ate"],["tan","nat"]]
//
// 【示例 2】
// 输入: strs = [""]
// 输出: [[""]]
//
// 【数据范围】
// 1 <= strs.length <= 10^4
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
// 【思路提示】每个单词排序后作为哈希 key 分组；
// 进阶用 26 位计数数组编码作 key（避免排序 O(L)）。
//
// 题目链接：https://leetcode.cn/problems/group-anagrams/
package main

// groupAnagrams 将字母异位词分组。
func groupAnagrams(strs []string) [][]string {
	// TODO: 实现
	return nil
}

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	groupAnagrams([]string{""})
	groupAnagrams([]string{"a"})
}
