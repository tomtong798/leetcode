//go:build p0041

// LeetCode 第141题 - 环形链表 (Linked List Cycle)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】链表、双指针、快慢指针
//
// 【题目描述】
// 给你一个链表的头节点 head，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达它，则链表中存在环。
//
// 进阶：你能用 O(1)（即常量）内存解决此问题吗？
//
// 【示例 1】
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
// 【示例 2】
// 输入：head = [1,2], pos = 0
// 输出：true
//
// 【示例 3】
// 输入：head = [1], pos = -1
// 输出：false
//
// 【数据范围】
// 链表中节点的数目范围是 [0, 10^4]
// -10^5 <= Node.val <= 10^5
//
// 【思路提示】Floyd 快慢指针：慢指针每次走 1 步、快指针走 2 步，若相遇则有环；
// 快指针到达 nil 即无环。
//
// 题目链接：https://leetcode.cn/problems/linked-list-cycle/
package main

import "fmt"

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 判断链表是否有环。
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil { // 🟡 冗余判断，后面循环会处理
		return false
	}

	var slow, fast = head, head.Next.Next // 🔴 BUG：fast 应该从 head 开始，不是 head.Next.Next
	for fast != nil {
		if slow.Val == fast.Val { // 🔴 BUG：应该比较指针 slow == fast，不是比较值
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}

	return false
}

// 【V0 点评】
// 编译通过，但逻辑错误：
//   🔴 第 50 行：比较值 slow.Val == fast.Val，应该比较指针 slow == fast
//      如果有两个不同节点值相同，会误判为有环
//      例如：1→2→3→2（值2出现两次），但实际无环，会误判
//   🔴 第 48 行：fast 初始化为 head.Next.Next，应该从 head 开始
//      标准做法：slow 和 fast 都从 head 开始，然后 slow 走 1 步，fast 走 2 步
//   🟡 第 44 行：head.Next.Next == nil 判断冗余，后面循环会处理
//   🟡 第 54-56 行：fast.Next == nil 判断可以合并到循环条件
//   ✅ 快慢指针思路正确

func main() {
	// 创建有环链表：3→2→0→-4→2
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2             // 形成环
	fmt.Println(hasCycleV1(node1)) // true

	// 无环链表：1→2→nil
	fmt.Println(hasCycleV1(&ListNode{Val: 1, Next: &ListNode{Val: 2}})) // false
}

// hasCycleV1 修复版：快慢指针，时间 O(n)，空间 O(1)
func hasCycleV1(head *ListNode) bool {
	slow, fast := head, head // ✅ 都从 head 开始

	for fast != nil && fast.Next != nil { // ✅ 快指针能走 2 步才继续
		slow = slow.Next      // 慢指针走 1 步
		fast = fast.Next.Next // 快指针走 2 步

		if slow == fast { // ✅ 比较指针，不是比较值
			return true
		}
	}

	return false
}

// 【V1 点评 - 修复版】
// ✅ 修复了比较逻辑：slow == fast 比较指针，不是比较值
// ✅ 修复了初始化：slow 和 fast 都从 head 开始
// ✅ 简化了循环条件：fast != nil && fast.Next != nil
// ✅ 时间 O(n)，空间 O(1)
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【为什么快慢指针可以判断链表中有环？】
//
// 直觉理解：
//   想象两个人在环形跑道上跑步：
//   - 慢指针：每分钟跑 1 圈
//   - 快指针：每分钟跑 2 圈
//   快指针一定会追上慢指针，因为每过一分钟，快指针就比慢指针多跑 1 圈。
//
// 数学证明：
//   关键观察：快指针每次比慢指针多走 1 步
//
//   第 0 步：慢=0, 快=0, 距离=0
//   第 1 步：慢=1, 快=2, 距离=1（快比慢多走1步）
//   第 2 步：慢=2, 快=4, 距离=2
//   第 k 步：距离=k
//
//   在环内，距离会"取模"：假设环长度为 L，实际距离是 k % L
//   当 k % L == 0 时，快慢指针相遇。
//
// 为什么不会"跳过"？
//   快指针每次只比慢指针多走 1 步，所以一定会"踩到"慢指针，不会跳过。
//
//   假设某时刻：慢指针在位置 5，快指针在位置 7（距离=2）
//   下一步：慢=6, 快=9（距离=3）
//   再下一步：慢=7, 快=11（距离=4）
//   ...
//   距离每次增加 1，不会从 0 直接跳到 2
//
// 为什么无环时不会误判？
//   如果无环，快指针会先到达 nil：
//   1→2→3→nil
//   第1步：慢=2, 快=3
//   第2步：慢=3, 快=nil（快指针到头了，无环）
//
// 画图理解：
//   假设链表：1→2→3→4→5→3（环：3→4→5→3，长度=3）
//
//   初始：slow=1, fast=1
//   第1步：slow=2, fast=3
//   第2步：slow=3, fast=5
//   第3步：slow=4, fast=4（相遇！）
//
// 核心总结：
//   | 问题 | 答案 |
//   |---|---|
//   | 为什么快指针能追上慢指针？ | 每次多走 1 步，距离每次减少 1 |
//   | 为什么不会跳过？ | 每次只多走 1 步，不会从 0 跳到 2 |
//   | 为什么无环时不会误判？ | 快指针会先到 nil |
//   | 时间复杂度？ | O(n)，最多走 n 步 |
//   | 空间复杂度？ | O(1)，只用两个指针 |
//
// 一句话记忆：
//   快指针每次比慢指针多走 1 步，在环内一定会追上（不会跳过），
//   无环时快指针先到 nil。
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法              | 时间   | 空间   | 面试期望        |
// |-------------------|--------|--------|-----------------|
// | 哈希表            | O(n)   | O(n)   | ⚠️ 可以          |
// | 快慢指针（V1）    | O(n)   | O(1)   | ✅ 推荐写法      |
//
// 【核心易错点】
// 1. 比较指针 slow == fast，不是比较值 slow.Val == fast.Val
// 2. slow 和 fast 都从 head 开始
// 3. 循环条件 fast != nil && fast.Next != nil（快指针能走 2 步）
// 4. 无环时快指针先到 nil，不会误判
//
// 【面试讲解建议】
// 1. 先说思路："快慢指针，快指针每次走 2 步，慢指针走 1 步"
// 2. 解释为什么能追上："每次距离减少 1，一定会相遇"
// 3. 画图说明相遇过程
// 4. 分析复杂度：时间 O(n)，空间 O(1)
// ──────────────────────────────────────────────────────────
