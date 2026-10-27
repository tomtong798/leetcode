//go:build p0041

// LeetCode 第207题 - 课程表 (Course Schedule)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】深度优先搜索、广度优先搜索、图、拓扑排序
//
// 【题目描述】
// 这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1。
// 在选修某些课程之前需要一些先修课程。先修课程按数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi]，表示如果要学习课程 ai 则必须先学习课程 bi。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true；否则，返回 false。
//
// 【示例 1】
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，需要完成课程 0。这是可能的。
//
// 【示例 2】
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前需要完成课程 0，学习课程 0 之前需要完成课程 1，这是不可能的。
//
// 【数据范围】
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
//
// 【思路提示】本质是判断有向图是否有环。拓扑排序（Kahn 算法）：
// 统计入度，入度为 0 的节点入队，逐个出队并给后继减入度；
// 最后出队节点数 == 课程数则无环。也可用 DFS 三色标记法。
//
// 题目链接：https://leetcode.cn/problems/course-schedule/
package main

// canFinish 判断课程安排是否可行（无环）。
func canFinish(numCourses int, prerequisites [][]int) bool {
	// TODO: 实现
	return false
}

func main() {
	canFinish(2, [][]int{{1, 0}})
	canFinish(2, [][]int{{1, 0}, {0, 1}})
}
