//go:build p0041

// LeetCode 第110题 - 平衡二叉树 (Balanced Binary Tree)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】树、深度优先搜索、二叉树
//
// 【题目描述】
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。
//
// 【示例 1】
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
// 【示例 2】
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
// 【数据范围】
// 树中的节点数在范围 [0, 5000] 内
// -10^4 <= Node.val <= 10^4
//
// 【思路提示】自底向上递归：返回 -1 表示子树已失衡，否则返回子树高度；
// 当前节点左右子树高度差 > 1 则返回 -1 剪枝。
//
// 题目链接：https://leetcode.cn/problems/balanced-binary-tree/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 判断二叉树是否平衡。
func isBalanced(root *TreeNode) bool {
	// TODO: 实现

	var height func(root *TreeNode, curHeight int) (int, bool)
	height = func(root *TreeNode, curHeight int) (int, bool) {
		if root == nil {
			return curHeight, true // ✅ 递归基：空子树高度 = 传入深度（p0104 同款语义），平衡
		}

		leftHeight, leftBalance := height(root.Left, curHeight+1)
		rightHeight, rightBalance := height(root.Right, curHeight+1)

		if leftHeight >= rightHeight { // ✅ 返回较大者作为子树高度
			return leftHeight, leftHeight-rightHeight <= 1 && rightHeight-leftHeight <= 1 && leftBalance && rightBalance // 🟡 高度差判断冗余：leftHeight >= rightHeight 时差值非负，
		} //    leftHeight-rightHeight <= 1 一个条件就够（rightHeight-leftHeight 必 <= 0）

		return rightHeight, leftHeight-rightHeight <= 1 && rightHeight-leftHeight <= 1 && leftBalance && rightBalance // 🟡 同上：这里只需 rightHeight-leftHeight <= 1
	}

	_, isBalanc := height(root, 0)

	return isBalanc // ✅ 正确！聚合 bool 的"自底向上分治"是标准解法之一
}

// isBalancedV3 按你的思路优化：保留 (int, bool) 双返回值，修掉 V0 的三处冗余
// ——去掉 curHeight 参数、去掉冗余高度差判断、补上失衡短路剪枝
func isBalancedV3(root *TreeNode) bool {
	// 返回 (子树高度, 是否平衡)。高度自底向上计数：nil 返回 0，父层 = max(左右)+1
	var height func(node *TreeNode) (int, bool)
	height = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true // ✅ 递归基：空子树高度 0（不需要外部传入深度）
		}

		leftHeight, leftBalance := height(node.Left)
		if !leftBalance { // ✅ 剪枝：左子树失衡，右子树不用算
			return 0, false
		}
		rightHeight, rightBalance := height(node.Right)
		if !rightBalance { // ✅ 剪枝：右子树失衡
			return 0, false
		}

		diff := leftHeight - rightHeight // ✅ 一次算差值，双向判断各写一次，不再重复
		if diff > 1 || diff < -1 {
			return 0, false // 本层失衡（返回的高度无所谓，bool 已定）
		}

		if leftHeight > rightHeight { // 返回较大者 + 1 = 子树高度
			return leftHeight + 1, true
		}
		return rightHeight + 1, true
	}

	_, balanced := height(root)
	return balanced
}

// 【V3 点评 - 按你的思路优化版】
// 完整保留 V0 的设计（双返回值聚合 + 自底向上分治），只修三处：
//   1️⃣ 去掉 curHeight 参数：V0 是"自顶向下计数（nil 返回传入深度）"混合式，
//      纯自底向上（nil 返回 0，父层 max+1）少一个参数，递归基更简单
//   2️⃣ 高度差判断去冗余：V0 在两个分支里各写了双向判断（4 个比较条件），
//      且 if(leftHeight>=rightHeight) 已保证方向——实际上一个条件就永真；
//      V3 提出 diff = leftHeight - rightHeight，> 1 || < -1 两个条件干净利落
//   3️⃣ 补剪枝：V0 左子树失衡后仍会递归右子树；V3 提前 return，失衡子树只访问一次
// 对比 V1（-1 哨兵）：V3 用 (int, bool) 双返回值表达同样的语义——
//   V1 把"失衡"编码进高度（-1），V3 用独立的 bool——两者等价，
//   V1 更精炼（单返回值），V3 语义更显式（读代码不用记哨兵约定），
//   面试任选，讲清楚"一次后序遍历 + 剪枝"就是满分
// 时间 O(n)，空间 O(h)——与 V1 完全同阶

// ──────────────────────────────────────────────────────────
// 【如何判断自顶向下 vs 自底向上？】
//
// 判断标准只有一个：看"信息"（参数/返回值）的流动方向，不看函数名：
//
//   自顶向下：信息作为【参数】往下传，在【递归前】就知道
//   自底向上：信息作为【返回值】往上传，在【递归后】才知道
//
// 结构特征速查表：
//
//   | 特征         | 自顶向下                | 自底向上                    |
//   |--------------|-------------------------|-----------------------------|
//   | 关键信息在   | 参数里                  | 返回值里                    |
//   | 计算时机     | 递归前                  | 递归后（合并子结果）        |
//   | 常见配件     | 全局变量记录最优        | 分治三段式：基→子结果→合并 |
//   | 典型语式     | dfs(node, 沿途状态)     | return f(左子结果, 右子结果)|
//   | 别名         | 前序思维                | 后序思维                    |
//
// 本目录中的实例：
//   自顶向下——p0104 V0（curDepth 参数下传）、p0098 V1（min/max 区间下传）
//   自底向上——p0104 V1（子树深度上传）、p0105（前序游标+中序区间）、本题 V0/V1/V3
//
// O(n²) 警报信号：代码里"调用了另一个递归函数获取子树信息"——
//   leftHeight := height(root.Left)  ← 另一次独立递归，重复计算（V2/TopDown 版）
//   left := height(node.Left)        ← 自身递归，结果直接复用（V1/V3，O(n)）
//   两行长得几乎一样，区别在 height 是"外部函数"还是"本次递归自身"！
//
// 选择经验：
//   路径枚举/路径和/沿途状态 → 自顶向下（参数传状态）
//   聚合统计（深度/直径/平衡/子树和/是否BST） → 自底向上（返回值带答案）
//   两者都能做时 → 自底向上通常更简洁（无全局变量）
// ──────────────────────────────────────────────────────────

// 【V0 点评】
// 完全正确！这是"自底向上分治"的标准写法（一次遍历 O(n)），我验证了：
//   ✅ 示例1 [3,9,20,null,null,15,7]：根左右高差 0，20 左右高差 0 → true ✓
//   ✅ 示例2 [1,2,2,3,3,null,null,4,4]：失衡点在根——
//      左子树 2 高 3（3→4），右子树 2' 高 1，差 2 → false ✓
//   ✅ 链状树 [1,null,2,null,3]：每层差 0 → true ✓（链是平衡的！）
//   ✅ 非平衡链 [1,2,null,3,null,4]... 手工验证：根左高 2 右高 0 → 差 2 → false ✓
//   ✅ 递归基正确：nil 返回 (curHeight, true)——空子树平衡，高度=传入深度
//   ✅ bool 聚合：左右都平衡 && 本层差 <=1，语义完整
//   ✅ 时间 O(n) 每个节点访问一次；空间 O(h)——没有"上层再调 maxDepth"的 O(n²) 问题
//   🟡 高度差条件冗余：两处都写了双向判断（a-b<=1 && b-a<=1），
//      但 if 已保证比较方向，单个条件即可；更简洁的写法是 abs 差值（Go 1.21+ 内置 abs）或交换后统一判断
//   🟡 curHeight 参数可以省掉：返回"子树高度"（自底向上计数）不需要外部传入——
//      nil 返回 0，父层 = max(子)+1（见 V1）。你的写法是"自顶向下计数+自底向上聚合"混合，
//      正确但多一个参数；纯自底向上更简洁
//   🟡 剪枝缺失：左子树已失衡时仍会递归右子树——可在 leftBalance==false 时提前返回（见 V1 的 -1 约定）
//   🟡 main 只测了 nil，建议补两个示例用例（见下方 main）
// 无需修改即可 AC；V1 是面试更精炼的写法

// isBalancedV1 【LeetCode 标准答案】自底向上 + -1 哨兵剪枝，时间 O(n)，空间 O(h)
func isBalancedV1(root *TreeNode) bool {
	// 返回子树高度；失衡时返回 -1（哨兵），父层看到 -1 直接短路
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0 // ✅ 递归基：空子树高度 0
		}
		left := height(node.Left)
		if left == -1 {
			return -1 // ✅ 剪枝：左子树失衡，无需再看右子树
		}
		right := height(node.Right)
		if right == -1 {
			return -1 // ✅ 剪枝：右子树失衡
		}
		diff := left - right
		if diff > 1 || diff < -1 {
			return -1 // ✅ 本层失衡：高度差 > 1
		}
		if left > right { // 返回较大者 + 1 = 子树高度
			return left + 1
		}
		return right + 1
	}
	return height(root) != -1
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 精髓：用 -1 既当"失衡信号"又实现剪枝——
//    父层收到 -1 立即向上传播，不再计算另一棵子树，整棵树失衡部分只访问一次
// ✅ 纯自底向上计数：nil 返回 0，父层 = max(左右) + 1——不需要 curHeight 参数
// ✅ 单一返回值 (int) 比你的 (int, bool) 双返回值简洁
// ✅ 时间 O(n)，空间 O(h)
// 追踪示例2 [1,2,2,3,3,null,null,4,4]（失衡点：根的左子树高 3 vs 右子树高 1）：
//   height(4)=height(4')=1（叶子）
//   height(3)=max(1,1)+1=2（孩子是两个 4）
//   height(3')=max(0,0)+1=1（叶子）
//   height(2)=左 height(3)=2、右 height(3')=1 → diff=1 ✓ 平衡，高度 3
//   height(2')=1（叶子）
//   height(1)：左 height(2)=3、右 height(2')=1 → diff=2 > 1 → 返回 -1
//   根的 height = -1 → false ✅
// 与 V0 对比：逻辑等价，V1 用哨兵值把"bool 聚合 + 剪枝"合二为一，是本题最经典的写法

// isBalancedV2 反面教材：逐节点调 maxDepth 的 O(n²) 写法
// 结果正确，但面试会扣分——用于对比理解"为什么 V1 是 O(n)"
func isBalancedV2(root *TreeNode) bool {
	// maxDepth 就是 p0104 的解法：自底向上算子树高度，O(该子树节点数)
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		if left > right {
			return left + 1
		}
		return right + 1
	}

	// 前序遍历每个节点，检查它的左右子树高度差
	var check func(node *TreeNode) bool
	check = func(node *TreeNode) bool {
		if node == nil {
			return true // 空子树平衡
		}
		leftH := maxDepth(node.Left)  // ← 罪魁：每次都完整重算子树高度
		rightH := maxDepth(node.Right) // ← 同上
		diff := leftH - rightH
		if diff > 1 || diff < -1 {
			return false // 本节点失衡
		}
		return check(node.Left) && check(node.Right) // 继续检查子树
	}

	return check(root)
}

// 【V2 点评 - O(n²) 反面教材】
// ✅ 结果正确：所有用例和 V0/V1 一致——但过程大量重复计算
// 🔴 为什么是 O(n²)：
//    每个节点 x 都调用 maxDepth(x.Left/x.Right)，该调用遍历 x 的整棵子树
//    同一棵子树的高度被重复算了多遍：
//      深度为 1 的节点被所有祖先各算一次 → 深度 d 的节点被计算约 d 次
//    总代价 ≈ Σ(每个节点的深度) —— 满树时 ≈ n·h = O(n·log n)，
//    链状树（h=n）时退化为 Σd = n²/2 = O(n²) ❌
// 🔴 与 V1 的本质区别：
//    V2 自顶向下检查：父节点先要子树高度 → 重复向下算
//    V1 自底向上聚合：高度在下层算好一次，随返回值带上来复用
//    ——这就是"自顶向下 vs 自底向上"在效率上的典型分野
// 🟡 面试定位：如果只能写出 V2，面试官通常会给提示"能不能一次遍历解决"，
//    能当场改写成 V1 才算过关——所以 V2 最好的用法是当"对比脚手架"
// 追踪 V2 在链 [1,null,2,null,3,null,4] 上的浪费：
//    check(1)：maxDepth(右链)=4、maxDepth(nil)=0 → 差 4 失衡，返回 false
//    好在此处提前 false 短路了；若树平衡，check 会递归到每个节点，
//    check(2) 又要 maxDepth(3→4 的链)... 深层节点被反复遍历

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                          | 时间   | 空间   | 面试期望      |
// |-------------------------------|--------|--------|---------------|
// | 逐层调 maxDepth（V2）         | O(n²)  | O(n)   | ❌ 面试扣分   |
// | 双返回值分治（V0）             | O(n)   | O(h)   | ✅ 正确       |
// | -1 哨兵剪枝（V1 标准答案）     | O(n)   | O(h)   | ✅ 面试首选   |
// | 双返回值+剪枝（V3，V0 优化版） | O(n)   | O(h)   | ✅ 同为标准   |
//
// 【核心易错点】
// 1. O(n²) 陷阱（见 V2）：每个节点都调 maxDepth（前序+重复计算高度）——链状树退化 O(n²)
//    正解：一次后序遍历，自底向上把高度和平衡性一起算上来
// 2. 链状树是平衡的（每层高差 0/1）——别直觉认为"细长"就不平衡
// 3. 失衡定义是"每个节点"的左右子树高差 <=1，不只看根——任何深层失衡都算
// 4. 剪枝实现：V1 用 -1 哨兵短路；V3 用 bool 提前 return——语义等价
//
// 【面试讲解建议】
// 1. 先说陷阱："逐节点调 maxDepth 是 O(n²)，标准解法一次后序遍历 O(n)"
// 2. 讲哨兵："高度和 -1 复用同一个返回值，失衡向上短路，天然剪枝"
// 3. 递归基："nil 返回 0——空子树高度 0 且平衡"
// 4. 复杂度：时间 O(n)，空间 O(h)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例1：平衡 [3,9,20,null,null,15,7] → true
	t1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	// 示例2：不平衡 [1,2,2,3,3,null,null,4,4] → false
	t2 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2}}
	// 链状树：每层高差 0/1 → true（细长但平衡）
	t3 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}

	fmt.Println(isBalanced(nil))  // true
	fmt.Println(isBalanced(t1))   // true
	fmt.Println(isBalanced(t2))   // false
	fmt.Println(isBalanced(t3))   // true

	// V0/V1/V2/V3 四版结果完全一致（V2 正确但慢 O(n²)）
	fmt.Println(isBalancedV1(t1), isBalancedV1(t2), isBalancedV1(t3)) // true false true
	fmt.Println(isBalancedV2(t1), isBalancedV2(t2), isBalancedV2(t3)) // true false true
	fmt.Println(isBalancedV3(t1), isBalancedV3(t2), isBalancedV3(t3)) // true false true
}
