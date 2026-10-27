//go:build p0041

// LeetCode 第104题 - 二叉树的最大深度 (Maximum Depth of Binary Tree)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、广度优先搜索
//
// 【题目描述】
// 给定一个二叉树 root，返回其最大深度。
// 二叉树的最大深度是指从根节点到最远叶子节点的最长路径上的节点数。
//
// 【示例 1】
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
// 【示例 2】
// 输入：root = [1,null,2]
// 输出：2
//
// 【数据范围】
// 树中节点数目在范围 [0, 10^4] 内
// -100 <= Node.val <= 100
//
// 【思路提示】递归：max(左子树深度, 右子树深度) + 1。最基础的树形 DFS，务必秒写。
//
// 题目链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 返回二叉树最大深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	// 到达所有的叶子结点，并记录深度
	// ✅ 思路正确：前序遍历携带当前深度，走到 nil（叶子的下一层）时更新全局最大值
	var helper func(root *TreeNode, curDepth int)
	helper = func(root *TreeNode, curDepth int) {
		if root == nil {
			if curDepth > depth { // ✅ 到达空节点 = 一条路径走完，更新最大深度
				depth = curDepth
			}
			return
		}

		helper(root.Left, curDepth+1)  // ✅ 深度 +1 下传
		helper(root.Right, curDepth+1)
	}

	helper(root, 0)

	return depth
}

// 【V0 点评】
// 完全正确！自顶向下 DFS（前序 + 全局变量）是标准解法之一：
//   ✅ 递归基：root == nil 时用 curDepth 更新最大值——nil 是"叶子下方的空位"，
//      传入的 curDepth 恰好等于该路径的深度（根从 0 计数，走 k 步到 nil 时 curDepth=k）
//   ✅ 自顶向下传参：curDepth+1 下传，不需要回溯时"减回来"（没有全局状态需要撤销）
//   ✅ 两个分支都递归：所有路径都会被走到，最大值必然被更新到
//   ✅ 时间 O(n) 每个节点访问一次；空间 O(h) 递归栈
//   ✅ 外层 root==nil 返回 0 与 helper 的 nil 处理一致（空树深度 0）
//   🟡 可简化：本题不需要 helper + 全局变量——"分治"一行搞定（见 V1）：
//      max(左子树深度, 右子树深度) + 1，代码量减半且无全局状态
//   🟡 main 测试太弱：只调用了 maxDepth(nil)，建议补示例树/单边树（见下方 main）
//   🟡 面试视角：V0 的"遍历带回"写法在路径类题（如路径总和、直径）有用武之地，
//      但求深度这类"聚合"题，分治写法更短更不易错

// maxDepthV1 【LeetCode 标准答案】分治递归，时间 O(n)，空间 O(h)
func maxDepthV1(root *TreeNode) int {
	if root == nil {
		return 0 // ✅ 空树深度 0（递归基）
	}
	left := maxDepthV1(root.Left)   // 左子树深度
	right := maxDepthV1(root.Right) // 右子树深度
	if left > right {
		return left + 1
	}
	return right + 1 // ✅ max(左, 右) + 1
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 分治思想：整棵树深度 = max(两棵子树深度) + 1（+1 是根自己这一层）
// ✅ 无 helper、无全局变量、无回溯——纯函数式，最不易错
// ✅ 3 行核心逻辑；Go 1.21+ 可直接写 max(left, right) + 1（内置 max 函数）
// ✅ 时间 O(n)，空间 O(h)（递归栈，最坏链状 O(n)，平衡 O(log n)）
// 追踪 root=[3,9,20,null,null,15,7]：
//   maxDepth(9)=1（叶子）、maxDepth(15)=maxDepth(7)=1
//   maxDepth(20)=max(1,1)+1=2
//   maxDepth(3)=max(1,2)+1=3 ✅
// 与 V0 对比：V0 是"自顶向下"（把深度传下去，走到头记录）；V1 是"自底向上"（子树答案汇总上来）
// 两种范式都要会：求深度/直径/平衡 → 分治；路径枚举/路径和 → 遍历带回

// maxDepthV2 BFS 层序遍历计数，时间 O(n)，空间 O(w)（w 为最大宽度）
func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		n := len(queue) // 当前层节点数
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++ // 处理完一层，深度 +1
	}
	return depth
}

// 【V2 点评 - BFS 层序计数】
// ✅ 思路：BFS 每处理完一层 depth++，处理了几层深度就是几
// ✅ 与 p0102 的分层 BFS 同款模板（记 n=len(queue)，弹 n 个，子节点入队）
// ✅ 只入队非 nil 子节点；时间 O(n)，空间 O(w)
// 求深度一般用 DFS（V1）；BFS 版的价值在于"层序天然带深度信息"，
// 最小深度、右视图、每层最大值等题用这个模板更顺手

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间   | 空间   | 面试期望      |
// |-------------------------|--------|--------|---------------|
// | 遍历带回（V0）           | O(n)   | O(h)   | ✅ 正确       |
// | 分治递归（V1 标准答案）  | O(n)   | O(h)   | ✅ 面试首选   |
// | BFS 层序计数（V2）       | O(n)   | O(w)   | ⚠️ 备选       |
//
// 【核心易错点】
// 1. 深度定义：节点数不是边数——[1,2] 深度是 2 不是 1；空树深度是 0
// 2. 递归基返回 0（空子树不贡献深度），+1 放在 max 之后（根贡献 1 层）
// 3. 两种范式区分：
//    自顶向下（V0）：把 curDepth 传下去，适合"路径枚举/记录沿途状态"
//    自底向上（V1）：先递归拿子树答案再合并，适合"聚合统计（深度/直径/平衡）"
// 4. 单边树是最深递归（O(n) 栈），10^4 节点不会栈溢出，但要能说出空间复杂度
//
// 【面试讲解建议】
// 1. 先说分治："深度 = max(左子树, 右子树) + 1"，一行核心
// 2. 递归基："nil 返回 0——空树不贡献深度"
// 3. 被追问时提 BFS："层序遍历数层数也行"
// 4. 复杂度：时间 O(n)，空间 O(h)，并说明链状树退化 O(n)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例树：[3,9,20,null,null,15,7] → 3
	tree := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	// 单边树：[1,null,2,null,3] → 3（最深递归形态）
	chain := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	// 单节点 → 1
	single := &TreeNode{Val: 1}

	fmt.Println(maxDepth(nil))    // 0
	fmt.Println(maxDepth(tree))   // 3
	fmt.Println(maxDepth(chain))  // 3
	fmt.Println(maxDepth(single)) // 1

	// V0/V1/V2 三版结果完全一致
	fmt.Println(maxDepthV1(tree), maxDepthV2(tree))   // 3 3
	fmt.Println(maxDepthV1(chain), maxDepthV2(chain)) // 3 3
	fmt.Println(maxDepthV1(single), maxDepthV2(single)) // 1 1
}
