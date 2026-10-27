//go:build p0041

// LeetCode 第70题 - 爬楼梯 (Climbing Stairs)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】动态规划、记忆化搜索、数学
//
// 【题目描述】
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 【示例 1】
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶：1 阶 + 1 阶；直接 2 阶。
//
// 【示例 2】
// 输入：n = 3
// 输出：3
// 解释：有三种方法：1+1+1、1+2、2+1。
//
// 【数据范围】
// 1 <= n <= 45
//
// 【思路提示】f(n) = f(n-1) + f(n-2)，本质是斐波那契数列，两个变量滚动即可，O(1) 空间。
//
// 题目链接：https://leetcode.cn/problems/climbing-stairs/
package main

// climbStairs 返回爬 n 阶楼梯的方法数。
func climbStairs(n int) int {
	// TODO: 实现
	return 0
}

func main() {
	climbStairs(2)
	climbStairs(3)
	climbStairs(45)
}
