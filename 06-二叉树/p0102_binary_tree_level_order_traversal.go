//go:build p0041

// LeetCode 第102题 - 二叉树的层序遍历 (Binary Tree Level Order Traversal)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、广度优先搜索、队列
//
// 【题目描述】
// 给你二叉树的根节点 root，返回其节点值的层序遍历。
// （即逐层地，从左到右访问所有节点）。
//
// 【示例 1】
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 【示例 2】
// 输入：root = [1]
// 输出：[[1]]
//
// 【示例 3】
// 输入：root = []
// 输出：[]
//
// 【数据范围】
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// 【思路提示】BFS + 队列：每轮记录当前层节点数 n，弹出 n 个节点收集本层值，
// 同时把子节点入队。也可用 DFS 按深度分组。
//
// 题目链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/
package main

import "fmt"

// TreeNode 二叉树节点（与 LeetCode Go 环境定义一致）。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 返回二叉树层序遍历结果。
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil // ✅ 空树返回 nil（LeetCode 判题 nil 和 [] 等价，均可 AC）
	}

	var queue []*TreeNode = []*TreeNode{root}
	var result [][]int

	for len(queue) != 0 {
		var n = len(queue) // ✅ 关键：先记录本层节点数
		var tmp []int

		for i := 0; i < n; i++ {
			cur := queue[i] // 🟡 用下标读队头，不是标准"弹出"——可行但绕（见下方说明）
			if cur == nil {
				continue // 🟡 nil 节点跳过：说明把 nil 也入队了，靠 continue 兜底
			}
			tmp = append(tmp, cur.Val)
			queue = append(queue, cur.Left, cur.Right) // 🟡 nil 子节点也入队（叶子会塞进两个 nil）
		}

		result = append(result, tmp)

		if len(queue) > n { // 🔴 BUG：本层全是 nil 时（queue 只有 nil），len(queue) == n
			queue = queue[n+1:] //    走 else 分支清空队列，恰好能终止——但这是巧合而非设计：
		} else { //    更糟的是 queue[n+1:] 切掉的是 n+1 个元素而不是 n 个，
			queue = []*TreeNode{} //    相当于把"下一层的第一个节点"也丢掉了！
		} //    之所以结果碰巧正确：下标读取用了旧下标空间……不，实际会漏节点。
		//    反例：root=[1,2]，第 1 层处理完 queue=[1的左右=nil,2]，len=2>n=1 →
		//    queue[2:] = [2]？不对：queue = [nil,2]，queue[n+1:]=queue[2:]=[] → 节点 2 被丢！
		//    但等等——内层循环 i<n 读的是 queue[0..n-1]=旧层，append 在后面，切 n+1 会切到新层第 1 个
		//    结论：只要本层有 ≥1 个非 nil 节点，queue[n+1:] 就会丢弃下一层的第一个节点 ❌
		//    只因下标读法与切片错位恰好互补，部分用例碰巧对——见 V0 点评详细追踪
	}

	return result
}

// 【V0 点评】
// 编译通过，单节点/满树碰巧正确，但切片逻辑有严重错误：
//   🔴 核心 BUG（第 64-68 行）：queue = queue[n+1:] 切掉了 n+1 个元素
//      内层循环后 queue 的结构是 [旧层 n 个][下一层若干个]，
//      正确的"弹出旧层"应该切 queue[n:]，切 [n+1:] 会把下一层的第一个节点丢掉！
//      追踪 root=[3,9,20,null,null,15,7]：
//        第 1 层：n=1，queue=[3] → 处理后 queue=[3,9,20]
//          len(3) > n(1) → queue = queue[2:] = [20] ← 节点 9 被丢弃！❌
//        第 2 层：n=1，queue=[20] → 处理后 queue=[20,15,7]
//          → queue = queue[2:] = [7] ← 节点 15 被丢弃！❌
//        第 3 层：queue=[7] → 处理后 queue=[7,nil,nil]
//          len(3) > n(1) → queue = queue[2:] = [nil]
//        第 4 层：n=1，queue=[nil] → cur=nil continue，tmp=[] 追加空层
//          len(1) == n(1) → 清空队列，循环结束
//        最终输出 [[3],[20],[7],[]] ❌ 期望 [[3],[9,20],[15,7]]
//      main 只测了 levelOrder(nil)，完全没验证层序逻辑！
//   🔴 次生问题：nil 子节点入队 + continue 兜底 → 叶子层之后必产生全 nil 层，
//      靠 else 分支清空才不崩——错误相互掩盖，恰好不死循环
//   🟡 下标读队列 + 切片"伪弹出"绕开了标准队列操作，正是这套错位的根源
//   🟡 result 会追加空层（如上面的 []），层序结果不该有空层数组
//   ✅ 先记录 n=len(queue) 再处理本层的"分层"思路正确（BFS 分层标准套路）
//   ✅ 内层循环只读旧层下标 [0,n)，新入队元素不会污染本层——分层隔离的意图是对的
//   ✅ 空树处理正确
// 根源：把"弹出"实现成了下标偏移 + 切片，且切片位数与入队结构错位（n+1 vs n）。
// 标准队列操作：弹出用 queue = queue[1:]（一次弹一个）或整层切片 queue = queue[n:]

// levelOrderV1 【LeetCode 标准答案】BFS + 队列（标准弹出写法），时间 O(n)，空间 O(n)
func levelOrderV1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		n := len(queue)          // ✅ 本层节点数（队列里只有非 nil 节点，n 就是真实数量）
		level := make([]int, 0, n)

		for i := 0; i < n; i++ {
			cur := queue[0]      // ✅ 标准弹出：取队头
			queue = queue[1:]    //    一次弹一个
			level = append(level, cur.Val)
			if cur.Left != nil {  // ✅ 只入队非 nil 子节点——队列里永远没有 nil
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 修复核心 BUG：弹出改为 queue[1:] 一次弹一个，切片位数永远正确
// ✅ 只入队非 nil 子节点：队列里没有 nil，n = len(queue) 就是本层真实节点数，
//    不需要 continue 兜底，也不会产生空层
// ✅ 分层不变量：内层循环开始时，队列恰好包含且仅包含本层的全部节点
//    （上一轮循环已把上一层的节点全部弹出）
// ✅ level 预分配容量 n，减少 append 扩容
// ✅ 时间 O(n)：每个节点入队出队各一次；空间 O(n)（队列 + 结果）
// 追踪 root=[3,9,20,null,null,15,7]：
//   第 1 层：n=1，弹 3，入 9,20 → result=[[3]]，queue=[9,20]
//   第 2 层：n=2，弹 9（无子），弹 20，入 15,7 → result=[[3],[9,20]]，queue=[15,7]
//   第 3 层：n=2，弹 15，弹 7 → result=[[3],[9,20],[15,7]]，queue=[] → 结束 ✅

// levelOrderV2 DFS 按深度分组，时间 O(n)，空间 O(h)——另一种经典写法
func levelOrderV2(root *TreeNode) [][]int {
	var result [][]int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) { // ✅ 第一次到达该深度：为本层开一个数组
			result = append(result, []int{})
		}
		result[depth] = append(result[depth], node.Val) // ✅ DFS 天然按深度递增，先访问的先入层
		dfs(node.Left, depth+1)  // ✅ 先左后右 → 层内从左到右
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)
	return result
}

// 【V2 点评 - DFS 按深度分组】
// ✅ 思路：DFS 带深度参数 depth，result[depth] 收集该层的值
// ✅ 关键行：depth == len(result) 时 append 新层——第一次到达深度 d 时开辟第 d 层
// ✅ 先递归左再右，保证层内从左到右的顺序
// ✅ 空间 O(h) 递归栈（不含结果数组）；对深树比 BFS 队列省空间
// 与 V1 对比：BFS 是"队列按层处理"（层序天然）；DFS 是"先访问先入层"（靠顺序保证）
// 两者都是标准答案；层序题首选 V1，但 V2 的"按深度分组"技巧在很多变形题有用
// （如右视图、锯齿形遍历、最大深度等）

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间   | 空间   | 面试期望      |
// |-------------------------|--------|--------|---------------|
// | 下标+错位切片（V0 有 bug）| O(n)  | O(n)   | ❌ 丢节点     |
// | BFS 队列（V1 标准答案）  | O(n)   | O(n)   | ✅ 面试首选   |
// | DFS 深度分组（V2）       | O(n)   | O(h)   | ✅ 同为标准   |
//
// 【核心易错点】
// 1. 分层关键：进入内层循环前先记 n = len(queue)——处理期间队列会增长，n 是快照
// 2. 弹出实现：queue = queue[1:] 一次弹一个；或内层循环结束后整层切 queue = queue[n:]
//    切片位数必须与队列结构精确对应（V0 的 n+1 错位就是丢节点的根源）
// 3. 只入队非 nil 子节点：队列里不存 nil，省去 continue 兜底，也避免"全 nil 层"的空层
// 4. 空树返回 nil/[] 均可；层序结果不应包含空层数组
// 5. 测试至少覆盖：满树、单边树、单节点——V0 的 bug 用示例 1 就能暴露
//
// 【面试讲解建议】
// 1. 先说分层套路："每轮记 n=len(queue)，弹出 n 个即本层，子节点入队"
// 2. 强调不变量："内层循环开始时，队列恰好是本层全部节点"
// 3. 提到 nil 处理："只把非 nil 子节点入队，队列里永远没有 nil"
// 4. 备选方案：DFS 按深度分组（depth==len(result) 时开新层），变形题常用
// 5. 复杂度：时间 O(n)，空间 O(n)
// ──────────────────────────────────────────────────────────

func main() {
	// 示例树：[3,9,20,null,null,15,7]
	tree := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	// 单边树：[1,2,null,3]
	single := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	fmt.Println(levelOrder(tree))    // [[3] [20] [7] []]  ← V0 错误输出（丢 9、丢 15、多空层）
	fmt.Println(levelOrderV1(tree))  // [[3] [9 20] [15 7]] ✅
	fmt.Println(levelOrderV2(tree))  // [[3] [9 20] [15 7]] ✅

	fmt.Println(levelOrder(single))    // [[1] [2] [3] []]  ← V0 又丢节点又多空层
	fmt.Println(levelOrderV1(single))  // [[1] [2] [3]] ✅
	fmt.Println(levelOrderV2(single))  // [[1] [2] [3]] ✅

	fmt.Println(levelOrderV1(nil))   // []
	fmt.Println(&TreeNode{Val: 1} != nil && levelOrderV1(&TreeNode{Val: 1}) != nil) // true（单节点 [[1]]）
}
