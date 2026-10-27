//go:build p0041

// LeetCode 第2题 - 两数相加 (Add Two Numbers)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】递归、链表、数学
//
// 【题目描述】
// 给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，
// 并且每个节点只能存储一位数字。请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 【示例 1】
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807。
//
// 【示例 2】
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
// 【数据范围】
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// 【思路提示】同时遍历两链表，逐位相加记进位；虚拟头节点简化边界。
//
// 题目链接：https://leetcode.cn/problems/add-two-numbers/
package main

import "fmt"

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 两链表逆序相加。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	var cur = result
	var jinwei int = 0
	for l1 != nil || l2 != nil {
		var v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		fmt.Println(v1, v2) // 🟡 调试代码，应删除
		jinwei = (v1 + v2 + jinwei) / 10 // 🔴 BUG：jinwei 先被更新了
		cur.Val = (v1 + v2 + jinwei) % 10 // 🔴 BUG：用的是新的 jinwei，不是旧进位
		fmt.Println(jinwei, cur.Val) // 🟡 调试代码，应删除
		if jinwei == 0 && l1 == nil && l2 == nil {
			cur.Next = nil
			break
		}
		cur.Next = &ListNode{} // 🟡 每次循环末尾预创建空节点，结束时可能多余
	}

	// 以下为废弃的旧代码，应删除
	// var cur = result
	// for i := 0; i < longLen-shortLen; i++ {
	// 	var next *ListNode = &ListNode{}
	// 	cur.Val = long.Val
	// 	long = long.Next
	// 	cur.Next = next
	// }

	// var jinwei int = 0
	// fmt.Println(jinwei, short, long)
	// for j := 0; j < shortLen; j++ {
	// 	var next *ListNode = &ListNode{}
	// 	jinwei = (long.Val + short.Val + jinwei) / 10
	// 	cur.Val = (long.Val + short.Val + jinwei) % 10
	// 	long = long.Next
	// 	short = short.Next
	// 	cur.Next = next
	// }

	return result
}

// 【V0 点评】
// 编译通过，但计算结果错误：
//   🔴 第 55-56 行：jinwei 先更新后使用，导致 cur.Val 用了新进位
//      例如 4+6=10，jinwei 先变成 1，cur.Val = (4+6+1)%10 = 1，应该是 0
//      修复：先算 sum := v1+v2+jinwei，再 cur.Val=sum%10，jinwei=sum/10
//   🔴 循环末尾预创建空节点，正常结束时最后一个节点多挂一个空节点
//   🟡 fmt.Println 是调试代码，应删除
//   🟡 第 65-82 行是废弃的旧代码，应删除
//   🟡 测试用例不足，只测了 [2]+[5,2]，没有测试进位场景
//   ✅ 虚拟头节点思路正确
//   ✅ 处理了链表长度不同的情况（v1/v2 默认 0）

func main() {
	// 342 + 465 = 807
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	printList(addTwoNumbersV1(l1, l2)) // [7,0,8]

	// 99 + 1 = 100（测试连续进位）
	l3 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l4 := &ListNode{Val: 1}
	printList(addTwoNumbersV1(l3, l4)) // [0,0,1]
}

// addTwoNumbersV1 修复版：标准迭代，时间 O(max(m,n))，空间 O(1)
func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // ✅ 虚拟头节点，简化边界
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 { // ✅ 三个条件：l1/l2 未遍历完 或 还有进位
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10     // ✅ 先算 sum，再分离 carry 和 digit
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}

// 【V1 点评 - 修复版】
// ✅ 修复了进位 bug：先算 sum，再 carry=sum/10，digit=sum%10
// ✅ 循环条件 carry > 0：正确处理最后进位（如 99+1=100）
// ✅ 不再预创建空节点：每次循环创建新节点并挂到 cur.Next
// ✅ 删除了调试代码和废弃注释
// ✅ 用 dummy 虚拟头节点，返回 dummy.Next，逻辑更清晰
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
// | 解法                | 时间          | 空间   | 面试期望    |
// |---------------------|---------------|--------|-------------|
// | 迭代（V0 有 bug）    | O(max(m,n))   | O(1)   | ❌ 进位错误  |
// | 迭代修复（V1）       | O(max(m,n))   | O(1)   | ✅ 推荐写法  |
// | 递归                | O(max(m,n))   | O(m+n) | ⚠️ 可以     |
//
// 【核心易错点】
// 1. 进位顺序：必须先算 sum，再 sum%10 和 sum/10，不能先更新 carry
// 2. 最后进位：循环条件要包含 carry > 0，否则 99+1 会漏掉最高位
// 3. 长度不同：短链表用 0 补齐，不要特殊处理
//
// 【虚拟头节点 vs 不用虚拟头节点】
//
// 不用虚拟头节点（V0 的写法）：
//   var result = &ListNode{}  // 预先创建第一个节点
//   var cur = result
//   for ... {
//       cur.Val = sum % 10    // 直接给第一个节点赋值
//       cur.Next = &ListNode{} // 后续节点通过 Next 创建
//       cur = cur.Next
//   }
//   return result
//
// 用虚拟头节点（V1 的写法）：
//   dummy := &ListNode{}      // 虚拟头节点，不存储有效数据
//   cur := dummy
//   for ... {
//       cur.Next = &ListNode{Val: sum % 10}  // 所有节点统一通过 Next 创建
//       cur = cur.Next
//   }
//   return dummy.Next
//
// 对比：
//   | 方面          | 不用虚拟头节点        | 用虚拟头节点          |
//   |---------------|-----------------------|-----------------------|
//   | 能否工作       | ✅ 能                 | ✅ 能                 |
//   | 代码一致性     | ❌ 第一个节点特殊处理 | ✅ 所有节点统一处理   |
//   | 边界处理       | ⚠️ 需要小心           | ✅ 更简单             |
//   | 面试印象       | ⚠️ 一般               | ✅ 更专业             |
//
// 结论：不用虚拟头节点可以，但代码风格不够优雅，面试官可能觉得不够熟练。
// 建议养成用虚拟头节点的习惯，尤其在：
//   - 可能需要删除/修改头节点的场景
//   - 链表可能为空的场景
//   - 多链表合并的场景
//
// 【面试讲解建议】
// 1. 先说思路："模拟竖式加法，逐位相加，记录进位"
// 2. 提到虚拟头节点简化边界处理
// 3. 强调循环条件包含 carry > 0（面试官常问的边界）
// 4. 分析复杂度：时间 O(max(m,n))，空间 O(1)
// ──────────────────────────────────────────────────────────
