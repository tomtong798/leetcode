//go:build p0041

// LeetCode 第210题 - 课程表 II (Course Schedule II)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】深度优先搜索、广度优先搜索、图、拓扑排序
//
// 【题目描述】
// 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
// 给你一个数组 prerequisites，其中 prerequisites[i] = [ai, bi] 表示在学习 ai 前必须先学 bi。
// 请返回你为了学完所有课程所安排的学习顺序。如果有多个顺序，返回任意一个即可。
// 如果不可能完成所有课程，返回空数组。
//
// 【示例 1】
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：[0,1]
//
// 【示例 2】
// 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// 输出：[0,2,1,3] 或 [0,1,2,3]
//
// 【示例 3】
// 输入：numCourses = 1, prerequisites = []
// 输出：[0]
//
// 【数据范围】
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= numCourses * (numCourses - 1)
//
// 【思路提示】Kahn 拓扑排序：统计入度，入度为 0 的入队，
// 逐个出队并给后继减入度；出队顺序即修课顺序。出队数 != numCourses 则有环返回空。
//
// 题目链接：https://leetcode.cn/problems/course-schedule-ii/
package main

// findOrder 返回拓扑排序结果。
func findOrder(numCourses int, prerequisites [][]int) []int {
	// TODO: 实现
	return nil
}

func main() {
	findOrder(2, [][]int{{1, 0}})
	findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	findOrder(1, nil)
}
