//go:build p0041

// LeetCode 第92题 - 反转链表 II (Reverse Linked List II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】链表
//
// 【题目描述】
// 给你单链表的头指针 head 和两个整数 left 和 right，其中 left <= right。
// 请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表。
//
// 【示例 1】
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
//
// 【示例 2】
// 输入：head = [5], left = 1, right = 1
// 输出：[5]
//
// 【数据范围】
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
// 【思路提示】虚拟头节点，走到 left 前驱，头插法把区间内节点逐个搬到前面。
//
// 题目链接：https://leetcode.cn/problems/reverse-linked-list-ii/
package main

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseBetween 反转链表 [left, right] 区间。
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}

	var leftNode, rightNode *ListNode

	var dummy = &ListNode{Next: head}

	for i := 0; i < right; i++ {
		var pre = dummy
		var cur = dummy.Next
		if i >= left-1 {
			tmp := cur.Next.Next
			pre.Next = cur.Next
			cur.Next.Next = cur
			cur.Next = tmp

			leftNode = cur
		}
		rightNode = cur
	}

	return nil
}

func main() {
	reverseBetween(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2, 4)
}
