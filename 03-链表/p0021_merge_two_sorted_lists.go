//go:build p0041

// LeetCode 第21题 - 合并两个有序链表 (Merge Two Sorted Lists)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】链表、双指针、递归
//
// 【题目描述】
// 将两个升序链表合并为一个新的升序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 【示例 1】
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 【示例 2】
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
// 【数据范围】
// 0 <= 两个链表的节点数目范围 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按非递减顺序排列
//
// 【思路提示】虚拟头节点 dummy + 尾插法：比较两链表当前值，小者接到结果尾部；
// 结束后把未走完的链表直接接上。也可写递归版本。
//
// 题目链接：https://leetcode.cn/problems/merge-two-sorted-lists/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists 合并两个升序链表。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil { // 🟡 冗余判断，后面逻辑已能处理
		return nil
	}

	var result = &ListNode{}

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val >= list2.Val { // 🔴 BUG：应该是 <=，升序合并选小的
				result.Next = &ListNode{Val: list1.Val} // 🔴 BUG：没有维护尾指针
				list1 = list1.Next
			} else {
				result.Next = &ListNode{Val: list2.Val} // 🔴 BUG：每次都覆盖 result.Next
				list2 = list2.Next
			}
			continue
		}

		if list1 != nil {
			result.Next = &ListNode{Val: list1.Val} // 🔴 BUG：同上，覆盖
			list1 = list1.Next
			continue
		}

		if list2 != nil {
			result.Next = &ListNode{Val: list2.Val} // 🔴 BUG：同上，覆盖
			list2 = list2.Next
			continue
		}
	}

	return result.Next
}

// 【V0 点评】
// 编译通过，但结果错误：
//   🔴 没有维护尾指针：每次 result.Next = ... 都覆盖之前的节点，
//      最终只保留最后一个节点，前面的全丢了
//      修复：用 cur 指针追踪尾部，cur = cur.Next
//   🔴 比较逻辑反了：list1.Val >= list2.Val 应该改为 <=
//      升序合并应该选小的，但当前选了大的
//   🟡 每次创建新节点浪费内存，应该直接复用原节点
//   🟡 三个 if 分支逻辑重复，可以简化
//   🟡 第 38-40 行空判断冗余
//   ✅ 循环条件 list1 != nil || list2 != nil 正确

func main() {
	mergeTwoLists(nil, nil)
}

// mergeTwoListsV1 修复版：虚拟头节点 + 尾指针，时间 O(m+n)，空间 O(1)
func mergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy // ✅ 维护尾指针

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val { // ✅ 升序合并，选小的
			cur.Next = list1 // ✅ 直接复用原节点，不创建新节点
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next // ✅ 移动到尾
	}

	// 接上剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

// 【V1 点评 - 修复版】
// ✅ 修复了尾指针：用 cur 追踪尾部，cur = cur.Next
// ✅ 修复了比较逻辑：list1.Val <= list2.Val，升序选小的
// ✅ 直接复用原节点，不创建新节点，节省内存
// ✅ 简化了逻辑：循环条件 list1 != nil && list2 != nil，结束后接剩余链表
// ✅ 用虚拟头节点 dummy，返回 dummy.Next
// 面试推荐写法

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间   | 空间   | 面试期望        |
// |-------------------------|--------|--------|-----------------|
// | 迭代（V0 有 bug）        | O(m+n) | O(1)   | ❌ 覆盖节点      |
// | 迭代修复（V1）           | O(m+n) | O(1)   | ✅ 推荐写法      |
// | 递归                    | O(m+n) | O(m+n) | ⚠️ 可以          |
//
// 【核心易错点】
// 1. 必须维护尾指针 cur，否则每次 result.Next = ... 会覆盖之前的节点
// 2. 升序合并：list1.Val <= list2.Val 时选 list1
// 3. 直接复用原节点，不要创建新节点
// 4. 循环结束后，把未走完的链表直接接上
//
// 【面试讲解建议】
// 1. 先说思路："双指针遍历两链表，小的接到结果尾部"
// 2. 提到虚拟头节点简化边界
// 3. 强调尾指针的重要性
// 4. 分析复杂度：时间 O(m+n)，空间 O(1)
// ──────────────────────────────────────────────────────────

func mergeTwoListsV2(list1 *ListNode, list2 *ListNode) *ListNode {
	var result = &ListNode{}
	var cur = result

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return result.Next
}
