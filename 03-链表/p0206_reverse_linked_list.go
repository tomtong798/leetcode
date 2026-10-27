//go:build p0041

// LeetCode 第206题 - 反转链表 (Reverse Linked List)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】递归、链表
//
// 【题目描述】
// 给你单链表的头节点 head，请你反转链表并返回反转后的链表。
//
// 【示例 1】
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
//
// 【示例 2】
// 输入：head = [1,2]
// 输出：[2,1]
//
// 【数据范围】
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
// 【思路提示】迭代三指针 prev/cur/next 头插；递归写法：反转后半段后接回头节点。
//
// 题目链接：https://leetcode.cn/problems/reverse-linked-list/
package main

import "fmt"

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 反转链表。
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 1. 保存（先记住下一个在哪）
		cur.Next = prev  // 2. 反转（掉头指回去）
		prev = cur       // 3. 前进（prev 和 cur 都往前走一步）
		cur = next
	}

	return prev
}

// 【V0 点评 - 迭代法】
// ✅ 三步口诀："保存、反转、前进"
// ✅ 时间 O(n)，空间 O(1)
// ✅ 空链表和单节点判断正确
// ✅ 循环条件 cur != nil（不是 cur.Next != nil）
// ✅ 返回 prev（不是 cur，cur 最后是 nil）
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【为什么反转链表总是写错？】
//
// 1. 顺序依赖：必须先保存 next，否则反转后丢失后续节点
// 2. 指针方向反直觉：cur.Next = prev 是"往回指"，和正常遍历方向相反
// 3. 返回值容易搞混：返回 prev 还是 cur？（cur 最后是 nil）
//
// 【记忆技巧】
//
// 技巧1：三步口诀 "保存、反转、前进"
//   next := cur.Next   // 1. 保存（先记住下一个在哪）
//   cur.Next = prev    // 2. 反转（掉头指回去）
//   prev = cur         // 3. 前进（prev 和 cur 都往前走一步）
//   cur = next
//
//   类比：像翻书一样
//   - 保存：记住下一页在哪
//   - 反转：把当前页翻过去
//   - 前进：手移到下一页
//
// 技巧2：画图法（推荐）
//   每一步都画出来，不要靠脑子想：
//
//   初始：1 → 2 → 3 → nil
//         ↑
//        prev cur
//
//   第1轮：nil ← 1   2 → 3 → nil
//          ↑   ↑
//         prev cur
//
//   第2轮：nil ← 1 ← 2   3 → nil
//              ↑   ↑
//             prev cur
//
//   第3轮：nil ← 1 ← 2 ← 3
//                  ↑   ↑
//                 prev cur (nil)
//
//   返回 prev（指向3）
//
// 技巧3：固定模板（背下来）
//   func reverseList(head *ListNode) *ListNode {
//       var prev *ListNode
//       cur := head
//       for cur != nil {
//           next := cur.Next
//           cur.Next = prev
//           prev = cur
//           cur = next
//       }
//       return prev
//   }
//
//   关键记忆点：
//   - prev 初始是 nil（反转后尾节点指向 nil）
//   - 循环条件是 cur != nil（不是 cur.Next != nil）
//   - 返回 prev（不是 cur）
//
// 技巧4：递归法（如果迭代记不住）
//   func reverseList(head *ListNode) *ListNode {
//       if head == nil || head.Next == nil {
//           return head
//       }
//       newHead := reverseList(head.Next)
//       head.Next.Next = head  // 后一个节点指向前一个
//       head.Next = nil        // 前一个节点指向 nil
//       return newHead
//   }
//
//   记忆口诀："递归到底，反向指回"
//
// 【练习建议】
// 1. 先画图再写代码：不要直接写，先在纸上画出每一步的指针变化
// 2. 默写 3 遍：连续默写 3 遍迭代法，形成肌肉记忆
// 3. 对比记忆：把迭代和递归都写一遍，理解两者的等价性
//
// 核心：反转链表是肌肉记忆题，不是理解题。画 10 遍图就永远不会忘了。
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法              | 时间   | 空间   | 面试期望        |
// |-------------------|--------|--------|-----------------|
// | 迭代（V0）        | O(n)   | O(1)   | ✅ 推荐写法      |
// | 递归              | O(n)   | O(n)   | ⚠️ 可以          |
//
// 【核心易错点】
// 1. 必须先保存 next = cur.Next，否则反转后丢失后续节点
// 2. 循环条件 cur != nil，不是 cur.Next != nil
// 3. 返回 prev，不是 cur（cur 最后是 nil）
// 4. 空链表和单节点要特殊处理
//
// 【面试讲解建议】
// 1. 先说思路："三指针迭代，逐个反转指针方向"
// 2. 画图说明指针变化过程
// 3. 提到递归写法（但强调空间复杂度 O(n)）
// 4. 分析复杂度：时间 O(n)，空间 O(1)
// ──────────────────────────────────────────────────────────

func main() {
	printList(reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})) // [3,2,1]
	printList(reverseList(nil))                                                                 // []
}

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
