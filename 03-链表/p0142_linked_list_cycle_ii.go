//go:build p0041

// LeetCode 第142题 - 环形链表 II (Linked List Cycle II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】链表、双指针、快慢指针
//
// 【题目描述】
// 给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回 null。
// 不允许修改链表。
//
// 【示例 1】
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
// 【示例 2】
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
//
// 【示例 3】
// 输入：head = [1], pos = -1
// 输出：返回 null
//
// 【数据范围】
// 链表中节点的数目范围在范围 [0, 10^4] 内
// -10^5 <= Node.val <= 10^5
//
// 【思路提示】快慢指针相遇后，把一个指针放回头节点，两个指针每次各走 1 步，
// 再次相遇的节点即入环口（数学推导：头到入环口距离 = 相遇点绕环剩余距离）。
//
// 题目链接：https://leetcode.cn/problems/linked-list-cycle-ii/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 返回入环的第一个节点。
func detectCycle(head *ListNode) *ListNode {
	// TODO: 实现
	return nil
}

func main() {
	detectCycle(nil)
}
