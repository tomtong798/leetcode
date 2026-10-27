//go:build p0041

// LeetCode 第399题 - 除法求值 (Evaluate Division)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】深度优先搜索、广度优先搜索、并查集、图、数组、字符串
//
// 【题目描述】
// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
// 其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i]。
// 每个 Ai 或 Bi 是一个表示单个变量的字符串。
// 另一些以 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，
// 请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。如果存在某个无法确定的答案，则用 -1.0 替代。
//
// 【示例 1】
// 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0],
//      queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
// 解释：a/b=2.0, b/c=3.0 → a/c=6.0, b/a=0.5
//
// 【数据范围】
// 1 <= equations.length <= 20
// 1 <= queries.length <= 20
// 方程式及结果中元素均为正数
//
// 【思路提示】带权有向图：a→b 权 w, b→a 权 1/w；
// 查询即从起点 DFS/BFS 找到终点的路径权值乘积；或带权并查集。
//
// 题目链接：https://leetcode.cn/problems/evaluate-division/
package main

// calcEquation 根据已知等式求查询结果。
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// TODO: 实现
	return nil
}

func main() {
	calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}})
}
