//go:build p0041

// LeetCode 第337题 - 打家劫舍 III (House Robber III)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树/动态规划/二叉树
//
// 【题目描述】
// 小偷发现房屋排列成二叉树，直接相连的两个房屋不能同一晚偷窃，求能偷窃到的最高金额。
//
// 【示例 1】
// 输入：root = [3,2,3,null,3,null,1]
// 输出：7（偷 3+3+1）
// 【数据范围】
// 树中节点数范围 [1, 10^4]；0 <= Node.val <= 10^4
//
// 【思路提示】树形 DP：后序遍历返回 [偷当前, 不偷当前] 两值，父节点由子节点合并。
//
// 题目链接：https://leetcode.cn/problems/house-robber-iii/
package main

// TreeNode 二叉树节点定义。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob337 返回树形房屋能偷窃到的最高金额。
func rob337(root *TreeNode) int {
	// TODO: 实现
	return 0
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	println(rob337(root)) // 7
}
