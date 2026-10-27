//go:build p0041

// LeetCode 第239题 - 滑动窗口最大值 (Sliding Window Maximum)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】队列、数组、滑动窗口、单调队列、堆
//
// 【题目描述】
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到最右侧。
// 你只能看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
//
// 【示例 1】
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
//
// 【示例 2】
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 【数据范围】
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length
//
// 【思路提示】单调递减双端队列存下标：
// 队头即窗口最大值的下标；新元素入队前从队尾弹出所有更小的；
// 队头下标超出窗口左边界时弹出。
//
// 题目链接：https://leetcode.cn/problems/sliding-window-maximum/
package main

// maxSlidingWindow 返回每个滑动窗口的最大值。
func maxSlidingWindow(nums []int, k int) []int {
	// TODO: 实现
	return nil
}

func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	maxSlidingWindow([]int{1}, 1)
}
