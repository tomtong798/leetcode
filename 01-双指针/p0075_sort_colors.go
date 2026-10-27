//go:build p0041

// LeetCode 第75题 - 颜色分类 (Sort Colors / 荷兰国旗)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、双指针、排序
//
// 【题目描述】
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums，
// 原地对它们进行排序，使得相同颜色相邻，并按红色、白色、蓝色顺序排列。
// 使用整数 0、1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决。
// 进阶：你能想出一个仅使用常数空间的一趟扫描算法吗？
//
// 【示例 1】
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
//
// 【示例 2】
// 输入：nums = [2,0,1]
// 输出：[0,1,2]
//
// 【数据范围】
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
// 【思路提示】三指针（Dutch National Flag）：p0 表 0 的右边界，p2 表 2 的左边界，
// cur 扫描；遇 0 换到前段，遇 2 换到后段（注意 cur 不前进需再判断），遇 1 前进。
//
// 题目链接：https://leetcode.cn/problems/sort-colors/
package main

// sortColors 原地排序 0/1/2 数组。
func sortColors(nums []int) {
	// TODO: 实现
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColors([]int{2, 0, 1})
}
