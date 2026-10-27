//go:build p0041

// LeetCode 第74题 - 搜索二维矩阵 (Search a 2D Matrix)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、二分查找、矩阵
//
// 【题目描述】
// 给你一个满足下述两条属性的 m x n 整数矩阵：
//   - 每行中的整数从左到右按非严格递增顺序排列
//   - 每行的第一个整数大于前一行的最后一个整数
//
// 给你一个整数 target，如果 target 在矩阵中返回 true，否则返回 false。
//
// 【示例 1】
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
//
// 【示例 2】
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
//
// 【数据范围】
// m == matrix.length, n == matrix[i].length
// 1 <= m, n <= 100
//
// 【思路提示】整体可拉平成一维有序数组，二分下标 mid 映射到 matrix[mid/n][mid%n]。
//
// 题目链接：https://leetcode.cn/problems/search-a-2d-matrix/
package main

import "fmt"

// searchMatrix74 在二维矩阵中查找 target。
func searchMatrix74(matrix [][]int, target int) bool {
	// 先比价当前数字是可能在哪几行
	// 然后遍历每一行，二分查找
	// 🟡 注释错别字："比价"应为"比较"
	// 🟡 思路可优化：由于"每行第一个整数大于前一行的最后一个整数"，
	//    矩阵可拉平成一维有序数组，只需一次全局二分 O(log(mn))，
	//    逐行二分是 O(m·log n)，没用上第二个性质（见 V1）

	var binarySearch func(nums []int, target int) bool
	binarySearch = func(nums []int, target int) bool {
		if len(nums) == 0 {
			return false
		}

		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				return true
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		return false
	} // ✅ 标准三分支二分（找存在性），终止条件、收缩、防溢出全对

	for _, row := range matrix {
		if row[0] > target {
			continue // ✅ 行首 > target，该行不可能有，跳过
		}

		if binarySearch(row, target) {
			return true
		}
	} // 🟡 只用了"行内有序"，没用"行间有序"：row[0] > target 后面所有行更不可能，
	//    应该 break 而不是 continue；且可先定位唯一候选行再二分（见 V1）

	return false
}

// 【V0 点评】
// 编译通过，结果正确，但算法没用上矩阵的第二个性质：
//   ✅ binarySearch 是标准的三分支查找（找存在性）：终止 left<=right、收缩 mid±1、防溢出 mid 全对
//   ✅ row[0] > target 跳过该行的剪枝方向正确
//   ✅ 结果正确：main 两个用例 + 边界（target 在首行/末行/小于全部/大于全部）都通过
//   🟡 算法复杂度 O(m·log n)：逐行二分，只用了"行内有序"
//      题目还有"每行第一个数 > 前一行最后一个数"——整个矩阵可拉平成一维有序数组
//      标准做法一次全局二分 O(log(mn))（见 V1）
//   🟡 row[0] > target 用 continue：后面所有行的行首更大，更不可能，直接 break 更好
//      （该剪枝不改变正确性，只影响提前退出的时机）
//   🟡 应该先二分定位"唯一候选行"（最后一个 row[0] <= target 的行），只对那一行二分：
//      找行 O(log m) + 行内找 O(log n) = O(log m + log n)，同样是标准答案
//   🟡 注释错别字："比价" → "比较"
//   🟡 测试建议补充：target 小于 matrix[0][0]、大于 matrix[m-1][n-1]、单行单列矩阵

// searchMatrix74V1 【LeetCode 标准答案】拉平一维 + 全局二分，时间 O(log(mn))，空间 O(1)
func searchMatrix74V1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1 // ✅ 把矩阵当作长度 m*n 的一维有序数组
	for left <= right {
		mid := left + (right-left)/2
		val := matrix[mid/n][mid%n] // ✅ 一维下标 → 二维坐标：行 = mid/n，列 = mid%n
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 【V1 点评 - LeetCode 标准答案】
// ✅ 核心：两个性质合起来 = 整个矩阵按行优先拉平后是一维升序数组
//    [[1,3,5,7],[10,11,16,20],[23,30,34,60]] 拉平后：
//    [1,3,5,7,10,11,16,20,23,30,34,60] —— 天然有序！
// ✅ 下标映射：一维 mid → matrix[mid/n][mid%n]
//    例：m=3,n=4，mid=7 → matrix[1][3]=20（第 7 个元素，1 行 3 列）
// ✅ 不真正展开数组，只是用除法/取模做坐标换算，空间 O(1)
// ✅ 一次二分 O(log(mn)) = O(log m + log n)，优于逐行二分的 O(m·log n)
// 追踪 V1 找 target=3（m=3,n=4，共 12 个元素）：
//   left=0,right=11,mid=5 → matrix[1][1]=11 > 3 → right=4
//   left=0,right=4, mid=2 → matrix[0][2]=5  > 3 → right=1
//   left=0,right=1, mid=0 → matrix[0][0]=1  < 3 → left=1
//   left=1,right=1, mid=1 → matrix[0][1]=3  == 3 → true ✅

// searchMatrix74V2 两次二分：先定位行再行内查找，时间 O(log m + log n)，空间 O(1)
func searchMatrix74V2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	// 第一次二分：找最后一个 row[0] <= target 的行（ans 记录法）
	row := -1
	top, bottom := 0, m-1
	for top <= bottom {
		mid := top + (bottom-top)/2
		if matrix[mid][0] <= target {
			row = mid      // mid 行可能是候选，记录
			top = mid + 1  // 尝试更靠后的行
		} else {
			bottom = mid - 1
		}
	}
	if row == -1 {
		return false // target 比所有行首都小
	}

	// 第二次二分：在该行内找 target
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 【V2 点评 - 两次二分法】
// ✅ 第一次二分用 ans 记录法找"最后一个行首 <= target 的行"（与 p0069 同一模板！）
//    target 只可能出现在这一行：更前的行最大值 < target，更后的行最小值 > target
// ✅ 第二次二分行内三分支查找
// ✅ O(log m + log n) = O(log(mn))，与 V1 同阶
// 对比 V1：V1 一次二分更简洁；V2 两次二分思路更"分层"（先定位区域再精找），
// 两者都是标准答案，面试任选其一，讲清楚下标映射/行定位逻辑即可

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                       | 时间           | 空间   | 面试期望      |
// |----------------------------|----------------|--------|---------------|
// | 逐行二分（V0）              | O(m·log n)     | O(1)   | ⚠️ 正确但没用上性质2 |
// | 拉平一维全局二分（V1 标准答案）| O(log(mn))   | O(1)   | ✅ 面试首选   |
// | 两次二分定位行（V2）         | O(log m+log n) | O(1)   | ✅ 同为标准答案 |
//
// 【核心易错点】
// 1. 关键洞察：两个性质合并 = 矩阵可拉平成一维有序数组；只用性质1会退化成 O(m·log n)
// 2. 一维转二维映射：行 = mid/n，列 = mid%n（n 是列数，别除反成 m）
// 3. 找候选行用 ans 记录法：找"最后一个行首 <= target 的行"，不是第一个
// 4. 边界：target < matrix[0][0] 时 row=-1；target > matrix[m-1][n-1] 时行内找不到，都返回 false
//
// 【面试讲解建议】
// 1. 先说洞察："行首大于前行末尾 → 整个矩阵拉平就是一维有序数组"
// 2. 讲映射："mid/n 是行，mid%n 是列"，举一个具体 mid 的例子
// 3. 提到备选方案："也可以两次二分，先 ans 记录法定位行，再行内查找"
// 4. 分析复杂度：时间 O(log(mn))，空间 O(1)
// ──────────────────────────────────────────────────────────

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix74(matrix, 3))  // true
	fmt.Println(searchMatrix74(matrix, 13)) // false

	// V0/V1/V2 结果一致
	fmt.Println(searchMatrix74(matrix, 1), searchMatrix74V1(matrix, 1), searchMatrix74V2(matrix, 1))         // true true true（首元素）
	fmt.Println(searchMatrix74(matrix, 60), searchMatrix74V1(matrix, 60), searchMatrix74V2(matrix, 60))       // true true true（末元素）
	fmt.Println(searchMatrix74(matrix, 0), searchMatrix74V1(matrix, 0), searchMatrix74V2(matrix, 0))          // false false false（小于全部）
	fmt.Println(searchMatrix74(matrix, 100), searchMatrix74V1(matrix, 100), searchMatrix74V2(matrix, 100))    // false false false（大于全部）
	fmt.Println(searchMatrix74V1([][]int{{1}}, 1), searchMatrix74V2([][]int{{1}}, 2))                         // true false（单行单列）
}
