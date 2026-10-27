//go:build p0041

// LeetCode 第124题 - 二叉树中的最大路径和 (Binary Tree Maximum Path Sum)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、动态规划
//
// 【题目描述】
// 二叉树中的路径被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
// 同一个节点在一条路径序列中至多出现一次。该路径至少包含一个节点，且不一定经过根节点。
// 路径和是路径中各节点值的总和。给你一个二叉树的根节点 root，返回其最大路径和。
//
// 【示例 1】
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3，和为 2 + 1 + 3 = 6。
//
// 【示例 2】
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7，和为 15 + 20 + 7 = 42。
//
// 【数据范围】
// 树中节点数目范围 [1, 3 * 10^4] 内
// -1000 <= Node.val <= 1000
//
// 【思路提示】与"二叉树直径"同套路：递归返回"单边最大贡献"（max(左贡献,右贡献,0)+node.Val），
// 全局最大路径和 = 左贡献 + 右贡献 + node.Val，注意负贡献直接舍弃（取 0）。
//
// 题目链接：https://leetcode.cn/problems/binary-tree-maximum-path-sum/
package main

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxPathSum 返回二叉树最大路径和。
func maxPathSum(root *TreeNode) int {
	// TODO: 实现
	return 0
}

func main() {
	maxPathSum(nil)
}
