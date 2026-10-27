//go:build p0041

// LeetCode 第88题 - 合并两个有序数组 (Merge Sorted Array)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】数组、双指针、排序
//
// 【题目描述】
// 给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n，
// 分别表示 nums1 和 nums2 中的元素数目。请你合并 nums2 到 nums1 中，使合并后的数组同样按非递减顺序排列。
// 注意：最终合并后数组不应由函数返回，而是存储在数组 nums1 中。
// nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0，应忽略。
//
// 【示例 1】
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
//
// 【示例 2】
// 输入：nums1 = [1], m = 1, nums2 = [], n = 0
// 输出：[1]
//
// 【数据范围】
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
//
// 【思路提示】从后往前三指针：p1=m-1, p2=n-1, 写指针 p=m+n-1，
// 比较 nums1[p1] 与 nums2[p2]，大的放末尾；避免覆盖 nums1 前段。
//
// 题目链接：https://leetcode.cn/problems/merge-sorted-array/
package main

// mergeSortedArray 合并 nums2 到 nums1（原地）。
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// TODO: 实现
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	mergeSortedArray(nums1, 3, []int{2, 5, 6}, 3)
}
