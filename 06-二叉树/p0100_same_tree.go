//go:build p0041

// LeetCode 第100题 - 相同的树 (Same Tree)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】树、深度优先搜索、广度优先搜索、二叉树
//
// 【题目描述】
// 给你两棵二叉树的根节点 p 和 q，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 【示例 1】
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
//
// 【示例 2】
// 输入：p = [1,2], q = [1,null,2]
// 输出：false
//
// 【数据范围】
// 两棵树上的节点数目都在范围 [0, 100] 内
// -10^4 <= Node.val <= 10^4
//
// 【思路提示】双递归同步比较：都空则相同，一空则不同，值不等则不同，否则递归比较左右子树。
//
// 题目链接：https://leetcode.cn/problems/same-tree/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree 判断两棵树是否相同。
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // ✅ 都空：结构相同
		return true
	}
	if p == nil || q == nil { // ✅ 一空一非空：结构不同
		return false
	}
	if p.Val != q.Val { // ✅ 值不同
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // ✅ 同步递归左右子树
}

// 【V0 点评】
// 完全正确！这就是标准答案的递归写法：
//   ✅ 三段判断顺序无懈可击：都空 → 一空 → 值不等 → 递归
//      （先判"都空"再判"一空"的顺序不能反，否则 p==nil&&q==nil 会先命中 || 的短路）
//   ✅ && 短路：左子树不同时不再比较右子树，天然剪枝
//   ✅ 时间 O(min(m,n))：节点少的那棵树遍历完即停；空间 O(h) 递归栈
//   ✅ 与思路提示"双递归同步比较"完全一致
//   🟡 main 测试太弱：只调用了 isSameTree(nil,nil)，没有任何断言输出
//      建议补：相同树、结构不同、值不同、单边树等用例（见下方 main）
//   🟡 可选进阶：非递归写法用队列做 BFS 同步层序比较（见 V2），面试可作为追问的备手
// 无需修改——本函数可直接提交 LeetCode 通过

// isSameTreeV1 【LeetCode 标准答案】递归版（与 V0 等价，合并前两个 if 展示紧凑写法）
func isSameTreeV1(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil { // ✅ 一箭双雕：任一为空时，返回"两个是否都为空"
		return p == q //    都空 → true；一空一非空 → false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTreeV1(p.Left, q.Left) && isSameTreeV1(p.Right, q.Right)
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 与 V0 逻辑完全等价，只是把前两个 if 合并成 return p == q：
//    p 为空或 q 为空时，p == q 当且仅当两者都空——空指针比较是安全的
// ✅ 6 行写完，是本题最简洁的标准写法
// ✅ 时间 O(min(m,n))，空间 O(h)
// V0 的三段式更利于讲解（意图逐条展开）；V1 的合并式更简洁——面试两种都行

// isSameTreeV2 BFS 迭代版：队列同步层序比较，时间 O(min(m,n))，空间 O(w)（w 为最大宽度）
func isSameTreeV2(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}} // ✅ 队列元素是节点对 (p节点, q节点)

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		a, b := pair[0], pair[1]

		if a == nil || b == nil {
			if a != b { // 一空一非空 → 结构不同
				return false
			}
			continue // 都空，继续处理下一对
		}
		if a.Val != b.Val {
			return false
		}
		// 子节点成对入队（nil 也入队，由上面的判断统一处理）
		queue = append(queue, [2]*TreeNode{a.Left, b.Left}, [2]*TreeNode{a.Right, b.Right})
	}
	return true
}

// 【V2 点评 - BFS 迭代版】
// ✅ 用队列做同步层序遍历：每弹出一对节点，比较值，再把左右子节点成对入队
// ✅ 优点：无递归栈溢出风险（本题 n≤100 无所谓，但海量数据时迭代更稳）
// ✅ 空间 O(w)：w 是最大层宽，最坏 O(n)；与递归版 O(h) 量级相当
// 🟡 细节：nil 子节点也要入队——空结构差异（一边有子树一边没有）正是在弹出时被发现的
// 面试通常写递归版即可；若面试官追问"能不能不用递归"，就写这个队列版

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间        | 空间   | 面试期望      |
// |-------------------------|-------------|--------|---------------|
// | DFS 递归（V0）           | O(min(m,n)) | O(h)   | ✅ 标准答案   |
// | DFS 递归紧凑版（V1）     | O(min(m,n)) | O(h)   | ✅ 面试首选   |
// | BFS 队列迭代（V2）       | O(min(m,n)) | O(w)   | ⚠️ 追问备手   |
//
// 【核心易错点】
// 1. 判断顺序：必须先判"都空"（&&）再判"一空"（||），反了会被短路逻辑坑
//    （或像 V1 用 return p == q 合并，天然无顺序问题）
// 2. 结构不同 ≠ 值不同：[1,2] vs [1,null,2] 结构就不同，nil 的位置就是结构
// 3. && 短路是天然剪枝：左子树已不同时，右子树不用看
// 4. 复杂度是 O(min(m,n)) 不是 O(m+n)：短树遍历完，长树剩余部分无需访问
//
// 【面试讲解建议】
// 1. 先说递归三部曲："都空 true → 一空 false → 值不等 false → 递归左右"
// 2. 提一句合并技巧："前两步可合并成 return p == q"
// 3. 被追问非递归时，讲队列成对入队的 BFS
// 4. 复杂度：时间 O(min(m,n))，空间 O(h)
// ──────────────────────────────────────────────────────────

func main() {
	// 相同的树：p=[1,2,3], q=[1,2,3] → true
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	// 结构不同：p=[1,2], q=[1,null,2] → false（2 一个在左一个在右）
	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	// 值不同：p=[1,2], q=[1,3] → false
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}
	// 一空一非空 → false
	p4 := &TreeNode{Val: 1}

	fmt.Println(isSameTree(nil, nil))  // true（V0 原有的调用，补上输出）
	fmt.Println(isSameTree(p1, q1))    // true
	fmt.Println(isSameTree(p2, q2))    // false
	fmt.Println(isSameTree(p3, q3))    // false
	fmt.Println(isSameTree(p4, nil))   // false

	// V0/V1/V2 三版结果完全一致
	fmt.Println(isSameTreeV1(p1, q1), isSameTreeV2(p1, q1)) // true true
	fmt.Println(isSameTreeV1(p2, q2), isSameTreeV2(p2, q2)) // false false
	fmt.Println(isSameTreeV1(p4, nil), isSameTreeV2(p4, nil)) // false false
}
