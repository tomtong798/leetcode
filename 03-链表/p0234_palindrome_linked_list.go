//go:build p0041

// LeetCode 第234题 - 回文链表 (Palindrome Linked List)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】栈、递归、链表、双指针
//
// 【题目描述】
// 给你一个单链表的头节点 head，请你判断该链表是否为回文链表。如果是，返回 true；否则返回 false。
//
// 【示例 1】
// 输入：head = [1,2,2,1]
// 输出：true
//
// 【示例 2】
// 输入：head = [1,2]
// 输出：false
//
// 【数据范围】
// 链表中节点数目在范围 [1, 10^5] 内
// 0 <= Node.val <= 9
//
// 【思路提示】快慢找中点 + 反转后半段逐一比对（恢复现场可选）。
//
// 题目链接：https://leetcode.cn/problems/palindrome-linked-list/
package main

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome 判断链表是否回文。
func isPalindrome(head *ListNode) bool {
	// TODO: 实现
	return false
}

func main() {
	isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}})
}
