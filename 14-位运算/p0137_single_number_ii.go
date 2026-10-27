//go:build p0041

// LeetCode 第137题 - 只出现一次的数字 II (Single Number II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】位运算/数组
//
// 【题目描述】
// 给定一个整数数组，某个元素只出现一次，其余每个元素均出现三次，找出那个只出现一次的元素。要求线性时间、常数空间。
//
// 【示例 1】
// 输入：nums = [2,2,3,2]
// 输出：3
// 【数据范围】
// 1 <= nums.length <= 3*10^4；-2^31 <= nums[i] <= 2^31 - 1；nums 中除一个元素外其余都出现三次
//
// 【思路提示】逐位统计 1 的个数再 %3 还原；或用 ones/twos 状态机（ones ^= n; ones &= ^twos ...）。
//
// 题目链接：https://leetcode.cn/problems/single-number-ii/
package main

// singleNumber137 返回数组中只出现一次的数字。
func singleNumber137(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(singleNumber137([]int{2, 2, 3, 2})) // 3
}
