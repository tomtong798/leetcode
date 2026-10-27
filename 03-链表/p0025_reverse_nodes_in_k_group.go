//go:build p0041

// LeetCode 第25题 - K 个一组翻转链表 (Reverse Nodes in k-Group)
//
// 【难度】困难 | 【面试频率】★★★★ | 【标签】链表、递归、迭代
//
// 【题目描述】
// 给你链表的头节点 head，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，
// 那么请将最后剩余的节点保持原有顺序。你不能只是单纯的改变节点内部的值，
// 而需要实际进行节点交换。
//
// 【示例 1】
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
//
// 【示例 2】
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
//
// 【数据范围】
// 链表长度范围 [1, 5000]
// 1 <= k <= 链表长度
//
// 【思路提示】虚拟头节点；每组先检查是否有 k 个节点，不足则结束；
// 组内用反转链表的迭代逻辑翻转，再把翻转后的子链表接回主链，prev 移动到原组头（翻转后组尾）。
//
// 题目链接：https://leetcode.cn/problems/reverse-nodes-in-k-group/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup 每 k 个一组翻转链表。
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO: 实现
	return nil
}

func main() {
	reverseKGroup(&ListNode{Val: 1}, 1)
}
