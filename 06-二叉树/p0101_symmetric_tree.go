//go:build p0041

// LeetCode 第101题 - 对称二叉树 (Symmetric Tree)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】树、深度优先搜索、广度优先搜索
//
// 【题目描述】
// 给你一个二叉树的根节点 root，检查它是否轴对称。
//
// 【示例 1】
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 【示例 2】
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
// 【数据范围】
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
// 【思路提示】递归比较一对节点 (l, r)：值相等，且 l.Left 对称于 r.Right、
// l.Right 对称于 r.Left。注意与"相同树"题（p100）的区别。
//
// 题目链接：https://leetcode.cn/problems/symmetric-tree/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 判断二叉树是否轴对称。
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true // ✅ 空树对称
	}

	var helper func(left *TreeNode, right *TreeNode) bool
	helper = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil { // ✅ 都空：镜像位置结构相同
			return true
		}
		if left == nil || right == nil { // ✅ 一空：结构不对称
			return false
		}
		if left.Val != right.Val { // ✅ 值不等：不对称
			return false
		}
		return helper(left.Left, right.Right) && helper(left.Right, right.Left) // ✅ 关键：交叉递归！
	}

	return helper(root.Left, root.Right) // ✅ 根自身不用比，从左右子树开始
}

// 【V0 点评】
// 完全正确！这就是标准答案的递归写法：
//   ✅ 核心洞察正确：对称 = 左子树和右子树互为镜像，
//      而"镜像比较"的关键是交叉递归：left.Left 对 right.Right，left.Right 对 right.Left
//      （不是 p0100 相同树的平行比较 left.Left 对 right.Left！）
//   ✅ helper 三段判断（都空→一空→值不等）顺序正确，与 p0100 同款套路
//   ✅ 根节点跳过自身，直接比较 root.Left 和 root.Right，正确
//   ✅ && 短路剪枝：外层镜像已不对称时，内层不再比较
//   ✅ 时间 O(n)：每个节点恰好访问一次；空间 O(h) 递归栈
//   ✅ root == nil 返回 true（LeetCode 保证 n>=1，防御性处理无妨）
//   🟡 main 测试太弱：只调用了 isSymmetric(nil)，没有断言输出
//      建议补：对称树、不对称树（结构/值）、单节点、null 位置不对称（见下方 main）
//   🟡 可选进阶：非递归用队列成对入队（见 V2），面试追问的备手
// 无需修改——本函数可直接提交 LeetCode 通过

// isSymmetricV1 【LeetCode 标准答案】递归紧凑版（合并空判断，展示与 p0100 的对称美）
func isSymmetricV1(root *TreeNode) bool {
	var mirror func(l, r *TreeNode) bool
	mirror = func(l, r *TreeNode) bool {
		if l == nil || r == nil {
			return l == r // 一箭双雕：都空 → true；一空一非空 → false
		}
		return l.Val == r.Val && //          值相等
			mirror(l.Left, r.Right) && //    外侧镜像：l 的左 vs r 的右
			mirror(l.Right, r.Left) //       内侧镜像：l 的右 vs r 的左
	}
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 与 V0 逻辑完全等价，两处精简：
//    1) 前两个 if 合并成 return l == r（p0100 同款技巧）
//    2) 值比较并入 && 链：l.Val == r.Val && mirror(外侧) && mirror(内侧)，一条表达式写完
// ✅ mirror 的语义：l 和 r 互为镜像吗？——轴对称的定义被翻译得非常干净
// ✅ 时间 O(n)，空间 O(h)
// 对比 p0100 相同树：
//    相同树：  same(l, r)  = l.Val==r.Val && same(l.Left, r.Left)  && same(l.Right, r.Right)  // 平行
//    对称树：mirror(l, r)  = l.Val==r.Val && mirror(l.Left, r.Right) && mirror(l.Right, r.Left) // 交叉
//    一字之差，两题互为镜像题——一起记效率最高

// isSymmetricV2 BFS 迭代版：队列成对入队，时间 O(n)，空间 O(w)（w 为最大宽度）
func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := [][2]*TreeNode{{root.Left, root.Right}} // 队列存"镜像位置对"

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		l, r := pair[0], pair[1]

		if l == nil || r == nil {
			if l != r { // 一空一非空 → 不对称
				return false
			}
			continue // 都空，处理下一对
		}
		if l.Val != r.Val {
			return false
		}
		// 交叉入队：l.Left 与 r.Right 是镜像对，l.Right 与 r.Left 是镜像对
		queue = append(queue, [2]*TreeNode{l.Left, r.Right}, [2]*TreeNode{l.Right, r.Left})
	}
	return true
}

// 【V2 点评 - BFS 迭代版】
// ✅ 队列存"镜像位置对"，弹出比较，子节点交叉成对入队（注意不是平行入队！）
// ✅ 与 p0100 的 BFS 版结构相同，仅"配对方式"从平行改成交叉
// ✅ 无递归栈溢出风险；空间 O(w)
// 面试通常写递归版；追问"不用递归"时写这个，强调交叉入队即可

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                  | 时间   | 空间   | 面试期望      |
// |-----------------------|--------|--------|---------------|
// | DFS 递归（V0）         | O(n)   | O(h)   | ✅ 标准答案   |
// | DFS 紧凑版（V1）       | O(n)   | O(h)   | ✅ 面试首选   |
// | BFS 队列迭代（V2）     | O(n)   | O(w)   | ⚠️ 追问备手   |
//
// 【核心易错点】
// 1. 最大陷阱：对称是"交叉比较"，不是相同树的"平行比较"
//    平行：left.Left 对 right.Left（p0100 相同树）
//    交叉：left.Left 对 right.Right（本题）——写错就变成判"相同树"
// 2. 根节点不参与比较：对称性由 root.Left 与 root.Right 的镜像关系决定
// 3. 空节点也是结构：[1,2,2,null,3,null,3] 中 null 的位置不对称 → false
// 4. 记忆法：照镜子时抬起左手，镜中人是抬右手——左对右、右对左
//
// 【面试讲解建议】
// 1. 先说定义："对称 = 左右子树互为镜像"
// 2. 讲递归："mirror(l,r)：值相等，且 l 左对 r 右、l 右对 r 左（交叉）"
// 3. 主动对比 p0100："相同树是平行比较，对称树是交叉比较，一字之差"
// 4. 复杂度：时间 O(n)，空间 O(h)；追问非递归时讲队列交叉入队
// ──────────────────────────────────────────────────────────

func main() {
	// 对称树：[1,2,2,3,4,4,3] → true
	t1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	// 不对称（null 位置）：[1,2,2,null,3,null,3] → false
	t2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	// 不对称（值）：[1,2,2,2,null,2,null]... 简单点：[1,2,2] 值不同 → false
	t3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	// 单节点 → true
	t4 := &TreeNode{Val: 1}

	fmt.Println(isSymmetric(t1)) // true
	fmt.Println(isSymmetric(t2)) // false
	fmt.Println(isSymmetric(t3)) // false
	fmt.Println(isSymmetric(t4)) // true
	fmt.Println(isSymmetric(nil)) // true

	// V0/V1/V2 三版结果完全一致
	fmt.Println(isSymmetricV1(t1), isSymmetricV2(t1)) // true true
	fmt.Println(isSymmetricV1(t2), isSymmetricV2(t2)) // false false
	fmt.Println(isSymmetricV1(t3), isSymmetricV2(t3)) // false false
}
