//go:build p0041

// LeetCode 第19题 - 删除链表的倒数第 N 个结点 (Remove Nth Node From End of List)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】链表、双指针
//
// 【题目描述】
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
//
// 【示例 1】
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 【示例 2】
// 输入：head = [1], n = 1
// 输出：[]
//
// 【示例 3】
// 输入：head = [1,2], n = 1
// 输出：[1]
//
// 【数据范围】
// 链表中结点的数目为 [1, 10^4]（进阶条件下 0~30 以内更常见）
// 0 <= n <= 链表结点数（实际保证 1 <= n <= Sz）
//
// 【思路提示】虚拟头节点 + 快慢指针：快指针先走 n+1 步，
// 然后快慢同步走，快指针到 nil 时慢指针正指向待删节点的前驱。
//
// 题目链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
package main

import "fmt"

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 删除倒数第 n 个节点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 倒数第 n 个节点，前后指针序号差为 n-1
	var left, right = head, head
	var rightIndex int
	var pre *ListNode

	for right != nil {
		if rightIndex >= n {
			pre = left
			left = left.Next
		}

		if right.Next != nil { // 🟡 rightIndex 递增方式不直观
			rightIndex += 1
		}
		right = right.Next
	}

	if pre == nil {
		return nil // 🔴 BUG：删除头节点时应返回 head.Next，不是 nil
	}
	pre.Next = right // 🔴 BUG：right 此时是 nil，应该 pre.Next = left.Next

	return head
}

// 【V0 点评】
// 编译通过，删除尾节点碰巧正确，但删除中间节点和头节点时错误：
//   🔴 第 63 行：pre.Next = right（right 是 nil）
//      应该 pre.Next = left.Next，跳过被删节点
//      删除尾节点时 left.Next == right == nil，碰巧正确
//   🔴 第 61 行：pre == nil 时 return nil
//      应该 return head.Next（删除头节点的情况）
//   🟡 rightIndex 递增方式不直观（只在 right.Next != nil 时递增）
//   ✅ 快慢指针思路正确

func main() {
	printList(removeNthFromEndV1(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)) // [1,2,3,5]
	printList(removeNthFromEndV1(&ListNode{Val: 1}, 1)) // []
	printList(removeNthFromEndV1(&ListNode{Val: 1, Next: &ListNode{Val: 3}}, 1)) // [1]
}

// removeNthFromEndV1 修复版：虚拟头节点 + 快慢指针，时间 O(L)，空间 O(1)
func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // ✅ 虚拟头节点，统一处理删除头节点的情况
	slow, fast := dummy, dummy

	// 快指针先走 n+1 步（多走 1 步，让慢指针停在待删节点的前驱）
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// 快慢同步走，快指针到 nil 时慢指针正好在待删节点前驱
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除节点
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 【V1 点评 - 修复版】
// ✅ 修复了删除逻辑：slow.Next = slow.Next.Next，正确跳过被删节点
// ✅ 修复了删除头节点：用虚拟头节点，不再需要特判
// ✅ 快指针先走 n+1 步，逻辑清晰直观
// ✅ 一趟扫描，时间 O(L)，空间 O(1)
// 面试推荐写法

// printList 辅助函数，打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间   | 空间   | 面试期望        |
// |-------------------------|--------|--------|-----------------|
// | 两趟扫描（先计数再删除） | O(L)   | O(1)   | ⚠️ 可以          |
// | 快慢指针（V0 有 bug）    | O(L)   | O(1)   | ❌ 删除逻辑错误  |
// | 快慢指针修复（V1）       | O(L)   | O(1)   | ✅ 推荐写法      |
//
// 【核心易错点】
// 1. 快指针先走 n+1 步（不是 n 步），让慢指针停在待删节点的前驱
// 2. 删除头节点的情况：用虚拟头节点统一处理，避免特判
// 3. 删除操作：slow.Next = slow.Next.Next，不是 slow.Next = nil
//
// 【面试讲解建议】
// 1. 先说思路："快慢指针，快指针先走 n+1 步，然后同步走"
// 2. 解释为什么多走 1 步：让慢指针停在待删节点的前驱，方便删除
// 3. 提到虚拟头节点统一处理边界
// 4. 分析复杂度：时间 O(L)，空间 O(1)，一趟扫描
// ──────────────────────────────────────────────────────────
