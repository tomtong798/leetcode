//go:build p0041

// LeetCode 第143题 - 重排链表 (Reorder List)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、递归、链表、双指针
//
// 【题目描述】
// 给定一个单链表 L 的头节点 head，单链表 L 表示为：L0 → L1 → … → Ln-1 → Ln
// 请将其重新排列后变为：L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 【示例 1】
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
//
// 【示例 2】
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
//
// 【数据范围】
// 链表中节点的数目在范围 [1, 5 * 10^4] 内
// 1 <= Node.val <= 1000
//
// 【思路提示】三步：快慢指针找中点 → 反转后半段 → 双指交替合并。
//
// 题目链接：https://leetcode.cn/problems/reorder-list/
package main

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList 重排链表。
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
func reorderList(head *ListNode) {
	// TODO: 实现
}

func main() {
	reorderList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}})
}
