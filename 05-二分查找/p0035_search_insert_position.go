//go:build p0041

// LeetCode 第35题 - 搜索插入位置 (Search Insert Position)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 【示例 1】
// 输入：nums = [1,3,5,6], target = 5
// 输出：2
//
// 【示例 2】
// 输入：nums = [1,3,5,6], target = 2
// 输出：1
//
// 【示例 3】
// 输入：nums = [1,3,5,6], target = 7
// 输出：4
//
// 【数据范围】
// 1 <= nums.length <= 10^4
// -10^4 <= nums[i] <= 10^4
// nums 为无重复元素的升序排列数组
//
// 【思路提示】本质是"第一个 >= target 的下标"（lower_bound），
// 左闭右开写法：while left < right，mid 偏左，nums[mid] < target 时 left = mid+1 否则 right = mid。
//
// 题目链接：https://leetcode.cn/problems/search-insert-position/
package main

import "fmt"

// searchInsert 返回 target 的下标或应插入位置。
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 🟡 空数组返回 -1 有争议：插入位置语义下应返回 0（LeetCode 保证 len>=1，不影响）
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if mid+1 < len(nums) && target < nums[mid+1] || mid+1 == len(nums) { // 🟡 && 优先级高于 ||，没加括号可读性差
				return mid + 1 // ✅ nums[mid] < target < nums[mid+1]，插在 mid+1，正确
			}
			left = mid + 1
		} else {
			if mid-1 > 0 && target > nums[mid-1] || mid-1 == -1 { // 🔴 BUG：mid-1 > 0 应为 >= 0
				return mid - 1 // 🔴 BUG 1：应 return mid，不是 mid-1！
				//    nums[mid-1] < target < nums[mid] 时，插入位置是 mid
				//    反例：nums=[2,4,6,8], target=7 → 返回 2 ❌，正确 3
			}
			right = mid - 1
		}
	}

	return 0 // 🟡 死代码：各分支的提前返回覆盖了所有路径，循环永远不会正常退出
}

// 【V0 点评】
// 编译通过，main 里的 3 个用例恰好全过，但有两个方向性 bug：
//   🔴 BUG 1（else 分支）：nums[mid-1] < target < nums[mid] 时 return mid-1，
//      插入位置应该是 mid！mid-1 会把 target 插到 nums[mid-1] 前面
//      反例：nums=[2,4,6,8], target=7 → 返回 2 ❌，正确 3
//           nums=[1,3,5,6,8,10], target=7 → 返回 3 ❌，正确 4
//   🔴 BUG 2（mid-1 == -1 分支）：mid==0 且 nums[0] > target 时 return -1，
//      应该 return 0（插在最前面）
//      反例：nums=[1,3,5,6], target=0 → 返回 -1 ❌，正确 0
//           nums=[5], target=3 → 返回 -1 ❌，正确 0
//   🟡 mid-1 > 0 应为 mid-1 >= 0（边界差一，当前被 BUG 1 掩盖）
//   🟡 a && b || c 没加括号，可读性差（Go 中 && 优先级高于 ||）
//   🟡 末尾 return 0 是死代码
//   🟡 测试覆盖不足：没测 target < nums[0] 和"target 落在右半窗口"的插入
//   ✅ nums[mid] < target 分支（return mid+1）逻辑正确
//   ✅ mid = left + (right-left)/2 防溢出写法正确
// 根源：用"提前返回 + 邻居判断"的写法，边界特判太多容易错。
// 标准做法是 lower_bound：判断收缩方向而不是判断邻居，天然无边界问题。

// searchInsertV1 修复版：lower_bound 模板，时间 O(log n)，空间 O(1)
func searchInsertV1(nums []int, target int) int {
	left, right := 0, len(nums) // ✅ 左闭右开：right 初始为 len(nums)，覆盖"插到末尾"
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // ✅ nums[mid] < target：答案在右半，收缩左边界
		} else {
			right = mid // ✅ nums[mid] >= target：mid 可能是答案，收缩右边界但保留 mid
		}
	}
	return left // ✅ 循环结束 left == right，就是第一个 >= target 的位置
}

// 【V1 点评 - 修复版（lower_bound 模板）】
// ✅ 修复 BUG 1/BUG 2：不判断邻居、不提前返回，只判断收缩方向，边界天然正确
// ✅ 三个反例全部通过：[1,3,5,6]+0→0、[2,4,6,8]+7→3、[5]+3→0
// ✅ 代码只有 8 行，没有任何特判
// ✅ 时间 O(log n)，空间 O(1)
// 面试推荐写法：这就是 lower_bound（第一个 >= target 的位置），
// 与 p0034 的 lowerBound 是同一个模板，掌握后可通吃所有插入位置/边界二分题

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                     | 时间     | 空间   | 面试期望       |
// |--------------------------|----------|--------|----------------|
// | 邻居判断法（V0 有 bug）   | O(log n) | O(1)   | ❌ 两个方向 bug |
// | lower_bound 模板（V1）    | O(log n) | O(1)   | ✅ 面试首选     |
//
// 【核心易错点】
// 1. nums[mid-1] < target < nums[mid] 时，插入位置是 mid，不是 mid-1（V0 BUG 1）
// 2. target < nums[0] 时，插入位置是 0，不是 -1（V0 BUG 2）
// 3. 邻居判断 + 提前返回的写法边界特判太多，极易出错；
//    lower_bound 只判断收缩方向（nums[mid] < target → 左边界收缩），无任何特判
// 4. lower_bound 返回值本身就是"找到的下标"或"应插入位置"，两种情况统一
//
// 【面试讲解建议】
// 1. 先说本质："这就是 lower_bound——找第一个 >= target 的位置"
// 2. 强调左闭右开 [left, right) 和循环条件 left < right 的配合
// 3. 提到为什么 return left：循环不变量是"答案始终在 [left, right) 内"
// 4. 分析复杂度：时间 O(log n)，空间 O(1)
// ──────────────────────────────────────────────────────────

// 左闭右闭
func searchInsertV2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 循环结束时，left 就是插入位置
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2 ✅
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1 ✅
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4 ✅

	// V0 的反例：V0 全错，V1 全对
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0), searchInsertV1([]int{1, 3, 5, 6}, 0)) // -1 0
	fmt.Println(searchInsert([]int{2, 4, 6, 8}, 7), searchInsertV1([]int{2, 4, 6, 8}, 7)) // 2 3
	fmt.Println(searchInsert([]int{5}, 3), searchInsertV1([]int{5}, 3))                   // -1 0
}
