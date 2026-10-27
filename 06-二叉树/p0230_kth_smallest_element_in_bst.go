//go:build p0041

// LeetCode 第230题 - 二叉搜索树中第K小的元素 (Kth Smallest Element in a BST)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、二叉搜索树、二叉树
//
// 【题目描述】
// 给定一个二叉搜索树的根节点 root，和一个整数 k，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化？
//
// 【示例 1】
// 输入：root = [3,1,4,null,2], k = 1
// 输出：1
//
// 【示例 2】
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3
//
// 【数据范围】
// 树中的节点数为 n
// 1 <= k <= n <= 10^4
// 0 <= Node.val <= 10^4
//
// 【思路提示】中序遍历 BST 即升序，计数到第 k 个提前返回。
// 进阶：每个节点维护子树大小，按 size 走树查找 O(h)。
//
// 题目链接：https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest 返回 BST 中第 k 小的元素。
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历的第K个元素

	var dfs func(root *TreeNode)
	var index int
	var result int
	dfs = func(root *TreeNode) {
		if root == nil || index >= k { // ✅ 双重终止：空节点 / 已找到（提前返回的剪枝）
			return
		}

		dfs(root.Left)  // 左
		index++         // 根：中序计数（BST 中序 = 升序，第 k 个即第 k 小）✅
		if index == k { // ✅ 命中：记录答案
			result = root.Val
		}
		dfs(root.Right) // 右
	}

	dfs(root)

	return result
}

// 【V0 点评】
// 完全正确！中序 + 计数 + 提前返回，标准写法一个不少：
//   ✅ 核心性质用对了：BST 中序遍历 = 升序序列 → 第 k 个访问的就是第 k 小
//   ✅ index >= k 剪枝：找到后整棵子树的递归被短路，不再白跑（很多解漏了这步）
//   ✅ index == k 记录后没有立刻 return 而是靠剪枝兜住——功能等价，
//      但"命中即 return"更直接（见 V1 的写法差异）
//   ✅ 时间 O(H+k)：H 是树高（先沉到最左），比无剪枝的 O(n) 更精确
//   ✅ 空间 O(h) 递归栈
//   验证示例1 [3,1,4,null,2], k=1：中序 1,2,3,4 → 第 1 个 = 1 ✓
//   验证示例2 [5,3,6,2,4,null,null,1], k=3：中序 1,2,3,4,5,6 → 第 3 个 = 3 ✓
//   验证边界：单节点 k=1 ✓；k=n（最后一个）✓；左偏链 k 中间值 ✓
//   🟡 main 只测了 kthSmallest(nil, 1)（返回 0，无意义），示例没验证（见下方 main）
//   🟡 可打磨细节：命中后用 return 提前退出比"记完继续走但被剪枝"语义更清晰（V1）
//   🟡 进阶问题没答：题目问"频繁查找如何优化"——答案见 V2（子树大小）
// 无需修改——可直接提交 AC

// kthSmallestV1 【LeetCode 标准答案】中序迭代栈版，命中即返回，时间 O(H+k)，空间 O(h)
func kthSmallestV1(root *TreeNode, k int) int {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		// 一路向左压栈（中序的"左"阶段）
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 弹出 = 中序访问"根"
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val // ✅ 命中即返回，剩下的完全不碰
		}
		// 转向右子树
		cur = cur.Right
	}
	return -1 // 不会到达（题目保证 1 <= k <= n）
}

// 【V1 点评 - 中序迭代栈版】
// ✅ 迭代中序的标准套路：外层 while + 内层"一路向左压栈"，弹出即访问，再转右
// ✅ 用 k-- 倒数代替 index++ 正数：k 减到 0 就是答案，少一个变量
// ✅ 命中即 return：真正的不访问剩余节点（V0 靠剪枝兜，这个更彻底）
// ✅ 时间 O(H+k)：先下沉 H 步到最左，再弹出 k 次；空间 O(h) 显式栈
// 与 V0 对比：递归版好写好读（面试先写它），迭代版展示"中序的栈实现"——
// 这也是"DFS 可以用迭代实现"的又一例证（呼应 p0199 的概念辨析）

// ──────────────────────────────────────────────────────────
// 【进阶答案：频繁修改 + 频繁查询的优化】
//
// 题目追问："BST 经常插入/删除，还要频繁查第 k 小，怎么优化？"
// 思路：每个节点维护【子树大小 size】（以该节点为根的子树节点数）
//
// type SizedNode struct {
//     Val   int
//     Size  int // 左右子树节点数 + 1
//     Left, Right *SizedNode
// }
//
// 查询 O(h)：从根往下走——
//   左子树大小 leftSize = size(node.Left)
//   k == leftSize+1  → node 就是第 k 小
//   k <= leftSize    → 答案在左子树，往左走
//   k > leftSize+1   → 答案在右子树，找第 k-leftSize-1 小，往右走
//
// 插入/删除 O(h)：沿路径修改每个节点的 size（+1/-1）
// 本质：这就是"顺序统计树"（order statistic tree），
// 把"数一遍 k 个"的 O(H+k) 优化成"按 size 导航"的 O(h)
// 面试只需讲清思路 + 复杂度，不必现场写完整实现
// ──────────────────────────────────────────────────────────

// kthSmallestV2 进阶实现：子树大小导航版，查询 O(h)，空间 O(h)
// （演示用：给 TreeNode 动态配一个 size map，真实工程应把 size 放进节点结构）
func kthSmallestV2(root *TreeNode, k int) int {
	size := make(map[*TreeNode]int) // 节点 → 子树大小

	var compute func(node *TreeNode) int
	compute = func(node *TreeNode) int { // 自底向上算 size：size = 左 + 右 + 1
		if node == nil {
			return 0
		}
		s := compute(node.Left) + compute(node.Right) + 1
		size[node] = s
		return s
	}
	compute(root) // 预处理 O(n)，之后每次查询 O(h)

	node := root
	for node != nil {
		leftSize := size[node.Left]
		switch {
		case k == leftSize+1:
			return node.Val // node 恰好排第 leftSize+1
		case k <= leftSize:
			node = node.Left // 答案在左子树
		default:
			k -= leftSize + 1 // 扣掉左子树和 node 自己
			node = node.Right
		}
	}
	return -1
}

// 【V2 点评 - 子树大小导航版】
// ✅ 演示进阶答案：预处理 size（O(n) 一次），每次查询 O(h) 按大小导航
// ✅ 三分支导航：k==leftSize+1 命中 / k 在左边 / k 在右边（减去左+自身）
// ✅ 频繁修改场景：插入/删除沿路径 O(h) 更新 size，查询仍是 O(h)——总量远小于每次 O(H+k)
// 🟡 演示用 map 存 size：真实实现应把 Size 字段放进节点（插入删除时顺手维护）
// 面试定位：讲清楚"size 导航 + O(h) 查询"的思路即可拿满进阶分

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间         | 空间   | 面试期望      |
// |-----------------------------|--------------|--------|---------------|
// | 中序递归+计数（V0）          | O(H+k)       | O(h)   | ✅ 标准答案   |
// | 中序迭代栈（V1）             | O(H+k)       | O(h)   | ✅ 同为标准   |
// | 中序全收集再取第 k 个        | O(n)         | O(n)   | ⚠️ 能过但平庸 |
// | 子树大小导航（V2 进阶）      | 预处理O(n)/查询O(h) | O(n) | ✅ 进阶满分 |
//
// 【核心易错点】
// 1. BST 中序 = 升序，是本题一切解法的基础性质（p0098 也用过它的等价形式）
// 2. k 从 1 开始计数：V1 用 k-- 倒数时，减到 0 恰好是第 k 个，别多减一次
// 3. 提前返回/剪枝：找到后不再遍历剩余——O(H+k) 而非 O(n)
// 4. 进阶追问必考："频繁修改怎么办"——答"节点维护子树 size，查询 O(h)"
//
// 【面试讲解建议】
// 1. 先说性质："BST 中序遍历就是升序，数到第 k 个即可"
// 2. 写递归中序 + 计数（或迭代栈版），强调命中即返回
// 3. 复杂度说精确："O(H+k)，先沉到最左再弹 k 个"
// 4. 进阶主动答："每个节点存子树大小，按 size 三分支导航，查询 O(h)——
//    这就是顺序统计树的思路"
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【附：中序遍历全过程图解（递归 vs 迭代）】
//
// 以本题示例 2 的树为例，画中序遍历（左→根→右）的完整过程：
//
//              5
//            /   \
//           3     6
//          / \
//         2   4
//        /
//       1
//
// 中序结果（升序）：1, 2, 3, 4, 5, 6
//
// ── 递归过程（调用栈视角）──
// 递归的本质：每个节点都要"三件事"——走左、访问自己、走右。
// 缩进表示调用深度，→ 访问 表示真正取出值的时刻：
//
//   dfs(5)
//   │
//   ├── dfs(5.Left) = dfs(3)          ① 走左
//   │   │
//   │   ├── dfs(3.Left) = dfs(2)      ① 走左
//   │   │   │
//   │   │   ├── dfs(2.Left) = dfs(1)  ① 走左
//   │   │   │   │
//   │   │   │   ├── dfs(1.Left) = dfs(nil) → 返回
//   │   │   │   ├── → 访问 1          ★ 输出 1
//   │   │   │   └── dfs(1.Right) = dfs(nil) → 返回
//   │   │   │
//   │   │   ├── → 访问 2              ★ 输出 2
//   │   │   └── dfs(2.Right) = dfs(nil) → 返回
//   │   │
//   │   ├── → 访问 3                  ★ 输出 3
//   │   └── dfs(3.Right) = dfs(4)     ② 走右
//   │       │
//   │       ├── dfs(4.Left) = dfs(nil) → 返回
//   │       ├── → 访问 4              ★ 输出 4
//   │       └── dfs(4.Right) = dfs(nil) → 返回
//   │
//   ├── → 访问 5                      ★ 输出 5
//   └── dfs(5.Right) = dfs(6)         ② 走右
//       │
//       ├── dfs(6.Left) = dfs(nil) → 返回
//       ├── → 访问 6                  ★ 输出 6
//       └── dfs(6.Right) = dfs(nil) → 返回
//
// 观察两个关键现象：
//   1. "访问"发生在两次递归【之间】——先沉到最左（1），再一路"回头访问"。
//      这就是为什么中序是升序：每个节点都是"左边的都访问完了，才轮到我"。
//   2. 栈深度 = 树高：最深时栈里压着 dfs(5)→dfs(3)→dfs(2)→dfs(1) 四层——
//      这就是空间 O(h) 的来源。
//
// ── 迭代过程（显式栈视角，V1 的写法）──
// 迭代用显式栈模拟递归的"走左"阶段；弹出即"访问"，再转右：
//
//   步骤   动作                     栈（底→顶）    输出
//   ─────  ──────────────────────  ────────────  ──────────
//    0     cur=5,5,3,2,1 依次压栈   [5 3 2 1]     （沉到最左）
//    1     弹 1，访问               [5 3 2]       1
//    2     1 无右，弹 2，访问       [5 3]         1 2
//    3     2 无右，弹 3，访问       [5]           1 2 3
//    4     3 的右是 4，压 4         [5 4]         1 2 3
//    5     弹 4，访问               [5]           1 2 3 4
//    6     4 无右，弹 5，访问       []            1 2 3 4 5
//    7     5 的右是 6，压 6         [6]           1 2 3 4 5
//    8     弹 6，访问               []            1 2 3 4 5 6
//    9     6 无右，栈空 → 结束                    1 2 3 4 5 6 ✅
//
//   对照树看压栈轨迹：
//              5          迭代轨迹：
//            /   \        ① 一路向左：5→3→2→1 全压栈
//           3     6       ② 弹1访问 → 弹2访问 → 弹3访问
//          / \            ③ 3 有右：压4 → 弹4访问
//         2   4           ④ 弹5访问（根最后轮到，因为中序）
//        /                ⑤ 5 有右：压6 → 弹6访问
//       1                 ⑥ 栈空且 cur=nil → 结束
//
// ── 递归 vs 迭代的对应关系 ──
//
//   递归                           迭代
//   ─────────────────────────     ─────────────────────────
//   隐式调用栈（函数自己压）        显式栈（自己 push/pop）
//   dfs(node.Left) 递归到底        for 内层：一路向左压栈
//   回到 node 时"访问"             弹出时"访问"
//   dfs(node.Right)               cur = cur.Right
//
//   同一棵树的同一条路：递归靠函数调用机制压栈，迭代靠手动 append/pop 压栈——
//   访问顺序分毫不差。这印证 p0199 的结论：DFS 绑定的是"栈"，
//   递归只是借用了调用栈。
//
// ── 三种遍历顺序的对比（顺带记忆）──
//
//   只需改"访问自己"的位置，走法完全相同：
//
//     前序（根左右）：访问在两个递归【前】   → 5,3,2,1,4,6
//     中序（左根右）：访问在两个递归【中】   → 1,2,3,4,5,6 ★ BST 升序
//     后序（左右根）：访问在两个递归【后】   → 1,2,4,3,6,5
//
//   同一棵树、同样的递归结构，只挪一行代码，顺序完全不同——
//   这也是 p0226 里"V0 前序思维 vs V1 后序思维"的底层机制。
// ──────────────────────────────────────────────────────────

func main() {
	// 示例1：[3,1,4,null,2], k=1 → 1
	t1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4}}
	// 示例2：[5,3,6,2,4,null,null,1], k=3 → 3（中序 1,2,3,4,5,6）
	t2 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6}}
	// 左偏链：[3,2,null,1], k=2 → 2（中序 1,2,3）
	t3 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}

	fmt.Println(kthSmallest(t1, 1)) // 1
	fmt.Println(kthSmallest(t1, 4)) // 4（k=n 最后一个）
	fmt.Println(kthSmallest(t2, 3)) // 3
	fmt.Println(kthSmallest(t3, 2)) // 2

	// V1/V2 与 V0 结果一致
	fmt.Println(kthSmallestV1(t1, 1), kthSmallestV1(t2, 3), kthSmallestV1(t3, 2)) // 1 3 2
	fmt.Println(kthSmallestV2(t1, 1), kthSmallestV2(t2, 3), kthSmallestV2(t3, 2)) // 1 3 2
}
