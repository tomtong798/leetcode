//go:build p0041

// LeetCode 第160题 - 相交链表 (Intersection of Two Linked Lists)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】链表、双指针
//
// 【题目描述】
// 给你两个单链表的头节点 headA 和 headB，请你找出并返回两个单链表相交的起始节点。
// 如果两个链表不存在相交节点，返回 null。整个链式结构中不存在环。
// 函数返回结果后，链表必须保持其原始结构。
//
// 【示例 1】
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
// 输出：Intersected at '8'
//
// 【示例 2】
// 输入：intersectVal = 2, listA = [1,9,1], listB = [3,2,3]（不相交）
// 输出：null
//
// 【数据范围】
// listA 中节点数目为 m，listB 中节点数目为 n
// 1 <= m, n <= 3 * 10^4
//
// 【思路提示】浪漫解法：两指针分别从 A、B 出发，走到头就换到另一条链表，
// 走过 a+c+b 与 b+c+a 步后必在相交点相遇（或同时为 nil）。
//
// 题目链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 返回两链表相交的起始节点。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// TODO: 实现
	return nil
}

func main() {
	getIntersectionNode(nil, nil)
}
