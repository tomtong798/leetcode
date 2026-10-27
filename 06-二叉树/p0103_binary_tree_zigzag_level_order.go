//go:build p0041

// LeetCode 第103题 - 二叉树的锯齿形层序遍历 (Binary Tree Zigzag Level Order Traversal)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】树、广度优先搜索、二叉树
//
// 【题目描述】
// 给你二叉树的根节点 root，返回其节点值的锯齿形层序遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 【示例 1】
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
//
// 【示例 2】
// 输入：root = [1]
// 输出：[[1]]
//
// 【数据范围】
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// 【思路提示】BFS 按层收集，偶数层（0-indexed）正序，奇数层反转；或用双端队列。
//
// 题目链接：https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
package main

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder 锯齿形层序遍历。
func zigzagLevelOrder(root *TreeNode) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	zigzagLevelOrder(nil)
}
