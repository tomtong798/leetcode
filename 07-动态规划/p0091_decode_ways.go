//go:build p0041

// LeetCode 第91题 - 解码方法 (Decode Ways)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】字符串/动态规划
//
// 【题目描述】
// 一条包含字母 A-Z 的消息通过如下映射编码为数字：'1'->'A'，'26'->'Z'。给定一个只含数字的字符串，计算解码方法总数。
//
// 【示例 1】
// 输入：s = "12"
// 输出：2（"AB" 或 "L"）
//
// 【示例 2】
// 输入：s = "226"
// 输出：3
//
// 【示例 3】
// 输入：s = "06"
// 输出：0
// 【数据范围】
// 1 <= s.length <= 100；s 只包含数字且可能含前导零
//
// 【思路提示】dp[i] = (s[i-1]!='0' ? dp[i-1] : 0) + (10<=两位数<=26 ? dp[i-2] : 0)，前导零无解。
//
// 题目链接：https://leetcode.cn/problems/decode-ways/
package main

// numDecodings 返回数字字符串的解码方法总数。
func numDecodings(s string) int {
	// TODO: 实现
	return 0
}

func main() {
	println(numDecodings("12")) // 2
	println(numDecodings("226")) // 3
	println(numDecodings("06")) // 0
}
