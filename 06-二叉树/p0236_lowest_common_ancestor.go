//go:build p0041

// LeetCode 第236题 - 二叉树的最近公共祖先 (Lowest Common Ancestor of a Binary Tree)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、二叉树
//
// 【题目描述】
// 给定一个二叉树，找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为："对于有根树 T 的两个节点 p、q，
// 最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大
// （一个节点也可以是它自己的祖先）。"
//
// 【示例 1】
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3。
//
// 【示例 2】
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出：5
// 解释：节点 5 和节点 4 的最近公共祖先是节点 5（节点也可以是自己的祖先）。
//
// 【数据范围】
// 树中节点数目在范围 [2, 10^5] 内
// -10^9 <= Node.val <= 10^9
// 所有节点值互不相同；p、q 为树中存在且值不同的节点
//
// 【思路提示】递归后序：左右子树分别找 p、q；
// 两边都找到 -> 当前节点即 LCA；只找到一边 -> 返回那一边的结果。
//
// 题目链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor 返回 p 和 q 的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p { // ✅ 递归基三合一：
		return root //    nil → nil（找不到）；root==p/q → root（找到了，且节点可以是自己的祖先）
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 左子树找
	right := lowestCommonAncestor(root.Right, p, q) // 右子树找

	if left != nil && right != nil { // ✅ p、q 分居两侧 → 我就是 LCA
		return root
	}

	if left != nil { // ✅ 只在左边找到 → 答案在左边（两目标都在左子树，或左边返回的已是 LCA）
		return left
	}

	return right // ✅ 只在右边（或都没找到返回 nil）
}

// 【V0 点评】
// 完全正确！这就是标准答案——LCA 递归的经典范式，一步没差：
//   ✅ 递归基巧妙三合一：nil / root==p / root==q 统一处理，
//      "节点可以是自己的祖先"直接由 root==p → return root 覆盖（示例2 的语义）
//   ✅ 后序自底向上：先在左右子树分别找，再合并结果——
//      "LCA 是第一个让 p、q 分居两侧的节点"，正是分治的聚合判断
//   ✅ 三分支合并逻辑：双侧命中→root；单侧命中→透传该侧；都空→nil
//   ✅ 时间 O(n) 每节点访问一次；空间 O(h) 递归栈
//   验证示例1 [3,5,1,...], p=5, q=1：5 在左子树、1 在右子树 → 根 3 双侧命中 → 3 ✓
//   验证示例2 p=5, q=4（4 在 5 的子树里）：左子树返回 5（先碰到 5 就返回，
//      不再深入找 4——"祖先找到了，后代必然在下面"），右侧 nil → 透传 5 ✓
//   🟡 隐藏细节值得说清（面试必被追问）：找到 p 后为什么不再深入找 q？
//      因为若 q 在 p 的子树里，p 自己就是 LCA（透传上去即可）；
//      若 q 不在 p 子树里，深入找也找不到——两种情况都不需要深入。
//      这个"提前返回"保证 O(n) 而非 O(n·h)
//   🟡 main 只测了 nil（老传统了），示例没验证（见下方 main）
// 无需修改——可直接提交 AC；这是自底向上分治的教科书级应用

// ──────────────────────────────────────────────────────────
// 【追踪示例 1：为什么"双侧命中→root"是对的】
//
//         3
//        / \
//       5   1        p=5, q=1
//      / \ / \
//     6  2 0  8
//       / \
//      7   4
//
//   LCA(3): 递归左右
//     LCA(5)：5==p → 直接返回 5（左子树结果 = 5）
//     LCA(1)：1==q → 直接返回 1（右子树结果 = 1）
//   3 看到 left=5, right=1 双侧非空 → p、q 分居两侧 → 返回 3 ✅
//
// 【追踪示例 2：为什么"单侧命中→透传"是对的】
//
//   p=5, q=4（4 在 5 的左子树... 不，4 是 2 的右孩子，5 的孙子）
//
//   LCA(3):
//     LCA(5)：5==p → 返回 5 ← 注意：没有继续下钻找 4！
//       （若 4 在 5 下面：5 就是 LCA；若不在：找也白找——两种情况都正确）
//     LCA(1)：没找到 → nil
//   3 看到 left=5, right=nil → 答案全在左边 → 透传 5 ✅
//
// 信息流动方向：p/q 的存在性从叶向根上传（自底向上），
// 第一个"双侧都有"的节点就是分叉点 = LCA——与 p0110 的分治范式同源。
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【深入辨析：为什么"率先命中一个就可以直接返回"？】
//
// 常见理解（有偏差）："先找到一个，就能确定另一个一定是它的子孙，所以返回。"
// 用示例 1 反证"确定"不成立：
//
//         3
//        / \
//       5    1        p=5, q=1
//      / \  / \
//     6  2 0   8
//
//   LCA(5)：5 == p → 立刻返回 5
//   此刻 q=1 在根 3 的右子树——【不在】5 的下面！
//   无法"确定"q 是 5 的子孙，但返回 5 依然正确（上层 LCA(1) 返回 1，
//   在 3 处双侧命中，LCA=3）。
//
// 正确的推理是"两 case 都不用深入"，而不是"确定在下面"：
//
//   | q 的位置            | 谁是 LCA          | 提前返回 p 为什么对           |
//   |---------------------|-------------------|------------------------------|
//   | ① q 在 p 的子树里    | p 自己（祖先可含自身） | 返回的就是答案本身           |
//   | ② q 在 p 子树之外    | 更高层的祖先节点    | 本子树向上"汇报 p"，LCA 由上层合成 |
//
//   关键：不需要知道 q 在哪。
//   深入 p 的子树去确认：case ① 确认完也只能得出"返回 p"——与提前返回相同；
//   case ② 深入了也找不到 q——白跑。两种情况深入都无收益——这才是完整的理由。
//
// 有趣的细分：这个"确定是子孙"的推理在【根层面】恰好成立——
//   若单侧命中发生在整棵树的根（left=x, right=nil）：
//     题目保证 p、q 都存在 + 右子树返回 nil（q 不在右边）
//     + 左侧一路下来所有旁支都是 nil（q 不在别处）
//     → 排除法：q 只能在 x 的子树里 ✓
//   但作为"任意时刻命中就返回"的普适理由不成立——中间层命中时
//   q 可能还在本子树之外（等上层去发现）。
//
// 一句话修正版：
//   不是"确定了另一个在下面所以返回"，而是
//   "在下面 → 我就是答案；不在下面 → 上层会合成答案——
//    两种情况提前返回都对，所以不需要深入"。
// ──────────────────────────────────────────────────────────

// lowestCommonAncestorV1 【LeetCode 标准答案】（与 V0 相同思路的紧凑写法，展示另一种合并风格）
func lowestCommonAncestorV1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorV1(root.Left, p, q)
	right := lowestCommonAncestorV1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil { // ✅ 两个 if 合并：left 空 → 必返回 right（无论是否 nil）
		return right
	}
	return left
}

// 【V1 点评 - 紧凑版】
// ✅ 与 V0 逻辑完全等价：V0 用"left!=nil 优先返回 left"，V1 用"left==nil 返回 right"——
//    只是合并了两个 if，少 4 行
// ✅ 语义：返回值三态——nil（没找到）/ p 或 q（找到一个）/ LCA 节点（都找到）
// 🟡 三态语义是这题最"绕"的点：返回的节点到底是"找到的目标"还是"LCA"？
//    对上层而言不需要区分——上层只关心 nil / 非 nil 和双侧是否都非空。
//    想通这一点，这题就没有玄学可言

// lowestCommonAncestorV2 存父指针法：先记录每个节点的父节点，p 向上走标记路径，q 向上走碰到标记即答案
// 时间 O(n)，空间 O(n)——思路更"人类"，也适用于多次查询的变体
func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	parent := map[*TreeNode]*TreeNode{} // 节点 → 父节点
	visited := map[*TreeNode]bool{}     // p 走过的祖先链

	var dfs func(node, par *TreeNode)
	dfs = func(node, par *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = par // ✅ 前序：带着父亲往下走（自上而下传参！）
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	// p 沿父指针一路向上，标记所有祖先（含自己）
	for cur := p; cur != nil; cur = parent[cur] {
		visited[cur] = true
	}
	// q 向上走，第一个碰到标记的就是 LCA
	for cur := q; cur != nil; cur = parent[cur] {
		if visited[cur] {
			return cur
		}
	}
	return nil
}

// 【V2 点评 - 存父指针法】
// ✅ 思路直白：树上找 LCA = 找两条到根路径的第一个交点（从交点往上路径重合）
// ✅ 注意 dfs(node, par) 是【自上而下】传参（父指针往下带）——
//    同一道题里 V0 用自底向上（聚合）、V2 用自顶向下（传父指针），
//    印证 p0226 的结论：两种范式解决同一问题的实例
// ✅ 时间 O(n) 三次遍历；空间 O(n) 两个 map
// 🟡 面试定位：V0 是首选答案；V2 在"树上带父指针/多次查询 LCA"的变体里更有用
//    （多次查询可先 O(n) 建父指针 map，之后每次查询 O(h)——类似 p0230 V2 的思路）

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间   | 空间   | 面试期望      |
// |-------------------------|--------|--------|---------------|
// | 后序分治（V0）           | O(n)   | O(h)   | ✅ 面试首选   |
// | 紧凑版（V1）             | O(n)   | O(h)   | ✅ 同为标准   |
// | 父指针法（V2）           | O(n)   | O(n)   | ⚠️ 变体备手   |
//
// 【核心易错点】
// 1. 递归基三合一：nil / ==p / ==q——"节点可以是自己的祖先"由此免费覆盖
// 2. 找到 p 后不再深入：若 q 在 p 子树内，p 即 LCA；否则深入也白找——O(n) 的保证
// 3. 返回值三态：nil / 目标节点 / LCA——上层只需区分 nil 和非 nil，不必区分后两者
// 4. 双侧命中 → 当前节点是 LCA；单侧命中 → 透传（答案全在那边）
// 5. 前序思维做不了这题：必须先知道子树里有没有 p/q（信息来自下方 → 自底向上）
//
// 【面试讲解建议】
// 1. 先说定义："LCA = 第一个让 p、q 分居两侧的节点"
// 2. 讲递归："后序，左右子树各找一遍，双侧非空返回我，单侧透传"
// 3. 主动讲清"找到就不深入"的正确性（必被追问）
// 4. 复杂度：时间 O(n)，空间 O(h)
// 5. 加分备手：父指针法（路径交点思路），多次查询变体用
// ──────────────────────────────────────────────────────────

func main() {
	// 示例树：[3,5,1,6,2,0,8,null,null,7,4]
	build := func() *TreeNode {
		node5 := &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
		node1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
		return &TreeNode{Val: 3, Left: node5, Right: node1}
	}
	root := build()
	node5 := root.Left
	node1 := root.Right
	node4 := root.Left.Right.Right // 4 是 5 的孙子

	// 示例1：p=5, q=1 → 3（分居两侧，根是 LCA）
	fmt.Println(lowestCommonAncestor(root, node5, node1).Val) // 3
	// 示例2：p=5, q=4 → 5（4 在 5 的子树里，5 是自己的祖先）
	fmt.Println(lowestCommonAncestor(root, node5, node4).Val) // 5
	// 深层 LCA：p=6, q=4 → 5（6 和 4 分居 5 的两侧）
	fmt.Println(lowestCommonAncestor(root, node5.Left, node4).Val) // 5
	// p=7, q=4 → 2（7、4 分居 2 的两侧；node7 是 2 的左孩子）
	node7 := root.Left.Right.Left
	fmt.Println(lowestCommonAncestor(root, node7, node4).Val) // 2

	// V1/V2 与 V0 结果一致
	fmt.Println(lowestCommonAncestorV1(root, node5, node1).Val) // 3
	fmt.Println(lowestCommonAncestorV1(root, node5, node4).Val) // 5
	fmt.Println(lowestCommonAncestorV2(root, node5, node1).Val) // 3
	fmt.Println(lowestCommonAncestorV2(root, node7, node4).Val) // 2
}
