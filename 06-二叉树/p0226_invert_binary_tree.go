//go:build p0041

// LeetCode 第226题 - 翻转二叉树 (Invert Binary Tree)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】树、深度优先搜索、广度优先搜索、递归
//
// 【题目描述】
// 给你一棵二叉树的根节点 root，翻转这棵二叉树，并返回其根节点。
//
// 【示例 1】
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//
// 【示例 2】
// 输入：root = [2,1,3]
// 输出：[2,3,1]
//
// 【示例 3】
// 输入：root = []
// 输出：[]
//
// 【数据范围】
// 树中节点数目范围 [0, 100]
// -100 <= Node.val <= 100
//
// 【思路提示】递归：交换左右子树指针，再分别递归翻转左右子树。经典"Homebrew 作者面试趣题"。
//
// 题目链接：https://leetcode.cn/problems/invert-binary-tree/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 翻转二叉树并返回根节点。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // ✅ 递归基：空树返回 nil
	}

	root.Left, root.Right = root.Right, root.Left // ✅ Go 多重赋值一行交换，无需临时变量

	invertTree(root.Left)  // ✅ 递归翻转（交换后的）左右子树
	invertTree(root.Right) // ✅ 注意：这里翻转的是"新"的左右——但因为整棵都要翻，顺序无所谓

	return root
}

// 【V0 点评】
// 完全正确！这就是标准答案，一步没差：
//   ✅ 递归基 nil → nil
//   ✅ Go 惯用交换：a, b = b, a 多重赋值，比临时变量 tmp 干净
//   ✅ 先交换再递归（或先递归再交换也对）：翻转 = 每个节点左右互换，
//      交换和递归的先后不影响结果——两种顺序都遍历全树、每个节点恰好交换一次
//   ✅ 时间 O(n) 每节点访问一次；空间 O(h) 递归栈
//   ✅ 返回 root：题目要求返回根节点（其实树是原地修改的，返回值只是为了链式调用方便）
//   🟡 递归返回值被丢弃（第 48-49 行）——功能正确，但更地道的写法是接住返回值：
//      root.Left = invertTree(root.Left)（见 V1），好处是和"分治"范式统一
//   🟡 main 只测了 nil（老习惯了），建议补示例 1 验证（见下方 main）
// 无需修改——可直接提交 AC

// invertTreeV1 【LeetCode 标准答案】分治写法（利用返回值），时间 O(n)，空间 O(h)
func invertTreeV1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTreeV1(root.Left)   // 翻转左子树，拿到新左子树
	right := invertTreeV1(root.Right) // 翻转右子树，拿到新右子树
	root.Left = right                 // 组装：新左 = 旧右（已翻转）
	root.Right = left
	return root
}

// 【V1 点评 - 分治写法】
// ✅ 与 V0 逻辑等价，风格不同：V0"先交换再递归"（遍历思维），
//    V1"先递归拿翻转结果再组装"（分治思维：自底向上）
// ✅ 接住返回值的好处：函数签名 *TreeNode 的意义完整——"返回翻转好的子树"
// ✅ 时间 O(n)，空间 O(h)
// 两种都是标准答案；V0 的多重赋值在 Go 里更地道，V1 的分治结构与
// p0105 构造二叉树（先递归构造左右再接上根）同一范式

// ──────────────────────────────────────────────────────────
// 【本题 = 两种递归范式的最小完整样本】
//
// 翻转二叉树问题足够简单，却是展示"自上而下 vs 自底向上"的最佳教具——
// 同一个问题，两种写法结构差异全部呈现：
//
//   V0 自上而下（top-down）：先处理自己，把问题抛给下一层
//     root.Left, root.Right = root.Right, root.Left  // ① 我先交换
//     invertTree(root.Left)                           // ② 再递归子树
//     invertTree(root.Right)
//
//   V1 自底向上（bottom-up）：先等子树算好，再组装我的答案
//     left := invertTreeV1(root.Left)   // ① 先递归拿翻转好的子树
//     right := invertTreeV1(root.Right)
//     root.Left = right                 // ② 再组装
//     root.Right = left
//
// 结构对照（对应 p0110 的判断标准：看信息流动方向）：
//
//   | 维度          | V0 自上而下           | V1 自底向上             |
//   |---------------|-----------------------|-------------------------|
//   | 处理顺序      | 先自己后子树          | 先子树后自己            |
//   | 遍历序        | 前序（根左右）        | 后序（左右根）          |
//   | 返回值        | 可丢弃（原地改）      | 必须接住（组装用）      |
//   | 信息方向      | 指令往下传"去翻转"    | 答案往上传"翻好的子树"  |
//   | 典型配件      | 全局/外部的累计变量   | return 合并子结果       |
//
// 本题的特殊性：翻转是"每个节点独立的局部操作"，两种范式结果完全相同——
// 所以才适合当教具：不用担心正确性差异，专心看结构差异。
//
// 但在别的题里两种范式不可互换：
//   信息由祖先决定 → 只能自上而下（p0098 BST：合法区间由祖先链给出）
//   答案由子树聚合 → 只能自底向上（p0104 深度：max(子树)+1，
//                    p0110 平衡：高度和平衡性从叶往上算）
// 两种范式都要会：见题先问"信息从哪来？"——祖先给的往下传，子树给的往上传。
// ──────────────────────────────────────────────────────────

// invertTreeV2 BFS 迭代版，时间 O(n)，空间 O(w)
func invertTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cur.Left, cur.Right = cur.Right, cur.Left // 出队时交换
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return root
}

// 【V2 点评 - BFS 迭代版】
// ✅ 层序遍历，每个节点出队时交换左右——不用分层计数（本题不关心层结构）
// ✅ 只入队非 nil 子节点；时间 O(n)，空间 O(w)
// 价值：被追问"不用递归怎么写"时的答案；也印证 p0199 的结论——DFS/BFS 都能配迭代

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                  | 时间   | 空间   | 面试期望      |
// |-----------------------|--------|--------|---------------|
// | 递归交换（V0）         | O(n)   | O(h)   | ✅ 面试首选   |
// | 分治组装（V1）         | O(n)   | O(h)   | ✅ 同为标准   |
// | BFS 迭代（V2）         | O(n)   | O(w)   | ⚠️ 追问备手   |
//
// 【核心易错点】
// 1. 本题几乎没有坑——递归基 + 交换即可；最常见"错误"是过度思考
// 2. 交换与递归的先后顺序不影响正确性（整棵树都要翻）
// 3. Go 多重赋值 a,b = b,a 交换指针，无需 tmp
// 4. 趣闻：Homebrew 作者 Max Howell 在 Google 面试挂在这题上（"翻转二叉树都写不出"，
//    他发推吐槽）——说明再资深的工程师也得会基本功
//
// 【面试讲解建议】
// 1. 一句话思路："每个节点左右互换，递归处理所有节点"
// 2. 写递归版（3 行核心），提一句"先交换或先递归都对"
// 3. 被追问迭代版：BFS 队列，出队时交换
// 4. 复杂度：时间 O(n)，空间 O(h)/O(w)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例1：[4,2,7,1,3,6,9] → [4,7,2,9,6,3,1]
	t1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}

	fmt.Println(invertTree(nil)) // <nil>

	printLevel(invertTree(t1)) // [4] [7 2] [9 6 3 1] —— 已翻转

	// 三版结果一致（重新建树分别测）
	t2 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	printLevel(invertTreeV1(t2)) // [4] [7 2] [9 6 3 1]

	t3 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	printLevel(invertTreeV2(t3)) // [2] [3 1]
}

// printLevel 层序打印，肉眼验证翻转结果
func printLevel(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		fmt.Print(level, " ")
	}
	fmt.Println()
}
