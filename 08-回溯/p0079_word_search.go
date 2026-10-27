//go:build p0041

// LeetCode 第79题 - 单词搜索 (Word Search)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、回溯、矩阵、深度优先搜索
//
// 【题目描述】
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word。如果 word 存在于网格中，
// 返回 true；否则，返回 false。单词必须按照字母顺序，通过相邻的单元格内的字母构成，
// 其中"相邻"单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中
// 不允许被重复使用。
//
// 【示例 1】
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
//
// 【示例 2】
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出：true
//
// 【示例 3】
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// 输出：false
//
// 【数据范围】
// 1 <= board.length <= 6
// 1 <= board[i].length <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
// 【思路提示】网格 DFS 回溯模板：从每个格子尝试匹配 word[0]；
// 递归四方向扩展，用临时标记（如改为 '#'）表示"访问中"，回溯时还原。
//
// 题目链接：https://leetcode.cn/problems/word-search/
package main

import "fmt"

// exist 判断 word 是否存在于网格中。
func exist(board [][]byte, word string) bool {

	var helper func(cnt int, positionX int, positionY int, visited [][]bool) bool
	helper = func(cnt int, positionX int, positionY int, visited [][]bool) bool {
		if visited == nil { // 🟡 visited 初始化散落在递归里：每个格子惰性 make——
			visited = make([][]bool, len(board)) //    能跑但每条路径独立建表，浪费；
		} //    标准做法：外层一次性建好，或直接改格子为 '#'（见 V1）

		if visited[positionX] == nil {
			visited[positionX] = make([]bool, len(board[0]))
		}

		visited[positionX][positionY] = true // ✅ 标记访问

		if board[positionX][positionY] == word[cnt] { // 🟡 判断顺序反了：先标访问后查匹配——
			return helper(cnt+1, positionX+1, positionY, visited) || helper(cnt+1, positionX, positionY+1, visited) || // 🔴 BUG 1：没有越界检查！
				helper(cnt+1, positionX-1, positionY, visited) || helper(cnt+1, positionX, positionY-1, visited) //    位置±1 直接递归，走到网格外 panic
		} // 🔴 BUG 2：没有终止条件！cnt == len(word) 时（完全匹配成功）
		//    word[cnt] 越界 panic——成功的那一步反而崩溃
		// 🔴 BUG 3：没有回溯还原！visited 标了 true 永不撤销——
		//    四方向 || 短路尝试时，第一个方向失败的路径仍标记着 visited，
		//    污染后续方向（例 3 "ABCB" 的 B 复用问题反而被这个 bug 掩盖/加剧）
		// 🔴 BUG 4：visited 按行惰性初始化 + [][]bool{} 传入：
		//    每条路径的新表只继承"已 make 的行"——外层传 [][]bool{}（空表）时
		//    每个起点全新，但递归内 visited 是值传递的 slice header——
		//    行指针共享：同一路径的标记能传递 ✓，但撤销缺失让"传递"变成"污染"
		// 🔴 BUG 5：visited[...]==nil 判断：第二行首访时 visited[1]==nil ✓，
		//    但若路径先访问行 0 再跳到行 2 再回行 1——每次进 helper 都重新 make 该行？
		//    不——make 只在 nil 时执行一次，但这依赖"行内格子先于判断被标记"的副作用，极脆弱

		return false

	}

	// 先找到所有首字母的位置
	// 遍历所有首字母为起点的路径
	// ✅ 预收集起点的思路本身没问题（也可以直接双重循环调 helper，等价）

	var starts [][]int
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				starts = append(starts, []int{i, j})
			}
		}
	}

	if len(starts) == 0 {
		return false
	}

	for _, s := range starts {
		if helper(0, s[0], s[1], [][]bool{}) { // 🟡 每个起点传全新空表 ✓ 语义对，
			return true //    但 [][]bool{} 与 nil 走不同分支（nil 触发 make，空表不触发
		} //    外层 make(len(board))）——初始化逻辑分叉，易埋坑
	}

	return false
}

// 【V0 点评】
// 编译通过，但一跑就 panic（越界），且存在多个结构性缺失：
//   🔴 BUG 1（必崩）：四方向递归无越界检查——positionX±1 走出网格直接 index out of range
//      追踪示例 1 起点 (0,0) 匹配 A 后：递归 (1,0) 匹配 B？S≠B → false；
//      但 (0,-1)（左）在访问 word[1] 前就 board[0][-1] panic
//   🔴 BUG 2（成功即崩）：无 cnt == len(word) 终止——完全匹配后下一步 word[cnt] 越界 panic
//      就算有越界检查，成功路径也永远走不到 return true
//   🔴 BUG 3（核心缺失）：没有回溯还原——visited 标记后不撤销，
//      || 的四个方向共享污染状态；"同一单元格不允许重复使用"的题意也无法正确实现
//      （这正是思路提示里"回溯时还原"的那一步——回溯三件套的第三件）
//   🟡 BUG 4/5（脆弱设计）：visited 惰性初始化散落递归内 + nil/空表双路径，
//      行 make 依赖副作用时序——功能勉强但极难维护
//   🟡 判断顺序：先标 visited 再查匹配——匹配失败也标了（配合无撤销，污染放大）
//   ✅ 值得肯定：四方向 || 短路（任一方向成功即 true）的结构是对的；
//      预收集首字母起点的思路清晰；cnt 与 word 下标对应关系正确
//   网格回溯的完整模板 = 边界检查 + 匹配判断 + 终止条件 + 标记 + 四方向 + 还原，
//   V0 缺了其中 3.5 件（边界/终止/还原 + 判断顺序）——见 V1 对照

// existV1 【LeetCode 标准答案】网格回溯完整模板，时间 O(m·n·3^L)，空间 O(L)
func existV1(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// inBoard + 匹配统一判断：越界或不匹配都返回 false（负条件收口）
	var dfs func(x, y, cnt int) bool
	dfs = func(x, y, cnt int) bool {
		// ① 边界 + 匹配检查（合并为一个负条件）
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[cnt] {
			return false
		}
		// ② 终止：最后一个字符匹配成功（在上面已保证 board[x][y]==word[cnt]）
		if cnt == len(word)-1 {
			return true
		}
		// ③ 标记：原地改 '#'——当前格子不能再被本路径使用
		board[x][y] = '#'
		// ④ 四方向（|| 短路：任一成功即 true）
		found := dfs(x+1, y, cnt+1) || dfs(x-1, y, cnt+1) ||
			dfs(x, y+1, cnt+1) || dfs(x, y-1, cnt+1)
		// ⑤ 回溯还原——恢复现场，别的路径还能用这个格子
		board[x][y] = word[cnt]
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) { // 每个格子都可能是起点（匹配检查在 dfs 内部统一做）
				return true
			}
		}
	}
	return false
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 五件套完整（对照 V0 缺的）：
//    ① 边界+匹配合并成一个负条件——V0 的 BUG 1 修复，且匹配检查前移（先查后标）
//    ② cnt == len(word)-1 终止——V0 的 BUG 2 修复：最后一个字符匹配即成功，不再多走一步
//    ③ 原地标记 board[x][y] = '#'——不用 visited 数组，省掉初始化和传参
//       （题目保证 word 由字母组成，'#' 不可能误匹配）
//    ⑤ 还原 board[x][y] = word[cnt]——V0 的 BUG 3 修复，回溯三件套补全
// ✅ 起点不需要预收集：匹配检查在 dfs 入口统一做，起点不匹配直接 false
// ✅ 时间 O(m·n·3^L)：每个起点 × 每步 3 个方向（不算回头路，第一个格子 4 方向）
//    空间 O(L) 递归栈（L = word 长度）
// 追踪例 2 "SEE"：起点 (1,0)S → 失败；(2,0)A≠S；(1,3)S → 下 E：(2,3)E ✓ →
//    下 E：(2,2)E ✓ cnt==2==len-1 → true ✅
// 追踪例 3 "ABCB"：所有路径 B(0,1) 已用过 → '#'≠B → false ✅（标记+还原配合正确）

// ──────────────────────────────────────────────────────────
// 【网格回溯模板五件套（V0 缺哪件的对照表）】
//
//   func dfs(x, y, cnt) bool {
//     ① 边界+匹配   x/y 出界 || 格子≠word[cnt] → false     【V0 缺 → panic】
//     ② 终止        cnt==len-1 → true                       【V0 缺 → 成功即 panic】
//     ③ 标记        board[x][y] = '#'（或 visited=true）    【V0 有但顺序错：先标后查】
//     ④ 四方向      dfs(四邻, cnt+1) 用 || 短路              【V0 ✓ 唯一完整的一件】
//     ⑤ 还原        board[x][y] = word[cnt]（或 =false）    【V0 缺 → 状态污染】
//   }
//
// 两个惯用变体：
//   标记方式：visited[][]bool 传参 vs 原地改 '#'——后者省空间省传参（推荐）
//   检查位置：入口统一查（起点也走同一条路）vs 外层预筛起点——前者代码短（推荐）
//
// 与线性回溯（p0216/p0090）的差异：
//   候选集：i+1..9 的"剩余数字" vs 四邻格的"可走方向"
//   剪枝：和/个数约束 vs 字母匹配（匹配本身就是最强剪枝）
//   还原：path 弹出 vs 格子还原（'#' 改回字母）
//   共同灵魂不变：做选择 → 递归 → 撤销选择
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                        | 时间         | 空间   | 面试期望      |
// |-----------------------------|--------------|--------|---------------|
// | V0（缺边界/终止/还原）       | —            | —      | ❌ 必 panic   |
// | 网格回溯 V1（标准答案）      | O(mn·3^L)    | O(L)   | ✅ 面试首选   |
// | Trie 优化（word 列表版）     | 见 p212      | O(单词总长) | ⚠️ 单词搜索 II 用 |
//
// 【核心易错点】
// 1. 越界检查必须最先做——四方向 ±1 天然会出界
// 2. 终止条件的写法：cnt == len(word)-1（最后一个字符匹配即成），
//    不是 cnt == len(word)（那一步 word[cnt] 已越界）——V0 的 BUG 2 正是此坑
// 3. 标记与还原必须配对出现——"#"改出去就要改回来，缺还原 = 后续全错
// 4. 匹配检查放在标记之前：不匹配的格子不该被标记
// 5. board 由字母组成，'#' 是安全的占位符（不会与 word 误匹配）
//
// 【面试讲解建议】
// 1. 先报模板："边界+匹配 → 终止 → 标记 → 四方向 → 还原"五件套
// 2. 讲 '#" 标记的巧思：省 visited 数组，还原时用 word[cnt]（就是原字符）
// 3. 复杂度：O(mn·3^L)——解释 3 不是 4（不走回头路）
// 4. 延伸：单词搜索 II（p212）用 Trie 前缀树批量匹配——被追问时的加分项
// ──────────────────────────────────────────────────────────

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	// 找到所有目标首字母在 board 中的位置
	// 以所有找到的位置开始遍历

	fmt.Println(existV1(board, "ABCCED")) // true
	fmt.Println(existV1(board, "SEE"))    // true
	fmt.Println(existV1(board, "ABCB"))   // false（B 不可复用）

	// V0 的 panic 复现（取消注释一跑就崩：无越界检查）
	// fmt.Println(exist(board, "ABCCED")) // panic: index out of range
	_ = exist
}
