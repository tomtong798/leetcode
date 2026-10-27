//go:build p0041

// LeetCode 第98题 - 验证二叉搜索树 (Validate Binary Search Tree)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、二叉搜索树
//
// 【题目描述】
// 给你一个二叉树的根节点 root，判断其是否是一个有效的二叉搜索树。
// 有效二叉搜索树定义如下：
//   - 节点的左子树只包含小于当前节点的数。
//   - 节点的右子树只包含大于当前节点的数。
//   - 所有左子树和右子树自身必须也是二叉搜索树。
//
// 【示例 1】
// 输入：root = [2,1,3]
// 输出：true
//
// 【示例 2】
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5，但右子节点的值是 4。
//
// 【数据范围】
// 树中节点数目范围 [1, 10^4] 内
// -2^31 <= Node.val <= 2^31 - 1
//
// 【思路提示】两种经典做法：
// 1) 递归传递上下界 (min, max)，节点值必须严格落在开区间内；
// 2) 中序遍历结果必须严格递增，只需比较前一个值。
//
// 题目链接：https://leetcode.cn/problems/validate-binary-search-tree/
package main

import (
	"fmt"
	"math"
)

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST 判断是否为合法 BST。
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false // 🔴 BUG 0：空树是合法 BST，应返回 true（LeetCode 保证非空，不影响判题）
	}
	// TODO: 实现
	if root.Left == nil && root.Right == nil {
		return true // ✅ 叶子节点是合法 BST
	}

	if root.Left == nil || root.Right == nil { // 🔴 BUG 1：只有一边为空就判 false
		return false //    单边子树完全可以是合法 BST！
	} //    反例：    1        ← 合法 BST，却返回 false
	//              \
	//               2
	//    反例：    2        ← 合法 BST，却返回 false
	//              /
	//             1

	if root.Left.Val < root.Val && root.Right.Val > root.Val { // 🔴 BUG 2：检查完直接 return true，
		return true //    没有递归！第 63 行的递归永远执行不到（死代码）
	} else { //    反例：  5        左右子都满足 3<5<8，
		return false //          / \       但 2 在右子树中（应 >5），不是 BST，却返回 true
	} //               3   8
	//                  /
	//                 2

	return isValidBST(root.Left) && isValidBST(root.Right) // 🔴 BUG 3：即使执行到这，
	//    只比较父子也不够：孙子可能违反祖先的约束
	//    反例：  5        3<5 ✓ 3<6 ✓ 6<10 ✓ 但 3 应 >5（在 5 的右子树里），不是 BST
	//           / \
	//          3  10
	//             /
	//            6  ← 没问题；换成下面这棵：
	//          反例：  5
	//                 / \
	//                3  10
	//                  /
	//                 1  ← 1 < 5 但在右子树，不是 BST；只比父子发现不了
}

// 【V0 点评】
// 编译通过，但三处逻辑错误，main 的 3 个用例只对 1 个：
//   🔴 BUG 1（第 53-55 行）：单边子树返回 false——[1,null,2] 这样的合法 BST 被误判
//   🔴 BUG 2（第 57-61 行）：比较完父子就 return true，没递归——多层树只检查了第一层
//      main 的 case3（1,[0,2]）恰好是两层满树且合法，掩盖了这个 bug
//   🔴 BUG 3（第 63 行）：死代码（前面已 return），且"只比父子"的思路本身不成立——
//      BST 约束是"左子树所有节点 < 根 < 右子树所有节点"，孙子要受祖先约束
//      反例：根 5，右子 10，右孙 1：1<10 ✓ 但 1 在 5 的右子树中应 >5，不是 BST
//   🔴 BUG 0：空树应返回 true（语义问题）
//   🟡 测试用例没覆盖：单边树、多层非法树（孙辈违反祖先约束）
//   ✅ 叶子返回 true 正确
//   ✅ 严格比较（< 和 >，不是 <= 和 >=）方向正确——BST 不允许重复值
// 根源：把 BST 定义理解成"父子比较"了。正确定义是"节点值必须在 (祖先链给定的) 开区间内"

// isValidBSTV1 【LeetCode 标准答案】递归传递上下界，时间 O(n)，空间 O(h)（h 为树高）
func isValidBSTV1(root *TreeNode) bool {
	// 用 int64 边界避开 int32 极值（节点值可达 ±2^31，用 int 的 MinInt/MaxInt 会挡不住边界值）
	var validate func(node *TreeNode, min, max int64) bool
	validate = func(node *TreeNode, min, max int64) bool {
		if node == nil {
			return true // ✅ 空子树是合法 BST（递归基）
		}
		v := int64(node.Val)
		if v <= min || v >= max { // ✅ 必须严格落在开区间 (min, max) 内
			return false
		}
		// ✅ 左子树收紧上界为当前值，右子树收紧下界为当前值
		return validate(node.Left, min, v) && validate(node.Right, v, max)
	}
	return validate(root, math.MinInt64, math.MaxInt64)
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 核心：每个节点的合法范围由祖先链决定——
//    进入左子树：上界收紧为父值（左子树所有节点 < 父）
//    进入右子树：下界收紧为父值（右子树所有节点 > 父）
//    这样孙辈自然受所有祖先约束，修复 V0 BUG 3
// ✅ 空子树返回 true（递归基），单边树正确处理，修复 BUG 0/1
// ✅ 每个节点恰好访问一次，时间 O(n)；递归栈深度 O(h)
// ✅ 用 int64 极值做初始边界：节点值可达 ±2^31，若用 int32 的 MinInt/MaxInt，
//    值恰好等于 ±2^31 的节点会被误判（v <= min 触发）
// 追踪反例 [5,3,10,null,null,null,1]（1 在 5 的右子树，非法）：
//   validate(5, -inf, +inf) ✓ → validate(3, -inf, 5) ✓（3 是叶子）
//   → validate(10, 5, +inf) ✓ → validate(1, 5, 10)：1 <= 5 → false ✅ 抓住了！

// isValidBSTV2 中序遍历法，时间 O(n)，空间 O(h)——利用"BST 中序严格递增"性质
func isValidBSTV2(root *TreeNode) bool {
	prev := int64(math.MinInt64) // 中序遍历的前一个值

	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) { // 左
			return false
		}
		if int64(node.Val) <= prev { // ✅ 中序序列必须严格递增（<= 即重复/降序，非法）
			return false
		}
		prev = int64(node.Val)     // 根：更新前驱值
		return inorder(node.Right) // 右
	}
	return inorder(root)
}

// 【V2 点评 - 中序遍历法】
// ✅ 性质：BST 的中序遍历严格递增——只需维护前驱 prev，逐个比较
// ✅ 不需要保存整个中序序列，O(h) 空间（递归栈），比"遍历存数组再检查"更优
// ✅ 同样修复所有反例：1 在 5 的右子树 → 中序 [3,1,5,10]，1 <= 3 之前……
//    实际中序是 [3,1,5,10]？不对，正确中序：3,1,5,10 中 1<5 但序列 3,1 已经降序 → false ✅
// 与 V1 对比：V1 更直观（区间约束）；V2 更巧妙（借助性质，代码更短）
// 两者都是标准答案，面试任选；讲 V2 时先说"BST 中序有序"这个性质

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间   | 空间   | 面试期望      |
// |-----------------------------|--------|--------|---------------|
// | 父子比较（V0 有 bug）        | O(n)   | O(h)   | ❌ 思路错误   |
// | 上下界递归（V1 标准答案）    | O(n)   | O(h)   | ✅ 面试首选   |
// | 中序遍历（V2）               | O(n)   | O(h)   | ✅ 同为标准   |
//
// 【核心易错点】
// 1. BST ≠ "每个节点满足父子关系"！约束是整棵子树：
//    左子树的所有节点（含孙子及更深）< 根 < 右子树的所有节点
// 2. 单边子树（一边为 nil）是完全合法的结构，别当成非法
// 3. 严格比较：不允许重复值，用 < 和 >，不是 <= 和 >=
// 4. 边界极值：节点值可达 ±2^31，初始上下界要用比 int 更宽的类型（int64），
//    否则值恰好等于极值的节点被误判
// 5. 测试反例集：单边树、[5,1,4,null,null,3,6]（右子违反）、[5,3,10,n,n,n,1]（孙辈违反）
//
// 【面试讲解建议】
// 1. 先点破易错点："只比较父子不够，孙辈要受祖先约束"
// 2. 讲 V1 思路："递归传递开区间 (min, max)，左子树收紧上界，右子树收紧下界"
// 3. 或讲 V2："BST 中序遍历严格递增，维护前驱比较即可"
// 4. 提到 ±2^31 边界用 int64，展示严谨性
// ──────────────────────────────────────────────────────────

func main() {
	case1 := &TreeNode{
		Val: 1,
	}
	case2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
	}
	case3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isValidBST(case1)) // true  ✅
	fmt.Println(isValidBST(case2)) // false ✅（V0 也判 false，但理由是"单边"，碰巧对）
	fmt.Println(isValidBST(case3)) // true  ✅ 两层满树，V0 只查了第一层碰巧对

	// V0 的反例：V0 全错，V1/V2 全对
	// 反例 1：单边合法树 [1,null,2]
	r1 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(isValidBST(r1), isValidBSTV1(r1), isValidBSTV2(r1)) // false true true
	// 反例 2：孙辈违反祖先约束 [5,1,6,null,null,null,7] 合法；
	// 非法的经典例 [5,1,4,null,null,3,6]：4 < 5 却在右子树
	r2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(r2), isValidBSTV1(r2), isValidBSTV2(r2)) // true false false
	//   ↑ V0 返回 true 是错的！只比了第一层 1<5<4? 不成立 → 其实返回 false，见下
	//   实际追踪：r2 第一层 1<5 ✓ 但 4>5 不成立 → V0 返回 false（碰巧对）
	// 真正暴露 BUG 2 的反例：第一层合法但深层非法 [5,3,10,null,null,7,null]
	//   （7 应 >5，在 10 的左子树中违反；V0 只查第一层 3<5<10 ✓ 返回 true ❌）
	r3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 7}}}
	fmt.Println(isValidBST(r3), isValidBSTV1(r3), isValidBSTV2(r3)) // true false false

	root := &TreeNode{Val: 50}
	// 左子树
	root.Left = &TreeNode{Val: 30}
	root.Left.Left = &TreeNode{Val: 20}
	root.Left.Left.Left = &TreeNode{Val: 10}
	root.Left.Right = &TreeNode{Val: 40}
	root.Left.Right.Left = &TreeNode{Val: 35}
	root.Left.Right.Left.Left = &TreeNode{Val: 33}
	// 右子树
	root.Right = &TreeNode{Val: 70}
	root.Right.Left = &TreeNode{Val: 60}
	root.Right.Left.Left = &TreeNode{Val: 55}
	root.Right.Left.Left.Right = &TreeNode{Val: 57}
	root.Right.Left.Right = &TreeNode{Val: 65}
	root.Right.Right = &TreeNode{Val: 80}
	root.Right.Right.Right = &TreeNode{Val: 90}

	fmt.Println(InorderTraversal(root))

}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var helper func(r *TreeNode)
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		helper(r.Left)
		result = append(result, r.Val)
		helper(r.Right)
	}

	helper(root)

	return result
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var helper func(r *TreeNode)
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		result = append(result, r.Val)
		helper(r.Left)
		helper(r.Right)
	}

	helper(root)

	return result
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var helper func(r *TreeNode)
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		helper(r.Left)
		helper(r.Right)
		result = append(result, r.Val)
	}

	helper(root)

	return result
}

func isValidBSTInorderTraversal(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(r *TreeNode, max int, min int) bool
	helper = func(r *TreeNode, max int, min int) bool {
		if r == nil {
			return true
		}

		if r.Val < min || r.Val > max {
			return false
		}

		return helper(r.Left, r.Val, min) && helper(r.Right, max, r.Val)

	}

	return helper(root, math.MaxInt64, math.MinInt64)
}
