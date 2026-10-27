//go:build p0041

// LeetCode 第31题 - 下一个排列 (Next Permutation)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、双指针、数学
//
// 【题目描述】
// 整数数组的一个排列就是将其所有成员进行序列化后的序列。
// 下一个排列是指其整数的下一个字典序更大的排列。
// 必须原地修改，只允许使用额外常数空间。
//
// 【示例 1】
// 输入：nums = [1,2,3]
// 输出：[1,3,2]
//
// 【示例 2】
// 输入：nums = [3,2,1]
// 输出：[1,2,3]
//
// 【示例 3】
// 输入：nums = [1,1,5]
// 输出：[1,5,1]
//
// 【数据范围】
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//
// 【思路提示】三步：从右往左找第一个升序对 i（nums[i] < nums[i+1]）；
// 再从右找第一个大于 nums[i] 的数与之交换；最后反转 i+1 到末尾。
// 若不存在升序对则整体反转。
//
// 题目链接：https://leetcode.cn/problems/next-permutation/
package main

// nextPermutation 原地生成下一个字典序排列。
func nextPermutation(nums []int) {
	// TODO: 实现
}

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{1, 1, 5})
}
