//go:build p0041

// LeetCode 第23题 - 合并 K 个升序链表 (Merge k Sorted Lists)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】链表、分治、堆（优先队列）、归并排序
//
// 【题目描述】
// 给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，
// 返回合并后的链表。
//
// 【示例 1】
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
//   1->4->5,
//   1->3->4,
//   2->6
// 将它们合并到一个有序链表中得到 1->1->2->3->4->4->5->6。
//
// 【示例 2】
// 输入：lists = []
// 输出：[]
//
// 【示例 3】
// 输入：lists = [[]]
// 输出：[]
//
// 【数据范围】
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// 【思路提示】三种解法：
// 1) 小顶堆：k 个头节点入堆，弹出最小接到结果，其 next 入堆，O(N log k)；
// 2) 分治：两两合并，递归折半；
// 3) 顺序两两合并（简单但 O(Nk)）。
//
// 题目链接：https://leetcode.cn/problems/merge-k-sorted-lists/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists 合并 k 个升序链表。
func mergeKLists(lists []*ListNode) *ListNode {
	// TODO: 实现
	return nil
}

func main() {
	mergeKLists(nil)
	mergeKLists([]*ListNode{{Val: 1, Next: &ListNode{Val: 4}}})
}
