//go:build p0041

// LeetCode 第231题 - 2 的幂 (Power of Two)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】位运算/数学
//
// 【题目描述】
// 给定一个整数 n，判断它是否是 2 的幂次方。
//
// 【示例 1】
// 输入：n = 1
// 输出：true
//
// 【示例 2】
// 输入：n = 16
// 输出：true
//
// 【示例 3】
// 输入：n = 3
// 输出：false
// 【数据范围】
// -2^31 <= n <= 2^31 - 1
//
// 【思路提示】2 的幂二进制中只有一个 1：判断 n > 0 && n&(n-1) == 0。
//
// 题目链接：https://leetcode.cn/problems/power-of-two/
package main

// isPowerOfTwo 判断整数 n 是否为 2 的幂。
func isPowerOfTwo(n int) bool {
	// TODO: 实现
	return false
}

func main() {
	println(isPowerOfTwo(1)) // true
	println(isPowerOfTwo(16)) // true
	println(isPowerOfTwo(3)) // false
}
