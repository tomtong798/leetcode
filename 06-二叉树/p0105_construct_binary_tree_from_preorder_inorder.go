//go:build p0041

// LeetCode 第105题 - 从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】树、数组、哈希表、分治、二叉树
//
// 【题目描述】
// 给定两个整数数组 preorder 和 inorder，其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//
// 【示例 1】
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 【示例 2】
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
//
// 【数据范围】
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// 无重复元素
//
// 【思路提示】前序首元素为根，哈希表存中序值→下标 O(1) 定位左子树长度，
// 递归构造左右子树。用全局 idx 指针推进前序更简洁。
//
// 题目链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package main

import "fmt"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 从前序与中序构造二叉树。
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil // ✅ 递归基：空序列 → 空子树
	}

	// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// 输出: [3,9,20,null,null,15,7]

	// 前序的第一个元素是根结点
	// 在中序中找到根结点，根结点右边是右子树，左边是左子树

	rootVal := preorder[0]
	var rootIndex int
	for i, e := range inorder { // 🟡 线性扫描定位根：每层 O(n)，总体 O(n²)
		if e == rootVal { //    3000 节点最坏 9×10^6 次比较，LeetCode 勉强能过但不优
			rootIndex = i //    标准做法：哈希表 值→中序下标，定位 O(1)（见 V1）
			break
		}
	}

	leftInorder := inorder[:rootIndex]      // ✅ 中序左段 = 左子树
	rightInorder := inorder[rootIndex+1:]   // ✅ 中序右段 = 右子树

	leftPreorder := preorder[1 : len(leftInorder)+1] // ✅ 前序[1..左长] = 左子树（跳过根，取左子树长度个）
	rightPreorder := preorder[len(leftInorder)+1:]   // ✅ 剩余 = 右子树

	root := &TreeNode{Val: rootVal}

	root.Left = buildTree(leftPreorder, leftInorder)  // ✅ 切片对应关系正确
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

// 【V0 点评】
// 完全正确！分治构造的经典写法，边界切片全对：
//   ✅ 核心三步全对：前序首元素定根 → 中序定位切左右 → 前序按左子树长度切段
//   ✅ 切片对应关系（最容易错的部分）：
//      leftPreorder  = preorder[1 : 1+len(leftInorder)]——根后面紧跟左子树的全部节点
//      rightPreorder = preorder[1+len(leftInorder):] ——剩下的是右子树
//      两段长度分别与 leftInorder/rightInorder 相等，递归参数自洽
//   ✅ 递归基 len==0 返回 nil；题目保证无重复元素，线性查找必命中
//   ✅ 验证：示例 [3,9,20,15,7]/[9,3,15,20,7] 构造出 [3,9,20,null,null,15,7] ✓
//      手工用例 [1,2]/[2,1] → 根1，左子树2 ✓；[1,2]/[1,2] → 根1，右子树2 ✓
//   🟡 效率：线性扫描找根，最坏（链状树）O(n²)；哈希表预定位可到 O(n)（见 V1）
//   🟡 每层递归都切片拷贝（4 个新切片），有额外 O(n) 空间/层；
//      V1 用 (preStart, inStart, inEnd) 下标区间代替切片，零拷贝
//   🟡 main 打印的是指针地址（*TreeNode 的默认 String），看不出结构——建议层序打印验证
// 无需修改即可 AC；V1 是面试更优写法

// buildTreeV1 【LeetCode 标准答案】哈希定位 + 下标区间递归，时间 O(n)，空间 O(n)
func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	// 哈希表：中序值 → 下标，定位根 O(1)
	idxMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		idxMap[v] = i
	}

	// preIdx 是全局推进的前序"游标"：构造顺序恰好与前序遍历顺序一致
	preIdx := 0

	// 在中序 [inStart, inEnd) 区间内构造
	var build func(inStart, inEnd int) *TreeNode
	build = func(inStart, inEnd int) *TreeNode {
		if inStart >= inEnd {
			return nil // 空区间 → 空子树
		}
		rootVal := preorder[preIdx] // ✅ 前序游标当前位置就是本区间根
		preIdx++                    // ✅ 消费掉这个根，游标前移

		root := &TreeNode{Val: rootVal}
		rootIdx := idxMap[rootVal]          // ✅ O(1) 定位根在中序中的位置
		root.Left = build(inStart, rootIdx)  // ✅ 左子树占 [inStart, rootIdx)
		root.Right = build(rootIdx+1, inEnd) // ✅ 右子树占 [rootIdx+1, inEnd)
		return root
	}

	return build(0, len(inorder))
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 两处关键优化：
//    1) 哈希表定位根：O(n) 次 O(1) 查找替代 O(n) 次 O(n) 线性扫描 → 整体 O(n)
//    2) 前序游标 preIdx 替代切片：不需要切 leftPreorder/rightPreorder——
//       因为"左子树先构造完才轮到右子树"，游标天然按前序顺序推进，
//       每个前序元素恰好被消费一次，与递归构造顺序完美同步
// ✅ 中序只用下标区间 [inStart, inEnd) 划分，零切片拷贝
// ✅ 时间 O(n)，空间 O(n)（哈希表 + 递归栈）
// 追踪示例 preorder=[3,9,20,15,7], inorder=[9,3,15,20,7]：
//   build(0,5): 根=pre[0]=3, preIdx→1, idxMap[3]=2
//     build(0,2): 根=pre[1]=9, preIdx→2, idxMap[9]=0
//       build(0,0)=nil, build(1,2)=nil → 9 是叶子 ✓
//     build(3,5): 根=pre[2]=20, preIdx→3, idxMap[20]=3
//       build(3,3)=nil
//       build(4,5): 根=pre[3]=15 → 叶子；build 内 preIdx→4, 其右 build(5,5)=nil
//       ... 最终 20 的左=15，右=7 ✓
// 关键洞察：为什么 preIdx 全局推进是对的？
//   前序 = [根, [左子树前序], [右子树前序]]，递归先构造整个左子树再构造右子树，
//   游标按"根→左→右"的顺序消费前序元素——与递归调用顺序严格一致，永不错位

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                          | 时间   | 空间   | 面试期望      |
// |-------------------------------|--------|--------|---------------|
// | 切片递归 + 线性找根（V0）      | O(n²)  | O(n)   | ⚠️ 正确但慢   |
// | 哈希 + 前序游标（V1 标准答案） | O(n)   | O(n)   | ✅ 面试首选   |
// | 哈希 + 区间参数递归（不切前序）| O(n)   | O(n)   | ✅ 同为标准   |
//
// 【核心易错点】
// 1. 切片对应关系是本题最容易写错的地方：
//    leftPreorder  = preorder[1 : 1+len(leftInorder)]（根后紧跟左子树全部）
//    rightPreorder = preorder[1+len(leftInorder):]（剩余是右子树）
//    两段长度必须与中序左右段相等
// 2. 前序游标为什么全局推进不乱：递归顺序（根→左→右）与前序定义一致
// 3. 为什么必须前序+中序成对：前序定根，中序分左右；
//    只有两个中序或两个前序无法唯一确定（缺"分左右"的信息）
// 4. 无重复元素是哈希定位的前提（本题保证）
//
// 【面试讲解建议】
// 1. 先说三步："前序首元素定根 → 中序切左右 → 递归构造"
// 2. 主动提优化："线性找根 O(n²)，用哈希表降到 O(n)"
// 3. 讲游标技巧："前序游标全局推进，因为递归顺序与前序一致"
// 4. 复杂度：时间 O(n)，空间 O(n)
// ──────────────────────────────────────────────────────────

func main() {
	printTree(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))  // [[3] [9 20] [15 7]]
	printTree(buildTree([]int{-1}, []int{-1}))                            // [[-1]]
	printTree(buildTree([]int{1, 2}, []int{2, 1}))                        // [[1] [2]]（左子树）
	printTree(buildTree([]int{1, 2}, []int{1, 2}))                        // [[1] [] [2]]（右子树）

	// V0/V1 构造结果完全一致
	printTree(buildTreeV1([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})) // [[3] [9 20] [15 7]]
	printTree(buildTreeV1([]int{1, 2, 3}, []int{2, 1, 3}))                 // [[1] [2 3]]
}

// printTree 层序打印树结构（nil 跳过），用于肉眼验证构造结果
func printTree(root *TreeNode) {
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
