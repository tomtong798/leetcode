//go:build p0041

// LeetCode 第82题 - 删除排序链表中的重复元素 II (Remove Duplicates from Sorted List II)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】链表、双指针
//
// 【题目描述】
// 给定一个已排序的链表的头 head，删除原始链表中所有重复数字的节点，只留下不同的数字。
// 返回已排序的链表。
//
// 【示例 1】
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
//
// 【示例 2】
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
//
// 【数据范围】
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
//
// 【思路提示】虚拟头节点 + 前驱指针，值相同的整段跳过。
//
// 题目链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
package main

import "fmt"

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates 删除所有重复节点。
// 【示例 1】
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
//
// 【示例 2】
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy = &ListNode{}
	dummy.Next = head

	for head.Next != nil { // 🔴 BUG：没有用 pre 指针，无法跳过重复段
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next // 🔴 BUG：只跳过了重复的第二个，第一个还在
		} else {
			head = head.Next
		}
	}

	return dummy.Next
}

// 【V0 点评】
// 编译通过，但结果错误：
//   🔴 核心错误：只跳过了重复的第二个，没有删除第一个
//      例如 [1,2,3,3,4,4,5]：
//      - 遇到 3==3 时，head.Next = 3.Next，链表变成 [1,2,3,4,4,5]
//      - 第一个 3 还在！应该整个 3 段都被删除
//   🔴 没有用 pre 指针追踪不重复的最后一个节点
//      无法在发现重复时跳过整个重复段
//   🔴 题目要求是"删除所有重复数字的节点"，不是"去重保留一个"
//      当前代码逻辑是去重（保留一个），不是删除重复（全部删除）
//   🟡 dummy 创建了但没用到（没有 pre 指针）
//   ✅ 空链表和单节点判断正确

func main() {
	printList(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5}}}}}}})) // [1,2,5]
	// printList(deleteDuplicatesV1(&ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val: 4,
	// 						Next: &ListNode{
	// 							Val: 5}}}}}}})) // [1,2,5]

	// printList(deleteDuplicatesV1(&ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 1,
	// 		Next: &ListNode{
	// 			Val: 1,
	// 			Next: &ListNode{
	// 				Val: 2,
	// 				Next: &ListNode{Val: 3}}}}})) // [2,3]
}

// deleteDuplicatesV1 修复版：虚拟头节点 + pre 指针，时间 O(n)，空间 O(1)
func deleteDuplicatesV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy // ✅ pre 追踪不重复的最后一个节点

	for pre.Next != nil {
		cur := pre.Next
		// 找到重复段的末尾
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			// 没有重复，pre 前进
			pre = pre.Next
		} else {
			// 有重复，跳过整个重复段
			pre.Next = cur.Next
		}
	}

	return dummy.Next
}

// 【V1 点评 - 修复版】
// ✅ 用 pre 指针追踪不重复的最后一个节点
// ✅ 发现重复时，跳过整个重复段（pre.Next = cur.Next）
// ✅ 没有重复时，pre 前进（pre = pre.Next）
// ✅ 正确处理了头节点重复的情况（如 [1,1,1,2,3]）
// ✅ 一次遍历，时间 O(n)，空间 O(1)
// 面试推荐写法

// printList 辅助函数
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
// | 去重（V0 有 bug）        | O(n)   | O(1)   | ❌ 逻辑错误      |
// | 删除重复修复（V1）       | O(n)   | O(1)   | ✅ 推荐写法      |
// | 递归                    | O(n)   | O(n)   | ⚠️ 可以          |
//
// 【核心易错点】
// 1. 题目要求"删除所有重复数字的节点"，不是"去重保留一个"
//    - 去重：[1,2,3,3,4] → [1,2,3,4]
//    - 删除重复：[1,2,3,3,4] → [1,2,4]
// 2. 必须用 pre 指针追踪不重复的最后一个节点
// 3. 头节点可能重复（如 [1,1,1,2,3]），需要虚拟头节点
// 4. 发现重复时，要跳过整个重复段，不是只跳过第二个
//
// 【面试讲解建议】
// 1. 先说思路："用 pre 指针追踪不重复的最后一个节点"
// 2. 画图说明发现重复时如何跳过整段
// 3. 强调虚拟头节点处理头节点重复的情况
// 4. 分析复杂度：时间 O(n)，空间 O(1)
// ──────────────────────────────────────────────────────────
