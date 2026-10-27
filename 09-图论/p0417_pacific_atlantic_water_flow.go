//go:build p0041

// LeetCode 第417题 - 太平洋大西洋水流问题 (Pacific Atlantic Water Flow)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】深度优先搜索、广度优先搜索、数组、矩阵
//
// 【题目描述】
// 有一个 m × n 的矩形岛屿，与太平洋和大西洋相邻。
// 太平洋处于大陆的左边界和上边界，而大西洋处于大陆的右边界和下边界。
// 这个岛被分割成一个由若干方格组成的网格。给定一个 m x n 的整数矩阵 heights，
// heights[r][c] 表示坐标 (r, c) 上单元格高于海平面的高度。
// 雨水可以从相邻格子流向高度相等或更低的格子。返回雨水流动即可到达太平洋和大西洋的坐标列表。
//
// 【示例 1】
// 输入：heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// 输出：[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
// 【数据范围】
// m == heights.length, n == heights[r].length
// 1 <= m, n <= 200
// 0 <= heights[r][c] <= 10^5
//
// 【思路提示】从两个大洋边界反向 DFS 上爬（邻格高度 >= 当前），
// 两个可达集合求交即为答案。
//
// 题目链接：https://leetcode.cn/problems/pacific-atlantic-water-flow/
package main

// pacificAtlantic 返回可同时流向两洋的坐标。
func pacificAtlantic(heights [][]int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}})
}
