//go:build p0041

// LeetCode 第63题 - 不同路径 II (Unique Paths II)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】数组/动态规划/矩阵
//
// 【题目描述】
// 一个机器人位于 m x n 网格的左上角，网格中有障碍物（obstacleGrid 值为 1 的格子），每次只能向下或向右移动一步，问到达右下角有多少条不同路径。
//
// 【示例 1】
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 【数据范围】
// 1 <= obstacleGrid.length, obstacleGrid[0].length <= 100；obstacleGrid[i][j] 为 0 或 1；起终点无障碍
//
// 【思路提示】动态规划同 62 题，遇到障碍位直接置 dp=0，注意首行首列的初始化传播。
//
// 题目链接：https://leetcode.cn/problems/unique-paths-ii/
package main

import "fmt"

// uniquePathsWithObstacles 返回带障碍网格中从左上到右下的不同路径数。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n = len(obstacleGrid), len(obstacleGrid[0])
	var dp = make([][]int, m)

	// dp[-1][-1] = 1

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			for k := i; k < m; k++ { // 🟡 冗余：dp 创建时已全零，break 即可
				dp[k][0] = 0
			}
			break
		}
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			for k := j; k < n; k++ { // 🟡 冗余：同上
				dp[0][k] = 0
			}
			break
		}
		dp[0][j] = 1
	}

	fmt.Println(dp) // 🟡 调试代码，应删除

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	fmt.Println(dp) // 🟡 调试代码，应删除

	return dp[m-2][n-1] + dp[m-1][n-2] // 🔴 应该返回 dp[m-1][n-1]，不是"终点上方+终点左方"
}

// 【V0 点评】
// 编译通过，示例碰巧正确，但返回值逻辑错误：
//   🔴 第 65 行：return dp[m-2][n-1] + dp[m-1][n-2]
//      应该 return dp[m-1][n-1]。DP 已经算好了每个位置的路径数，
//      直接返回终点即可。当前写法在 3×3 示例碰巧正确
//      （因为 dp[2][2] = dp[1][2] + dp[2][1]），但更大网格或有障碍时会错
//   🔴 起点是障碍时未处理（obstacleGrid[0][0]==1 应返回 0）
//   🟡 首行首列初始化中，障碍后面的 for 循环冗余（dp 已全零，break 即可）
//   🟡 fmt.Println(dp) 是调试代码，应删除
//   🟡 main 中变量名 nuns 拼写错误，应为 grid

func main() {
	var grid = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	println("V0:", uniquePathsWithObstacles(grid))    // 2
	println("V1:", uniquePathsWithObstaclesV1(grid))  // 2
}

// uniquePathsWithObstaclesV1 修复版：标准 DP，时间 O(mn)，空间 O(n)
func uniquePathsWithObstaclesV1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0 // 起点或终点有障碍，直接返回 0
	}

	dp := make([]int, n)
	dp[0] = 1 // 起点

	// 初始化第一行：遇到障碍后面全是 0
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0
		} else {
			dp[j] = dp[j-1] // 只能从左方来
		}
	}

	// 逐行更新
	for i := 1; i < m; i++ {
		// 每行第一个元素：遇到障碍置 0，否则保持上一行的值（只能从上方来）
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0 // 障碍位置路径数为 0
			} else {
				dp[j] = dp[j] + dp[j-1] // 上方 + 左方
			}
		}
	}

	return dp[n-1]
}

// 【V1 点评 - 修复版（滚动数组 DP）】
// ✅ 修复了返回值：直接返回 dp[n-1]
// ✅ 修复了起点障碍：开头检查 obstacleGrid[0][0]
// ✅ 空间优化到 O(n)：用一维滚动数组代替二维数组
// ✅ 删除了调试代码
// ✅ 首行初始化正确处理障碍传播
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【面试策略】
//
// 面试时写二维数组 DP 完全可以，不会有影响！
//
// 面试官看重的是：
//   1. 思路正确：状态转移方程 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//   2. 边界处理：首行首列初始化、障碍处理
//   3. 代码清晰：能跑、能讲清楚
//
// 二维数组 vs 滚动数组只是空间优化，属于"加分项"不是"必选项"。
//
// 【推荐面试流程】
//   第一遍：写二维数组 DP（清晰易懂，不容易出错）
//      ↓
//   面试官问："能优化空间吗？"
//      ↓
//   第二遍：改成滚动数组（展示优化能力）
//
// 【优先级排序】
//   1. 思路正确，能讲清楚 DP 状态转移
//   2. 代码能跑，边界处理正确
//   3. 时间复杂度 O(mn)
//   4. 空间优化到 O(n)（加分项）
//
// 【数据范围考虑】
//   LeetCode 63 题：100×100，二维数组完全没问题
//   如果网格特别大（10000×10000），才需要考虑滚动数组避免 MLE
// ──────────────────────────────────────────────────────────
