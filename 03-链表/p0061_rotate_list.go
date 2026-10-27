//go:build p0041

// LeetCode 第61题 - 旋转链表 (Rotate List)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】链表、双指针
//
// 【题目描述】
// 给你一个链表的头节点 head，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 【示例 1】
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
//
// 【示例 2】
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
//
// 【数据范围】
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10^9
//
// 【思路提示】连成环：先求长度 len，k %= len，从新头走 len-k 步断开。
//
// 题目链接：https://leetcode.cn/problems/rotate-list/
package main

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// rotateRight 向右旋转链表 k 个位置。
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
func rotateRight(head *ListNode, k int) *ListNode {
	var lenght int // 🟡 拼写错误：应为 length
	var cur = head

	// 找到截断的节点、尾节点

	var tail *ListNode
	for cur != nil {
		lenght++
		tail = cur
		cur = cur.Next
	}

	if lenght == 0 { // 🔴 缺少空链表判断，下面 k % lenght 会除零 panic
		return nil
	}

	var rightMove = k % lenght
	var newHead = head
	var newHeadPre *ListNode
	for i := 0; i < rightMove; i++ { // 🔴 BUG：向右旋转应该走 length-rightMove 步
		newHeadPre = newHead          //    当前走 rightMove 步，变成向左旋转了
		newHead = newHead.Next
	}

	if newHeadPre != nil {
		tail.Next = head
		newHeadPre.Next = nil
	}

	return newHead
}

// 【V0 点评】
// 编译通过，但结果错误：
//   🔴 旋转方向反了：走 rightMove 步是左旋，应该走 length-rightMove 步
//      例如 [1,2,3,4,5] k=2：当前输出 [3,4,5,1,2]（左旋），期望 [4,5,1,2,3]（右旋）
//   🔴 空链表时 k % lenght 除零 panic，需要加 lenght==0 判断
//   🟡 变量名 lenght 拼写错误，应为 length
//   🟡 main 没有输出结果
//   ✅ 连成环再断开的思路正确
//   ✅ k % length 处理大 k 值正确

func main() {
	printList(rotateRightV1(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)) // [4,5,1,2,3]
	printList(rotateRightV1(&ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, 4)) // [2,0,1]
}

// rotateRightV1 修复版：成环再断开，时间 O(n)，空间 O(1)
func rotateRightV1(head *ListNode, k int) *ListNode {
	if head == nil { // ✅ 空链表判断
		return nil
	}

	length := 1 // ✅ 从 1 开始，因为 tail 已经指向最后一个节点
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	rightMove := k % length
	if rightMove == 0 { // ✅ 不需要旋转
		return head
	}

	// 成环
	tail.Next = head

	// 找到新尾节点：走 length - rightMove - 1 步
	newTail := head
	for i := 0; i < length-rightMove-1; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next // ✅ 新头
	newTail.Next = nil      // ✅ 断开

	return newHead
}

// 【V1 点评 - 修复版】
// ✅ 修复了旋转方向：走 length-rightMove 步到新头
// ✅ 修复了空链表：加了 head == nil 判断
// ✅ 优化了遍历：一次遍历同时求长度和找尾节点
// ✅ 加了 rightMove == 0 的优化：不需要旋转时直接返回
// ✅ 用 length 从 1 开始，tail 从 head 开始，减少一次遍历
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
// | 成环断开（V0 有 bug）    | O(n)   | O(1)   | ❌ 方向反了      |
// | 成环断开修复（V1）       | O(n)   | O(1)   | ✅ 推荐写法      |
// | 快慢指针                | O(n)   | O(1)   | ⚠️ 可以          |
//
// 【核心易错点】
// 1. 右旋 k 步 = 最后 k 个节点移到前面，不是前 k 个移到后面
// 2. 新尾节点位置：length - k - 1（从 0 开始计数）
// 3. 空链表和 k % length == 0 的边界处理
// 4. k 可能很大（2*10^9），必须先 k %= length
//
// 【面试讲解建议】
// 1. 先说思路："连成环，再在正确位置断开"
// 2. 画图说明旋转后新头新尾的位置
// 3. 强调 k %= length 处理大 k
// 4. 分析复杂度：时间 O(n)，空间 O(1)
// ──────────────────────────────────────────────────────────
