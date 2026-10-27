//go:build p0041

// LeetCode 第617题 - 合并二叉树 (Merge Two Binary Trees)
//
// 【难度】简单 | 【面试频率】★★★ | 【标签】树、深度优先搜索、广度优先搜索、二叉树
//
// 【题目描述】
// 给你两棵二叉树：root1 和 root2。想象一下，当你将其中一棵覆盖到另一棵之上时，
// 两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。
// 合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；
// 否则不为 null 的节点将直接作为新二叉树的节点。
//
// 【示例 1】
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
//
// 【示例 2】
// 输入：root1 = [1], root2 = [1,2]
// 输出：[2,2]
//
// 【数据范围】
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10^4 <= Node.val <= 10^4
//
// 【思路提示】递归同步遍历：都非空则相加；一方为空直接返回另一方。
//
// 题目链接：https://leetcode.cn/problems/merge-two-binary-trees/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mergeTrees 合并两棵二叉树。
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var dfs func(root1 *TreeNode, root2 *TreeNode)
	dfs = func(root1 *TreeNode, root2 *TreeNode) {
		if root1 == nil && root2 == nil { // ✅ 都空：无需处理
			return
		}

		if root1 == nil { // 🔴 BUG：root1 为空时新建了局部节点，但【没有挂到树上】！
			root1 = &TreeNode{ //    root1 是【值传递】的参数，赋值只改了局部变量，
				Val: root2.Val, //    调用方的 root1.Left/Right 仍然是 nil
			} //    → 树 1 缺失的子树永远补不进结果
		} else if root2 == nil { // 🟡 同款问题：新建 root2 局部节点也没挂回去
			root2 = &TreeNode{ //    （碰巧"不挂也行"——见点评分析，但语义是错的）
				Val: root1.Val,
			}
		} else {
			root1.Val = root1.Val + root2.Val // ✅ 都非空：相加
			root2.Val = root2.Val + root1.Val // 🔴 BUG：此时 root1.Val 已更新，
			//    root2.Val = root2旧值 + root1新值 = root2旧 + (root1旧+root2旧) —— 脏值！
			//    反例：root1.Val=1, root2.Val=2 → root1.Val=3, root2.Val=5（应为 3）
		}

		dfs(root1.Left, root2.Left)
		dfs(root1.Right, root2.Right)
	}

	return root1 // 🔴 BUG：两树都空时返回 nil ✓；但 root1==nil 而 root2 非空时，
	//    dfs 里新建的根（局部）没被返回——整棵树 2 丢失，返回 nil ❌
	//    反例：mergeTrees(nil, [1,2]) 应返回 [1,2]，实际返回 nil
}

// 【V0 点评】
// 编译通过，但三个致命 bug，示例 1 就错：
//   🔴 BUG 1（新建节点没挂树）：root1 == nil 时 new 一个节点赋给【值传递的参数】，
//      调用方的 child 指针纹丝不动——树 1 缺失的整棵子树进不了结果
//      反例（示例 2）：root1=[1], root2=[1,2]：递归左孩子时 root1.Left=nil，
//      新建节点丢失 → 输出 [2] ❌ 应为 [2,2]
//   🔴 BUG 2（相加顺序）：root2.Val 用了已更新的 root1.Val —— 算出脏值
//      1+2 → root1=3, root2=5。虽然最终 return root1 不读 root2.Val，
//      但只要交换两行顺序或删掉这行才安全；当前写法靠"没人读 root2.Val"侥幸
//   🔴 BUG 3（返回值丢树 2）：root1==nil 且 root2 非空时返回 nil
//      反例：mergeTrees(nil, [1,2]) → nil ❌ 应返回 [1,2] 整棵
//   根源：想"原地修改树 1"但 Go 参数是值传递，nil 参数上 new 节点无法反映到调用方；
//   同时"双树都改"的对称写法引入了顺序依赖。
//   正解思路：让递归【返回】合并后的节点，父层接住挂上（自底向上组装）——见 V1

// mergeTreesV1 【LeetCode 标准答案】递归返回新树（自底向上组装），时间 O(min(m,n))，空间 O(h)
func mergeTreesV1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2 // ✅ 一方为空：返回另一方（整棵子树直接用，不用复制）
	}
	if root2 == nil {
		return root1
	}
	// 都非空：合并节点（复用 root1 作为新节点）
	root1.Val += root2.Val
	root1.Left = mergeTreesV1(root1.Left, root2.Left)    // ✅ 接住递归返回值挂回去
	root1.Right = mergeTreesV1(root1.Right, root2.Right) //    ——解决"新建节点没挂树"的根本手法
	return root1
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 与 p0100（相同树）同款"同步递归"骨架，但方向相反：
//    p0100 同步走是为了【比较】（返回 bool），本题同步走是为了【合并】（返回节点）
// ✅ 递归基是精髓："一方空 → 返回另一方"——缺失部分整棵接上，一个 if 解决 V0 的
//    BUG 1（没挂树）和 BUG 3（丢树 2）：挂接由"父层接住返回值"完成
// ✅ 复用 root1 节点（原地改）：省内存；想不动原树就 new 节点（见 V2）
// ✅ 时间 O(min(m,n))：短树走完，长树剩余部分整体接上不再遍历；空间 O(h)
// 追踪示例2 root1=[1], root2=[1,2]：
//   merge(1,1)：Val=2；左 = merge(nil,2) → 返回 2 挂上；右 = merge(nil,nil) → nil
//   输出 [2,2] ✅

// mergeTreesV2 不修改原树的版本：全部新建节点（保持输入树完整）
func mergeTreesV2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2 // 注意：这里仍直接返回 root2（共享子树）；要完全隔离可复制
	}
	if root2 == nil {
		return root1
	}
	node := &TreeNode{Val: root1.Val + root2.Val} // ✅ 新建节点，不动原树
	node.Left = mergeTreesV2(root1.Left, root2.Left)
	node.Right = mergeTreesV2(root1.Right, root2.Right)
	return node
}

// 【V2 点评 - 新建节点版】
// ✅ 与 V1 唯一区别：重叠处 new 节点而不是改 root1——原树保持完整
// ✅ 适用场景：调用方还要继续用原树（函数式风格/并发安全）
// 🟡 严格隔离需连"单方为空的子树"也深拷贝，否则与原树共享节点——
//    一般题目不要求，口头提一句即可
// 面试定位：先写 V1（标准）；被问"能不能不动原树"时切换 V2 并说明共享语义

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间        | 空间   | 面试期望      |
// |-----------------------------|-------------|--------|---------------|
// | 原地改+参数挂树（V0 有 bug） | O(n)        | O(h)   | ❌ 值传递坑   |
// | 返回值挂接（V1 标准答案）    | O(min(m,n)) | O(h)   | ✅ 面试首选   |
// | 新建节点（V2）               | O(min(m,n)) | O(h)   | ✅ 追问备手   |
//
// 【核心易错点】
// 1. Go 值传递：在 nil 参数上 new 节点无法反映到调用方——必须由父层接住【返回值】挂接
//    （这是链表/树题的高频坑：p0002 尾指针、p0021 合并链表同款思想）
// 2. 对称双写有顺序依赖：root1.Val += root2.Val 后再 root2.Val += root1.Val 是脏值
// 3. 递归基"一方空返回另一方"一举解决挂接和丢树两个问题
// 4. root1==nil && root2==nil 的双重判断不需要——第一个 if 已经兜住（V1 只有两个 if）
//
// 【面试讲解建议】
// 1. 先说递归基："一方空返回另一方——缺失子树整体接上"
// 2. 讲挂接："父层接住递归返回值挂回去，Go 参数值传递改不了孩子指针"
// 3. 主动提版本选择："原地改省内存；不动原树就新建节点"
// 4. 复杂度：时间 O(min(m,n))，空间 O(h)
// ──────────────────────────────────────────────────────────

func main() {
	printTree := func(root *TreeNode) {
		if root == nil {
			fmt.Println("[]")
			return
		}
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			n := len(queue)
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				fmt.Printf("%d ", cur.Val)
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
			}
			fmt.Print("| ")
		}
		fmt.Println()
	}

	// 示例1：[1,3,2,5] + [2,1,3,null,4,null,7] → [3,4,5,5,4,null,7]
	r1a := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	r1b := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	printTree(mergeTreesV1(r1a, r1b)) // 3 | 4 5 | 5 4 7 |

	// 示例2：[1] + [1,2] → [2,2]
	r2a := &TreeNode{Val: 1}
	r2b := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	printTree(mergeTreesV1(r2a, r2b)) // 2 | 2 |

	// V0 的反例：root1=nil, root2=[1,2] → V0 返回 nil ❌，V1 返回整棵树 2
	printTree(mergeTreesV1(nil, r2b)) // 2 | 2 |

	// V2 与 V1 结果一致（不动原树）
	r3a := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	r3b := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	printTree(mergeTreesV2(r3a, r3b)) // 3 | 4 5 | 5 4 7 |
}
