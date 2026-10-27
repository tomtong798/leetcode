//go:build p0041

// LeetCode 第153题 - 寻找旋转排序数组中的最小值 (Find Minimum in Rotated Sorted Array)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
// 例如，原数组 nums = [0,1,2,4,5,6,7] 在旋转后可能变为 [4,5,6,7,0,1,2]。
// 给你元素值互不相同的数组 nums，它原来是一个升序数组，并按上述情形进行了多次旋转。
// 请你找出并返回数组中的最小元素。你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 【示例 1】
// 输入：nums = [3,4,5,1,2]
// 输出：1
// 解释：原数组为 [1,2,3,4,5]，旋转 3 次得到输入数组。
//
// 【示例 2】
// 输入：nums = [4,5,6,7,0,1,2]
// 输出：0
//
// 【示例 3】
// 输入：nums = [11,13,15,17]
// 输出：11
//
// 【数据范围】
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// nums 中的所有整数互不相同
//
// 【思路提示】与右端点比较：nums[mid] > nums[right] 说明最小值在 mid 右侧，
// left = mid+1；否则最小值在 mid 及其左侧，right = mid。
//
// 题目链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
package main

// findMin 返回旋转数组中的最小值。
func findMin(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	findMin([]int{3, 4, 5, 1, 2})
	findMin([]int{4, 5, 6, 7, 0, 1, 2})
	findMin([]int{11, 13, 15, 17})
}
