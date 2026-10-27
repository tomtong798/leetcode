//go:build p0041

// LeetCode 第154题 - 寻找旋转排序数组中的最小值 II (Find Minimum in Rotated Sorted Array II)
//
// 【难度】困难 | 【面试频率】★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
// 注意，数组 [a[0], a[1], ..., a[n-1]] 旋转一次的结果为数组 [a[n-1], a[0], ..., a[n-2]]。
// 给你一个可能存在重复元素值的数组 nums，原来是一个升序排列的数组，并进行了上述旋转。
// 请返回数组中的最小元素。你必须尽可能减少整个过程的操作步骤。
//
// 【示例 1】
// 输入：nums = [1,3,5]
// 输出：1
//
// 【示例 2】
// 输入：nums = [2,2,2,0,1]
// 输出：0
//
// 【数据范围】
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
// 【思路提示】与第 153 题类似，但含重复元素：nums[mid] == nums[right] 时无法判断方向，
// 只能 right-- 逐步缩小（最坏 O(n)）。
//
// 题目链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
package main

// findMin154 返回旋转数组（含重复）中的最小值。
func findMin154(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	findMin154([]int{1, 3, 5})
	findMin154([]int{2, 2, 2, 0, 1})
}
