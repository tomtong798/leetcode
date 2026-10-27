//go:build p0041

// LeetCode 第199题 - 二叉树的右视图 (Binary Tree Right Side View)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、广度优先搜索、二叉树
//
// 【题目描述】
// 给定一个二叉树的根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，
// 返回从右侧所能看到的节点值。
//
// 【示例 1】
// 输入：root = [1,2,3,null,5,null,4]
// 输出：[1,3,4]
//
// 【示例 2】
// 输入：root = [1,null,3]
// 输出：[1,3]
//
// 【数据范围】
// 二叉树的节点个数的范围是 [0,100]
// -100 <= Node.val <= 100
//
// 【思路提示】BFS 每层最后一个节点即为右视图；
// 或 DFS 根右左，按深度首访记录（每层只记第一个访问的节点）。
//
// 题目链接：https://leetcode.cn/problems/binary-tree-right-side-view/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView 返回二叉树右视图。
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var queue = []*TreeNode{root}
	var result []int

	for len(queue) > 0 {
		n := len(queue)
		result = append(result, queue[n-1].Val) // ✅ 巧：先记队尾再弹出——队尾就是本层最右节点
		//    （比常规"弹完一层再记最后一个"少一次循环外的判断，等价但更早收集）

		for n > 0 { // ✅ 分层模板：弹 n 个 = 本层，子节点入队 = 下一层
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil { // ✅ 只入队非 nil 子节点（p0102 教训已吸收）
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			n--
		}
	}

	return result
}

// 【V0 点评】
// 完全正确！这是 BFS 解法的标准写法，一次通过：
//   ✅ 核心洞察：每层的最右节点 = 队列中本层区间的最后一个（queue[n-1]）
//      收集时机放在内层循环【前】——此时队列还没动，前 n 个恰好是本层，
//      n-1 下标就是最右——比"边弹边记最后一个"更直接
//   ✅ 分层 BFS 模板完整：记 n → 弹 n 个 → 子节点入队（p0102 的标准套路）
//   ✅ 只入队非 nil 子节点：队列无 nil，n 就是真实层宽
//   ✅ 空树返回 nil 正确
//   ✅ 时间 O(n)，空间 O(w)（w 为最大层宽）
//   验证示例1 [1,2,3,null,5,null,4]：
//     第1层队列 [1] → 记 1；入 2,3
//     第2层队列 [2,3] → 记 3（最右）✓；2 入 5，3 入 4 → 队列 [5,4]
//     第3层队列 [5,4] → 记 4 ✓；无子节点
//     输出 [1,3,4] ✅
//   验证边界 [1,2]（看不到 2 吗？不——2 在左子树，但第 2 层只有它）：
//     第2层队列 [2] → 记 2 ✓ 输出 [1,2]（右视图取"每层最右"，不是"只有右子树"）
//   🟡 main 只测了 nil，示例用例没验证（见下方 main）
//   🟡 可选进阶：DFS 版（根右左 + 每层首访记录）代码更短（见 V1）；
//      以及"收集时机"的另一种等价写法：内层循环里 i==n-1 时收集
// 无需修改——本函数可直接提交 LeetCode 通过

// rightSideViewV1 【LeetCode 标准答案】DFS 根右左 + 深度首访，时间 O(n)，空间 O(h)
func rightSideViewV1(root *TreeNode) []int {
	var result []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) { // ✅ 本层第一次被访问——根右左顺序下第一个到的就是最右
			result = append(result, node.Val)
		}
		dfs(node.Right, depth+1) // ✅ 先右后左：保证每层最先访问的是最右节点
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return result
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 核心：DFS 按"根→右→左"访问，每层第一个访问到的节点就是该层最右的
//    depth == len(result) 判断"本层是否已记录"——与 p0104 V2 的分层技巧完全同源！
//    （p0104 层序 V2：depth==len(result) 时开新层数组；本题：depth==len(result) 时记录首访值）
// ✅ 先递归右子树：顺序即答案——不需要"找到再判断是否最右"，访问顺序天然保证
// ✅ 时间 O(n)，空间 O(h) 递归栈（深树比 BFS 的 O(w) 省空间）
// 追踪示例1 [1,2,3,null,5,null,4]：
//   dfs(1,0)：depth 0==len 0 → 记 1；先右
//     dfs(3,1)：depth 1==len 1 → 记 3；先右
//       dfs(4,2)：depth 2==len 2 → 记 4
//       dfs(nil,3) 返回；3 的左 nil 返回
//     dfs(2,1)：depth 1 != len 3（第 1 层已记过 3）→ 跳过
//       dfs(nil,2)；dfs(5,2)：depth 2 != len 3 → 已记过 4，跳过
//   输出 [1,3,4] ✅
// BFS vs DFS 选择：
//   BFS（V0）：直观（层序=视图），队列 O(w) 空间
//   DFS（V1）：代码更短，O(h) 空间，但"根右左+首访"的推理要多想一步
//   两者都是标准答案；变形单侧视图题（左视图）DFS 只需改访问顺序为根左右

// ──────────────────────────────────────────────────────────
// 【DFS/BFS vs 递归/迭代——概念辨析】
//
// 关键结论：DFS/BFS 是【遍历策略】，递归/迭代是【实现手段】，两个维度不能等价：
//
//   "深度优先搜索就是递归，广度优先搜索就是迭代"  ← 错误等价！
//
// 正确的对应关系：
//
//              递归实现        迭代实现
//   DFS      ✅ 常见直观     ✅ 用显式栈
//   BFS      ⚠️ 少见别扭     ✅ 常见直观
//
// 真正绑定的是【数据结构】：DFS ↔ 栈（递归栈或显式栈），BFS ↔ 队列。
// 递归只是"借用函数调用栈"的 DFS——所以 DFS 常用递归写，但不是必须。
// BFS 也能用递归写（把队列推进写成递归），只是别扭、无收益，没人这么写。
//
// 思维模型差异（为什么本题 DFS 不直观）：
//
//   DFS 沿一条路走到黑再回头——右视图的"跳过"发生在回溯途中，看不见
//   BFS 每轮面对完整的一层横截面——谁右谁左一目了然，所见即所得
//
// 右视图题的直观度差异根源：
//   BFS：定义"每层最右"与 BFS 的"层"天然同构，queue[n-1] 闭眼取
//   DFS：正确性靠两个隐性约定叠加——
//        ① "根右左"保证每层最先碰到的是最右节点
//        ② depth == len(result) 恰好表示"这层还没记过"
//     两个约定都不写在代码表面，要在脑中模拟几层才能相信它对
//
// 选择建议：
//   只求快速写对 → BFS（直观）
//   空间敏感（深树 O(h) vs 队列 O(w)）→ DFS
//   变形题：路径收集/每层最值/根到叶路径 → DFS 带深度参数更顺手
//   面试策略 → 先用最顺的 BFS 写对，再口头补"DFS 根右左+首访，空间 O(w)→O(h)"
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                      | 时间   | 空间   | 面试期望      |
// |---------------------------|--------|--------|---------------|
// | BFS 层序取队尾（V0）       | O(n)   | O(w)   | ✅ 标准答案   |
// | DFS 根右左首访（V1）       | O(n)   | O(h)   | ✅ 面试首选   |
//
// 【核心易错点】
// 1. 右视图 = 每层最右节点，不是"只有右子树上的节点"——
//    [1,2] 输出 [1,2]（第 2 层最右是左子树的 2）
// 2. BFS 收集时机：进入本层循环前，队列前 n 个就是本层，queue[n-1] 即最右
// 3. DFS 必须先右后左：先左后右就变成"左视图"
// 4. depth == len(result) 判"首访"——这个技巧与层序建层（p0102 V2/p0104 V2）同源
//
// 【面试讲解建议】
// 1. 先说 BFS："层序遍历，每层最后一个就是右视图"——最直观
// 2. 再讲 DFS："根右左 + 每层首访记录，先到先得"
// 3. 主动提空间："DFS 递归栈 O(h)，比 BFS 队列 O(w) 更省（深树场景）"
// 4. 复杂度：时间都是 O(n)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例1：[1,2,3,null,5,null,4] → [1,3,4]
	t1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}
	// 示例2：[1,null,3] → [1,3]
	t2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}
	// 易错边界：[1,2] → [1,2]（第 2 层最右是"左子树"的 2）
	t3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	// 完全树：[1,2,3] → [1,3]
	t4 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	fmt.Println(rightSideView(nil)) // []
	fmt.Println(rightSideView(t1))  // [1 3 4]
	fmt.Println(rightSideView(t2))  // [1 3]
	fmt.Println(rightSideView(t3))  // [1 2] ← 左子树也能进右视图
	fmt.Println(rightSideView(t4))  // [1 3]

	// V0/V1 结果完全一致
	fmt.Println(rightSideViewV1(t1)) // [1 3 4]
	fmt.Println(rightSideViewV1(t3)) // [1 2]
	fmt.Println(rightSideViewV1(t4)) // [1 3]
}
