//go:build p0041

// LeetCode 第543题 - 二叉树的直径 (Diameter of Binary Tree)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、二叉树
//
// 【题目描述】
// 给你一棵二叉树的根节点，返回该树的直径长度。
// 直径长度是指任意两个节点之间路径长度的最大值（边数）。
// 这条路径可能经过也可能不经过根节点。
// 两节点之间路径的长度由它们之间边数表示。
//
// 【示例 1】
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：路径 [4,2,1,3] 或 [5,2,1,3] 的长度为 3。
//
// 【示例 2】
// 输入：root = [1,2]
// 输出：1
//
// 【数据范围】
// 树中节点数目在范围 [1, 10^4] 内
// -100 <= Node.val <= 100
//
// 【思路提示】递归求每个节点的"深度"，同时用 左深+右深 更新全局最大直径。
// 答案在递归过程中维护，函数返回值是深度、全局变量是直径——树形 DP 的入门套路。
//
// 题目链接：https://leetcode.cn/problems/diameter-of-binary-tree/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// diameterOfBinaryTree 返回二叉树直径（最长路径边数）。
func diameterOfBinaryTree(root *TreeNode) int {
	// 最长直径一定是两个叶子结点之间的距离
	// 两个叶子结点有一个共同祖先
	// 左子树的最大高度+右子树的最大高度
	// 🟡 注释第一行不准确：最长路径端点未必是叶子——
	//    单边树 [1,2,null,3]：路径 3→2→1 长 2，端点是根 1（非叶子）
	//    准确表述："路径的最高点"是某个节点 x，直径 = max(x左深 + x右深)

	var dfs func(root *TreeNode) int
	var maxDiameter int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0 // ✅ 递归基：空子树高度 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		// 以 root 为根结点的左右子树的最长直径
		// 🟡 注释表述不准：这行是"以 root 为【最高点】的路径长度 = 左深+右深"，
		//    用它更新全局直径——不是"左右子树的最长直径"

		if maxDiameter < left+right { // ✅ 核心一步：经过 root 的最长路径 = left+right（边数）
			maxDiameter = left + right
		}

		if left > right { // ✅ 返回值是高度（供父层拼接路径用），不是直径
			return left + 1
		}

		return right + 1
	}

	dfs(root)

	return maxDiameter // ✅ 答案在全局变量：树形 DP 的标准结构
}

// 【V0 点评】
// 完全正确！一次通过，这是树形 DP 的教科书实现：
//   ✅ 双职责设计正确（本题的灵魂）：返回值=子树高度（给父层用），
//      全局变量=最大直径（最终答案）——"深度上传、直径全局记"缺一不可
//   ✅ left+right 恰好是边数：深度按"节点数"计，路径 4→2→1→3 的边数
//      = 深度(4侧)=2 + 深度(3侧)=1 = 3 ✓（不用 +1/-1 修正，直接对）
//   ✅ "路径可能不经过根"处理对了：每个节点都用 left+right 更新，
//      不只看根——最高点在哪一层都会被扫到
//   ✅ 时间 O(n)，空间 O(h)
//   验证示例1 [1,2,3,4,5]：节点 2 处 left=1(4)、right=1(5) → 2；
//    节点 1 处 left=2(2子树)、right=1(3) → 3 ✅ 最大
//   验证示例2 [1,2]：根处 1+0=1 ✅
//   验证单边链 [1,2,null,3]：每层 left+right 依次 1,1 → 直径 2 ✅（端点是根，非叶子）
//   验证空树：dfs(nil)=0，maxDiameter=0 ✅（题目保证 n>=1，防御性正确）
//   🟡 两处注释表述可打磨（见行内标注）："叶子端点"和"左右子树直径"都不准确，
//      准确心智模型是"路径的最高点 x，直径 = max(x.left深度 + x.right深度)"
//   🟡 max 函数可简化 if（Go 1.21+ 有内置 max），纯风格问题
//   🟡 main 只测了 nil（老传统），示例没验证（见下方 main）
// 无需修改——可直接提交 AC

// ──────────────────────────────────────────────────────────
// 【心智模型：为什么"每个节点的 左深+右深"取 max 就是对的？】
//
// 任何一条路径，形状都是"∧"：先一直向上到某个最高点 x，再一直向下：
//
//        x   ← 路径最高点（路径在这里"拐弯"）
//       / \
//      ↑   ↓  上行段 = x 左子树的一条深链（左深度条边）
//     ↑     ↓ 下行段 = x 右子树的一条深链（右深度条边）
//
//   路径边数 = 左子树深度 + 右子树深度（以 x 为根计）
//
// 关键观察：
//   1. 每条路径【唯一确定】一个最高点 x——所以枚举所有 x 取 max 就覆盖所有路径
//   2. x 的左右两段要尽量长 → 恰好是左右子树的【深度】
//   3. 路径不必经过树根：x 可以是任何节点——所以递归里每个节点都更新
//   4. 端点未必是叶子：单边树的最长路径，一端就是根（注释第一行的坑）
//      （把叶子换掉为"路径端点"，模型就严密了）
//
// 与 p0104 的关系：本题 = p0104 求深度 + 一行"用左右深度更新全局答案"。
// 这就是【树形 DP】入门：递归返回"子树信息"（高度），顺手用信息更新答案，
// 与 p0110（返回高度+检查平衡）、p0236（返回找到与否+双侧合成）同一范式。
// ──────────────────────────────────────────────────────────

// diameterOfBinaryTreeV1 【LeetCode 标准答案】（与 V0 等价，用内置 max 的简洁版）
func diameterOfBinaryTreeV1(root *TreeNode) int {
	maxDiameter := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		maxDiameter = max(maxDiameter, left+right) // 经过 node 的最长路径
		return max(left, right) + 1                // 返回高度
	}

	depth(root)
	return maxDiameter
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 与 V0 逻辑完全等价：只是用内置 max（Go 1.21+）替代两个 if，核心 8 行
// ✅ max(maxDiameter, left+right) 一行同时完成"比较+更新"，意图更醒目
// ✅ 双职责注释化：返回值=高度，maxDiameter=直径
// 面试先写 V0 风格（if 更显式）或 V1 风格（max 更简洁）都对，
// 讲清楚"返回高度 / 全局记直径"的双职责就是满分

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间   | 空间   | 面试期望      |
// |-----------------------------|--------|--------|---------------|
// | 树形DP：高度上传（V0）       | O(n)   | O(h)   | ✅ 面试首选   |
// | max 简洁版（V1）             | O(n)   | O(h)   | ✅ 同为标准   |
// | 逐节点调 depth（前序）       | O(n²)  | O(n)   | ❌ 同 p0110 陷阱 |
//
// 【核心易错点】
// 1. 双职责缺一不可：只返回高度不更新直径 → 答案错；只算直径不返回高度 → 父层拼不了路径
// 2. 边数不用修正：深度按节点数计，left+right 直接就是边数（画图验证）
// 3. 路径未必经过根、端点未必是叶子：每个节点都要更新（单边树反例）
// 4. O(n²) 陷阱与 p0110 同款：逐节点调 depth 重复计算——树形 DP 一次遍历解决
// 5. 直径的"路径"是简单路径（不重复经过节点），∧ 模型天然满足
//
// 【面试讲解建议】
// 1. 先讲心智模型："每条路径都是 ∧ 形，最高点为 x，长度 = x左右深度之和"
// 2. 讲双职责："递归返回高度供父层拼接，全局变量记最大直径"
// 3. 主动提单边树反例："端点未必是叶子、路径未必过根——所以每个节点都更新"
// 4. 复杂度：时间 O(n)，空间 O(h)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例1：[1,2,3,4,5] → 3（路径 4→2→1→3）
	t1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3}}
	// 示例2：[1,2] → 1
	t2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	// 单边链：[1,2,null,3] → 2（端点是根，非叶子——注释坑的反例）
	t3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	// 直径不过根的用例：根只有左子树，最长路径的最高点在节点 2
	// 路径 6→4→2→5（长 3），不经过根 1
	t4 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 5}}}

	fmt.Println(diameterOfBinaryTree(nil)) // 0
	fmt.Println(diameterOfBinaryTree(t1))  // 3
	fmt.Println(diameterOfBinaryTree(t2))  // 1
	fmt.Println(diameterOfBinaryTree(t3))  // 2
	fmt.Println(diameterOfBinaryTree(t4))  // 3（不过根：最高点是 2，不是根 1）

	// V1 与 V0 结果一致
	fmt.Println(diameterOfBinaryTreeV1(t1), diameterOfBinaryTreeV1(t3), diameterOfBinaryTreeV1(t4)) // 3 2 3
}
