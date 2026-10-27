//go:build p0041

// LeetCode 第133题 - 克隆图 (Clone Graph)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】深度优先搜索、广度优先搜索、图、哈希表
//
// 【题目描述】
// 给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）。
// 图中的每个节点都包含它的值 val（int）和其邻居的列表（[]Node）。
//
// 【示例 1】
// 输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
// 输出：[[2,4],[1,3],[2,4],[1,3]]
// 解释：图中有 4 个节点。节点 1 的值是 1，它有两个邻居节点 2 和 4。
//
// 【数据范围】
// 节点数不超过 100
// 1 <= Node.val <= 100
// Node.val 互不相同
// 无向图，无自环、无重边
//
// 【思路提示】哈希表 原节点→克隆节点，DFS/BFS 遍历建邻接关系。
//
// 题目链接：https://leetcode.cn/problems/clone-graph/
package main

// Node 无向图节点。
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph 深拷贝无向图。
func cloneGraph(node *Node) *Node {
	// TODO: 实现
	return nil
}

func main() {
	cloneGraph(nil)
}
