//go:build p0041

// LeetCode 第191题 - 位 1 的个数 (Number of 1 Bits)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】位运算/分治
//
// 【题目描述】
// 给定一个 32 位无符号整数（二进制串），返回其中等于 1 的位的个数（汉明重量）。
//
// 【示例 1】
// 输入：n = 00000000000000000000000000001011（11）
// 输出：3
// 【数据范围】
// 输入是 32 位二进制表示的整数；0 <= n < 2^32
//
// 【思路提示】n &= n-1 每次消去最低位的 1，循环计数；或逐位右移检查最低位。
//
// 题目链接：https://leetcode.cn/problems/number-of-1-bits/
package main

// hammingWeight 返回 32 位无符号整数中 1 的个数。
func hammingWeight(n uint32) int {
	// TODO: 实现
	return 0
}

func main() {
	println(hammingWeight(11)) // 3
}
