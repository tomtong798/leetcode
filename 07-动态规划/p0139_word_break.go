//go:build p0041

// LeetCode 第139题 - 单词拆分 (Word Break)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】字典树、记忆化搜索、动态规划
//
// 【题目描述】
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 【示例 1】
// 输入：s = "leetcode", wordDict = ["leet", "code"]
// 输出：true
// 解释：返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
// 【示例 2】
// 输入：s = "applepenapple", wordDict = ["apple", "pen"]
// 输出：true
// 解释：返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//
// 【示例 3】
// 输入：s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出：false
//
// 【数据范围】
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和所有 wordDict[i] 都由小写英文字母组成
//
// 【思路提示】dp[i] 表示 s 前 i 个字符能否被拼出：dp[i] = 存在 j < i 使 dp[j] 为真且 s[j:i] 在字典中。
// 用哈希集合存字典；内层只需尝试字典中最长单词长度以内的切分点。
//
// 题目链接：https://leetcode.cn/problems/word-break/
package main

// wordBreak 判断 s 能否由字典单词拼接而成。
func wordBreak(s string, wordDict []string) bool {
	// TODO: 实现
	return false
}

func main() {
	wordBreak("leetcode", []string{"leet", "code"})
	wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
}
