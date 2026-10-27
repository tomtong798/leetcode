//go:build p0041

// LeetCode 第62题 - 不同路径 (Unique Paths)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数学/动态规划/组合数学
//
// 【题目描述】
// 一个机器人位于 m x n 网格的左上角，每次只能向下或向右移动一步，问到达右下角总共有多少条不同的路径。
//
// 【示例 1】
// 输入：m = 3, n = 7
// 输出：28
// 【数据范围】
// 1 <= m, n <= 100；题目数据保证答案小于等于 2*10^9
//
// 【思路提示】dp[i][j] = dp[i-1][j] + dp[i][j-1]；或直接用组合数 C(m+n-2, m-1)。
//
// 题目链接：https://leetcode.cn/problems/unique-paths/
package main

// uniquePaths 返回 m x n 网格从左上到右下的不同路径数。

// 从位置 （0,0） 到 （m-1，n-1） 的不同路径
func uniquePaths(m int, n int) int {
	var result int

	var curQueue = [][]int{{0, 0}}

	for len(curQueue) > 0 { // 🔴 这是在枚举所有路径，时间复杂度 O(2^(m+n))
		var nextQueue = [][]int{}

		for _, ele := range curQueue {
			x, y := ele[0], ele[1]
			if x == m-1 && y == n-1 {
				result++
				continue
			}

			if x < m-1 {
				nextQueue = append(nextQueue, []int{x + 1, y})
			}
			if y < n-1 {
				nextQueue = append(nextQueue, []int{x, y + 1})
			}
		}

		curQueue = nextQueue
	}

	return result
}

// 【V0 点评】
// 编译通过，小数据正确，但算法复杂度过高：
//   🔴 算法是暴力枚举所有路径，时间复杂度 O(2^(m+n))，指数级
//   🔴 同一个位置被重复访问多次，没有记忆化/去重
//   🔴 当 m=100, n=100 时必然超时（路径数约 10^58）
//   🟡 curQueue/nextQueue 命名暗示 BFS，但实际是路径枚举
// 正确做法：动态规划 O(mn) 或组合数学 O(m)

// uniquePathsV1 用动态规划，时间 O(mn)，空间 O(n)
func uniquePathsV1(m int, n int) int {
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1 // 第一行全是 1（只能从左边来）
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] // dp[j] 是上方（上一行的值），dp[j-1] 是左方（当前行已更新的值）
		}
	}
	return dp[n-1]
}

// 【V1 点评 - 动态规划】
// ✅ 时间 O(mn)，空间 O(n)（滚动数组优化）
// ✅ 思路：dp[i][j] = dp[i-1][j] + dp[i][j-1]（上方 + 左方）
// ✅ 用一维数组滚动更新，空间从 O(mn) 优化到 O(n)
// 面试推荐写法，简洁高效

// uniquePathsV2 用组合数学，时间 O(min(m,n))，空间 O(1)
func uniquePathsV2(m int, n int) int {
	// 从 (0,0) 到 (m-1,n-1) 需要走 m-1 步向下、n-1 步向右，共 m+n-2 步
	// 问题等价于：从 m+n-2 步中选 m-1 步向下，即 C(m+n-2, m-1)
	a, b := m-1, n-1
	if a > b {
		a, b = b, a // 取较小的减少计算量
	}
	result := 1
	for i := 0; i < a; i++ {
		result = result * (m + n - 2 - i) / (i + 1)
	}
	return result
}

// 【V2 点评 - 组合数学】
// ✅ 时间 O(min(m,n))，空间 O(1)
// ✅ 思路：路径总数 = C(m+n-2, m-1) = (m+n-2)! / ((m-1)! * (n-1)!)
// ✅ 最优解法，但需要理解组合数学推导
// 面试时先写 V1（DP），如果被追问"有没有更优解"再写 V2

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                    | 时间        | 空间      | 面试期望          |
// |-------------------------|-------------|-----------|-------------------|
// | 暴力枚举（V0）           | O(2^(m+n))  | O(路径数) | ❌ 超时            |
// | 动态规划滚动数组（V1）    | O(mn)       | O(n)      | ✅ 推荐写法        |
// | 组合数学（V2）           | O(min(m,n)) | O(1)      | ✅ 最优，可追问时写 |
// | 递归无记忆化（V3）        | O(2^(m+n))  | O(m+n)    | ❌ 超时            |
// | 递归带记忆化（V3-fix）    | O(mn)       | O(mn)     | ⚠️ 可以但不如迭代 |
// | DP 用 map（V4）          | O(mn)       | O(mn)     | ⚠️ 性能差          |
// | 标准二维数组 DP（V4-fix） | O(mn)       | O(mn)     | ✅ 清晰但空间可优化 |
//
// 【面试讲解建议】
// 1. 先说 DP 思路："每个位置的路径数 = 上方 + 左方"
// 2. 写出状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 3. 提到可以用滚动数组优化空间到 O(n)（V1）
// 4. 如果被追问，再提组合数学方法（V2）
// ──────────────────────────────────────────────────────────

func uniquePathsV3(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return uniquePathsV3(m-1, n) + uniquePathsV3(m, n-1)
}

// 【V3 点评 - 递归（无记忆化）】
// ✅ 思路正确：边界返回 1，递归求上方 + 左方
// 🔴 没有记忆化，时间复杂度 O(2^(m+n))，和 V0 一样会超时
// 🔴 大量重复计算：f(2,2) 在 f(3,3) 的递归树中被计算 2 次
// 修复方法：加 memo 缓存，变成带记忆化的递归（见 V3-fix）

// uniquePathsV3Fix 带记忆化的递归，时间 O(mn)，空间 O(mn)
func uniquePathsV3Fix(m int, n int) int {
	memo := make(map[[2]int]int)
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x == 1 || y == 1 {
			return 1
		}
		key := [2]int{x, y}
		if v, ok := memo[key]; ok {
			return v
		}
		memo[key] = dfs(x-1, y) + dfs(x, y-1)
		return memo[key]
	}
	return dfs(m, n)
}

// 【V3-fix 点评 - 带记忆化的递归】
// ✅ 时间 O(mn)，空间 O(mn)
// ✅ 用 memo 缓存已计算的结果，避免重复计算
// 🟡 递归有栈开销，面试时不如直接写迭代 DP（V1）简洁

func uniquePathsV4(m int, n int) int {
	var paths = make(map[[2]int]int)
	paths[[2]int{1, 1}] = 1
	for i := 1; i <= n; i++ { // 🟡 i 对应 n（列），命名反了，容易混淆
		for j := 1; j <= m; j++ { // 🟡 j 对应 m（行）
			if i == 1 && j == 1 {
				continue
			}
			paths[[2]int{i, j}] = paths[[2]int{i - 1, j}] + paths[[2]int{i, j - 1}]
			// 🟡 访问 [0,j] 或 [i,0] 时 map 返回 0，虽然正确但不够直观
		}
	}

	return paths[[2]int{n, m}]
}

// 【V4 点评 - DP 用 map】
// ✅ 思路正确：用 map 存储每个位置的路径数，自底向上 DP
// 🟡 用 map[[2]int]int 性能差：有哈希开销，不如二维数组或一维滚动数组
// 🟡 循环变量命名反了：i 对应 n（列），j 对应 m（行），容易混淆
// 🟡 边界处理隐式：访问越界时 map 返回 0，虽然正确但不够直观
// 修复方法：改成标准二维数组 DP 或一维滚动数组（见 V4-fix）

// uniquePathsV4Fix 标准二维数组 DP，时间 O(mn)，空间 O(mn)
func uniquePathsV4Fix(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 第一列全是 1（只能从上方来）
	}
	for j := range dp[0] {
		dp[0][j] = 1 // 第一行全是 1（只能从左方来）
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] // 上方 + 左方
		}
	}
	return dp[m-1][n-1]
}

// 【V4-fix 点评 - 标准二维数组 DP】
// ✅ 时间 O(mn)，空间 O(mn)
// ✅ 边界显式初始化，逻辑清晰
// ✅ 比 V4 的 map 版本性能更好，无哈希开销
// 🟡 空间可以进一步优化到 O(n)（见 V1 的滚动数组）

func main() {
	println("V1 (DP 滚动数组):", uniquePathsV1(3, 7)) // 28
	println("V2 (组合数学):", uniquePathsV2(3, 7))   // 28
	println("V3-fix (递归+记忆化):", uniquePathsV3Fix(3, 7)) // 28
	println("V4-fix (二维数组 DP):", uniquePathsV4Fix(3, 7)) // 28
}

// f(x,y) = f(x-1, y) + f(x, y-1)

// f(3,7) = f(2, 7) + f(3, 6)
