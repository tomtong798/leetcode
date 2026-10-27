//go:build p0041

// LeetCode 第114题 - 二叉树展开为链表 (Flatten Binary Tree to Linked List)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、树、深度优先搜索、链表、二叉树
//
// 【题目描述】
// 给你二叉树的根结点 root，请你将它展开为一个单链表：
//   - 展开后的单链表应该同样使用 TreeNode，其中 right 子指针指向链表中下一个结点，而 left 子指针始终为 null
//   - 展开后的单链表应该与二叉树先序遍历顺序相同
//
// 【示例 1】
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
//
// 【示例 2】
// 输入：root = []
// 输出：[]
//
// 【数据范围】
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// 【思路提示】方法一：逆后序（右左根）递归记录前驱；方法二（原地 O(1)）：
// 找当前节点左子树最右节点，接上原右子树，左子树移到右边，左置空，沿 right 前进。
//
// 题目链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flatten 将二叉树展开为链表（原地）。
func flatten(root *TreeNode) {
	// TODO: 实现

	flatten(root.Right) // 🔴 BUG：root 可能为 nil（子树为空/空树），直接 panic
	flatten(root.Left)  //    缺递归基 if root == nil { return }

	right := root.Right
	root.Right = root.Left
	root.Left = nil

	var last, pre *TreeNode = root, nil // 🟡 pre 声明在外，靠循环推进——空链表时 pre 为 nil
	for last != nil {                   //    （本函数 root 非 nil 时不会发生，但写法脆弱）
		pre = last
		last = last.Right
	}

	pre.Right = right // 🟡 尾部遍历 O(h)：整个算法最坏 O(n²)（链状左树时每层都走到尾）
}

// 【V0 点评】
// 编译通过，但运行必 panic；补上递归基后逻辑正确：
//   🔴 BUG：缺递归基——flatten(nil) 或任何空子树都会 panic: nil pointer dereference
//      main 只测了 flatten(nil)，一跑就崩；补 if root == nil { return } 即可
//   ✅（修复后）递归顺序正确：先展右、再展左、最后拼接——因为拼接时
//      要找"新右链（原左子树）"的尾节点，左子树必须已完全展平
//   ✅ 拼接三步正确：右 = 左、左置 nil、旧右链接到新右链尾部
//   🟡 效率：每个节点都要沿 right 走到链尾找 pre，最坏（左偏链）O(n²)
//      标准优化：a) 前序遍历 + 反向拼接（O(n)）；b) 原地 O(1) 空间的 Morris 式（见 V1/V2）
//   🟡 逆后序（右→左→根）是这题的巧思解法，你的"后处理拼接"是等价变形，思路对
//   🟡 main 只测了 nil（必 panic），没有验证正确性（见下方 main）
// 修复递归基后，用 [1,2,5,3,4,null,6] 验证输出 1→2→3→4→5→6 ✓

// 对上述代码的修改
func flattenV4(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归展开左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 保存右子树
	right := root.Right

	// 将左子树移到右子树位置
	root.Right = root.Left
	root.Left = nil

	// 找到新右子树的最右节点
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	// 将原右子树接到最右节点
	cur.Right = right
}

// flattenV1 【LeetCode 标准答案】前序遍历 + 前驱拼接，时间 O(n)，空间 O(h)
func flattenV1(root *TreeNode) {
	var prev *TreeNode // 全局记录"已构建链表的尾节点"（上一个访问的节点）

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return // ✅ 递归基：空子树直接返回
		}
		dfs(node.Right) // ✅ 逆前序 = 右→左→根：先处理右边
		dfs(node.Left)  // ✅ 再左边

		// 此时左右子树都已展平成链，且 prev 指向"应排在我后面的链头"
		node.Right = prev // ✅ 我接到已构建链的前面
		node.Left = nil   // ✅ 左指针清空（题目要求）
		prev = node       // ✅ 我成为新的链头
	}
	dfs(root)
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 精髓：按"根→左→右"的【逆序】访问（右→左→根），每访问一个节点就把它
//    头插到已构建链表前面——逆序访问 + 头插 = 正序结果
// ✅ 不需要找尾节点：prev 始终是链头（等价于你 V0 里"走到尾"的信息，但 O(1) 拿到）
// ✅ 时间 O(n) 每个节点访问一次；空间 O(h) 递归栈
// 追踪 [1,2,5,3,4,null,6]（前序 1,2,3,4,5,6，逆序访问 6,5,4,3,2,1）：
//   访问 6：6.Right=nil(prev), prev=6          链：6
//   访问 5：5.Right=6, prev=5                  链：5→6
//   访问 4：4.Right=5, prev=4                  链：4→5→6
//   访问 3：3.Right=4, prev=3                  链：3→4→5→6
//   访问 2：2.Right=3, prev=2                  链：2→3→4→5→6
//   访问 1：1.Right=2, prev=1                  链：1→2→3→4→5→6 ✅
// 与 V0 对比：你的"递归后找尾拼接"和 V1 的"逆序访问头插"是同一信息的两种获取方式——
//   V0 每次花 O(h) 走到尾部（慢），V1 用 prev 变量 O(1) 记住（快）

// flattenV2 原地 O(1) 空间解法（Morris 式），时间 O(n)，空间 O(1)
func flattenV2(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// 找左子树的最右节点（前序意义下 cur 的直接前驱）
			rightmost := cur.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}
			// 拼接：前驱的右 = 原右子树；左子树整体搬到右边
			rightmost.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right // 沿 right 前进到下一个待处理节点
	}
}

// 【V2 点评 - 原地 O(1) 解法】
// ✅ 核心洞察：前序中 cur 的下一个节点，是它左子树的"最后一个节点"（最右路径）
// ✅ 三步拼接：左子树最右节点接上原右子树 → 左子树搬到右 → 左置空
// ✅ 没有递归没有栈，空间 O(1)——本题最优空间
// ✅ 时间 O(n)：每个节点被"访问"不超过 2 次（一次作为 cur，一次作为最右被找）
// 追踪 [1,2,5,3,4,null,6]：
//   cur=1：左=2子树，最右=4 → 4.Right=5子树, 1.Right=2子树, 1.Left=nil
//          树变为 1→2→3→4→(5→6)，cur 前进
//   cur=2：左=3子树，最右=3（3无右） → 3.Right=nil(原右nil), 2.Right=3, 左=nil
//   cur=3,4：无左，直接前进
//   cur=5：左=nil（原5的左本来就是null）... 5.Right=6 已接好，前进
//   cur=6：无左 → 前进 → nil 结束 ✅ 最终 1→2→3→4→5→6
// 面试价值：被追问"能不能 O(1) 空间"时写出这个，是本题的满分答案

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间   | 空间   | 面试期望      |
// |-----------------------------|--------|--------|---------------|
// | 后处理找尾拼接（V0 有 bug）  | O(n²)  | O(h)   | ❌ 补基后可用 |
// | 逆前序+头插（V1 标准答案）   | O(n)   | O(h)   | ✅ 面试首选   |
// | Morris 原地（V2）            | O(n)   | O(1)   | ✅ 满分备选   |
//
// 【核心易错点】
// 1. 递归基必须有：flatten 会在空子树上被调用，root==nil 直接 panic
// 2. 拼接前左右子树必须已展平：V0/V1 都靠"先递归后拼接"保证
// 3. 找尾 O(h) 的代价：V0 每层找尾，左偏链退化 O(n²)——用 prev 变量替代（V1）
// 4. 逆序访问 + 头插 = 正序链表：这是很多"链表重建"题的通用技巧
// 5. Morris 式的关键：cur 的前序后继 = 左子树最右节点
//
// 【面试讲解建议】
// 1. 先说思路："前序遍历顺序，逆序访问每个节点头插到链前，prev 记链头"
// 2. 讲为什么 O(n)："prev 替代了每次找尾的 O(h) 遍历"
// 3. 被追问 O(1) 空间时：讲 Morris 式"左子树最右节点接原右子树"
// 4. 复杂度：V1 O(n)/O(h)；V2 O(n)/O(1)
// ──────────────────────────────────────────────────────────

// printFlatten 沿 right 打印展开后的链表，用于验证
func printFlatten(root *TreeNode) {
	for root != nil {
		fmt.Printf("%d", root.Val)
		if root.Right != nil {
			fmt.Print(" -> ")
		}
		root = root.Right
	}
	fmt.Println()
}

// buildTree114 按"左孩子、右孩子"数组快速建树（这里手工建更直观）
func main() {
	// 示例：[1,2,5,3,4,null,6] → 1→2→3→4→5→6
	build := func() *TreeNode {
		return &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	}

	// flatten(nil) // V0 原 main：一跑就 panic（缺递归基），注释掉保平安

	t1 := build()
	flattenV1(t1)
	printFlatten(t1) // 1 -> 2 -> 3 -> 4 -> 5 -> 6

	t2 := build()
	flattenV2(t2)
	printFlatten(t2) // 1 -> 2 -> 3 -> 4 -> 5 -> 6

	t3 := build()
	// t3 修复版 V0（用户思路 + 递归基）：直接在 copy 上试
	var flattenFixed func(root *TreeNode)
	flattenFixed = func(root *TreeNode) {
		if root == nil {
			return // 补上的递归基
		}
		flattenFixed(root.Right)
		flattenFixed(root.Left)

		right := root.Right
		root.Right = root.Left
		root.Left = nil

		var last, pre *TreeNode = root, nil
		for last != nil {
			pre = last
			last = last.Right
		}
		pre.Right = right
	}
	flattenFixed(t3)
	printFlatten(t3) // 1 -> 2 -> 3 -> 4 -> 5 -> 6
}
