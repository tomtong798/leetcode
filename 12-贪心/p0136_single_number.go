//go:build p0041

// LeetCode 第136题 - 只出现一次的数字 (Single Number)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】位运算/数组
//
// 【题目描述】
// 给定一个非空整数数组，某个元素只出现一次，其余每个元素均出现两次，找出那个只出现一次的元素。要求线性时间、常数空间。
//
// 【示例 1】
// 输入：nums = [4,1,2,1,2]
// 输出：4
// 【数据范围】
// 1 <= nums.length <= 3*10^4；-3*10^4 <= nums[i] <= 3*10^4
//
// 【思路提示】全体元素异或：成对的数相互抵消为 0，剩下唯一出现一次的那个数。
//
// 题目链接：https://leetcode.cn/problems/single-number/
package main

// singleNumber 返回数组中只出现一次的数字。
func singleNumber(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(singleNumber([]int{4, 1, 2, 1, 2})) // 4
}
