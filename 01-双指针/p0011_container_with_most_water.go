//go:build p0041

// LeetCode 第11题 - 盛最多水的容器 (Container With Most Water)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、双指针、贪心
//
// 【题目描述】
// 给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。注意你不能倾斜容器。
//
// 【示例 1】
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：选择第 2 条线（高8）和第 9 条线（高7），宽度 min(8,7)=7 * 宽度8 = 49。
//
// 【示例 2】
// 输入：height = [1,1]
// 输出：1
//
// 【数据范围】
// 2 <= height.length <= 10^5
// 0 <= height[i] <= 10^4
//
// 【思路提示】左右指针从两端开始，每次移动较矮的那一侧（移动高的一侧不可能更优），
// 每步用 area = min(h[l], h[r]) * (r-l) 更新答案。
//
// 题目链接：https://leetcode.cn/problems/container-with-most-water/
package main

// maxArea 返回容器可容纳的最大水量。
func maxArea(height []int) int {
	// TODO: 实现
	return 0
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	maxArea([]int{1, 1})
}
