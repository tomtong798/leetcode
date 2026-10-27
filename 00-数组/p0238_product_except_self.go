//go:build p0041

// LeetCode 第238题 - 除自身以外数组的乘积 (Product of Array Except Self)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、前缀和
//
// 【题目描述】
// 给你一个整数数组 nums，返回数组 answer，其中 answer[i] 等于 nums 中除 nums[i] 之外
// 其余各元素的乘积。题目数据保证数组之中任意元素的全部前缀元素和后缀的乘积都在 32 位整数范围内。
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 【示例 1】
// 输入：nums = [1,2,3,4]
// 输出：[24,12,8,6]
//
// 【示例 2】
// 输入：nums = [-1,1,0,-3,3]
// 输出：[0,0,9,0,0]
//
// 【数据范围】
// 2 <= nums.length <= 10^5
// -30 <= nums[i] <= 30
// 保证数组中任意元素的全部前缀元素和后缀的乘积都在 32 位整数范围内
//
// 【思路提示】answer[i] = (i 左边所有数的乘积) * (i 右边所有数的乘积)，
// 先从左到右填前缀乘积，再从右到左乘上后缀乘积，可用 answer 自身节省空间。
//
// 题目链接：https://leetcode.cn/problems/product-of-array-except-self/
package main

// productExceptSelf 返回除自身以外元素的乘积数组（不使用除法）。
func productExceptSelf(nums []int) []int {
	// TODO: 实现
	return nil
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
	productExceptSelf([]int{-1, 1, 0, -3, 3})
}
