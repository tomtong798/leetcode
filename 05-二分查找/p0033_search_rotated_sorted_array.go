//go:build p0041

// LeetCode 第33题 - 搜索旋转排序数组 (Search in Rotated Sorted Array)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 整数数组 nums 按升序排列，数组中的值互不相同。
// 在传给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]。
// 给你旋转后的数组 nums 和一个整数 target，如果 nums 中存在这个目标值，则返回它的下标，
// 否则返回 -1。你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 【示例 1】
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
// 【示例 2】
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
//
// 【示例 3】
// 输入：nums = [1], target = 0
// 输出：-1
//
// 【数据范围】
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// nums 中每个值都独一无二，nums 必然在某个下标上旋转过
//
// 【思路提示】二分后至少有一半是有序的：判断 nums[left] <= nums[mid] 确定左半是否有序，
// 再看 target 是否落在有序那半的范围内，据此收缩边界。
//
// 题目链接：https://leetcode.cn/problems/search-in-rotated-sorted-array/
package main

import "fmt"

// searchRotated 在旋转升序数组中查找 target。
func searchRotated(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var search func(start, end int) int
	search = func(start, end int) int {
		if start == end {
			if nums[start] != target {
				return -1
			} else {
				return start
			}
		}

		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		// 经过旋转之后最后一个元素一定是小于等于第一个元素的
		// 🟡 这条注释不准确：k=0（未旋转）时 nums[end] > nums[start]，且与下方判断无关
		if nums[mid] > target {
			if nums[start] <= nums[mid] && target >= nums[start] {
				return search(start, mid-1) // ✅ 左半有序且 target 在左半范围内，向左
			} else {
				return search(mid+1, end) // 🔴 BUG 1：左半无序时（nums[start] > nums[mid]），
				//    右半 [mid..end] 有序且所有元素 >= nums[mid] > target，
				//    target 只能在左半，应该 search(start, mid-1)，这里却去了右半
				//    反例：nums=[5,1,2,3,4], target=1 → 返回 -1，正确答案 1
			}
		} else {
			if nums[start] <= nums[mid] {
				return search(mid+1, end) // ✅ 左半有序则左半全部 <= nums[mid] < target，向右
			} else {
				return search(start, mid) // 🔴 BUG 2：左半无序时右半 [mid+1..end] 有序，
				//    若 nums[mid] < target <= nums[end]，target 就在右半，
				//    应先判断 target <= nums[end] 再决定方向，这里却无条件向左
				//    另外 mid 已确认 != target，带着 mid 递归也是冗余的
				//    反例：nums=[5,1,3], target=3 → 返回 -1，正确答案 2
			}
		}

	}

	return search(0, len(nums)-1)
}

// 【V0 点评】
// 编译通过，main 里的 5 个用例恰好全过，但存在两个方向判断 bug：
//   🔴 BUG 1（第 66 行）：nums[mid] > target 且左半无序时去了右半。
//      此时右半有序且最小值是 nums[mid] > target，右半不可能有 target，应向左
//      反例：nums=[5,1,2,3,4], target=1 → 返回 -1，正确答案 1
//   🔴 BUG 2（第 72 行）：nums[mid] < target 且左半无序时无条件向左。
//      此时右半有序，若 nums[mid] < target <= nums[end]，target 就在右半
//      反例：nums=[5,1,3], target=3 → 返回 -1，正确答案 2
//   🟡 第 61 行注释不准确（k=0 未旋转时不成立）
//   🟡 测试覆盖不足：5 个用例都在幸运路径上，没覆盖"target 在旋转点后的小值区"
//   🟡 base case 只有 start == end：改成 mid-1 收缩后可能出现 start > end 的空区间，
//      需要改成 start > end 判空（见 V1），否则有空区间死循环风险
//   ✅ mid = start + (end-start)/2 防溢出写法正确
//   ✅ 左半有序时的两个分支（第 63、69 行）逻辑正确
// 核心口诀：先判断哪半有序，再判断 target 是否落在有序那半的 [min, max] 范围内

// searchRotatedV1 修复版：递归二分，时间 O(log n)，空间 O(log n)（递归栈）
func searchRotatedV1(nums []int, target int) int {
	var search func(start, end int) int
	search = func(start, end int) int {
		if start > end { // ✅ 终止条件改为 start > end，空区间直接返回 -1
			return -1
		}

		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[start] <= nums[mid] { // 左半 [start, mid] 有序
			if nums[start] <= target && target < nums[mid] {
				return search(start, mid-1) // target 在左半范围内，向左
			}
			return search(mid+1, end) // 不在，去右半
		}

		// 右半 [mid, end] 有序
		if nums[mid] < target && target <= nums[end] {
			return search(mid+1, end) // target 在右半范围内，向右
		}
		return search(start, mid-1) // 不在，去左半
	}

	return search(0, len(nums)-1)
}

// 【V1 点评 - 修复版】
// ✅ 修复 BUG 1 和 BUG 2：不再"猜方向"，统一结构——
//    先判断哪半有序，再判断 target 是否在有序那半的范围内，命中才去那边
// ✅ 终止条件 start > end：配合 mid-1 / mid+1 收缩，空区间安全返回 -1
// ✅ 标准四行判断：nums[start] <= target && target < nums[mid]（左半有序时）
//                  nums[mid] < target && target <= nums[end]（右半有序时）
// 面试推荐写法（迭代版同理：把递归改成 for left <= right 即可）

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                  | 时间     | 空间     | 面试期望       |
// |-----------------------|----------|----------|----------------|
// | 递归二分（V0 有 bug）  | O(log n) | O(log n) | ❌ 方向判断错误 |
// | 递归二分修复（V1）     | O(log n) | O(log n) | ✅ 推荐         |
// | 迭代二分（标准答案）    | O(log n) | O(1)     | ✅ 面试首选     |
//
// 【核心易错点】
// 1. 二分后至少有一半是有序的：nums[start] <= nums[mid] → 左半有序，否则右半有序
// 2. 只能对"有序的那半"做范围判断：
//    左半有序：nums[start] <= target && target < nums[mid]
//    右半有序：nums[mid] < target && target <= nums[end]
// 3. 终止条件：迭代用 left <= right；递归用 start > end 返回 -1
// 4. 测试要覆盖"target 在旋转点后小值区"的用例，如 [5,1,2,3,4] 找 1
//
// 【面试讲解建议】
// 1. 先说性质："旋转数组二分后，至少有一半是有序的"
// 2. 画图说明两段结构：[大数段 | 小数段]，旋转点在中间
// 3. 强调判断顺序：先判哪半有序，再判 target 范围
// 4. 分析复杂度：时间 O(log n)，空间 O(1)（迭代）
// ──────────────────────────────────────────────────────────

func main() {
	fmt.Println(searchRotatedV1([]int{4, 5, 6, 7, 0, 1, 2}, 6)) // 2
	fmt.Println(searchRotatedV1([]int{4, 5, 6, 7, 0, 1, 2}, 4)) // 0
	fmt.Println(searchRotatedV1([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(searchRotatedV1([]int{4, 5, 6, 7, 0, 1, 2}, 7)) // 3
	fmt.Println(searchRotatedV1([]int{1}, 0))                   // -1

	// V0 的两个反例：V0 返回 -1（错误），V1 返回正确下标
	fmt.Println(searchRotated([]int{5, 1, 2, 3, 4}, 1), searchRotatedV1([]int{5, 1, 2, 3, 4}, 1)) // -1 1
	fmt.Println(searchRotated([]int{5, 1, 3}, 3), searchRotatedV1([]int{5, 1, 3}, 3))             // -1 2
}
