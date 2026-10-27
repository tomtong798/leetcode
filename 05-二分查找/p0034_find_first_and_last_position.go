//go:build p0041

// LeetCode 第34题 - 在排序数组中查找元素的第一个和最后一个位置 (Find First and Last Position of Element in Sorted Array)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、二分查找
//
// 【题目描述】
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
// 请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
// 【示例 1】
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
//
// 【示例 2】
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
//
// 【示例 3】
// 输入：nums = [], target = 0
// 输出：[-1,-1]
//
// 【数据范围】
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
// nums 是一个非递减数组
//
// 【思路提示】两次二分：lower_bound 找第一个 >= target 的位置，
// 再找第一个 > target 的位置减一；校验区间合法性。
// 熟练掌握 lower_bound / upper_bound 两个模板即可通吃此类题。
//
// 题目链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

import "fmt"

// searchRange 返回 target 在 nums 中的首尾下标。
func searchRange(nums []int, target int) []int {
	// 起始位置：当前元素 == target && 上一个元素 < target
	// 终止位置：当前元素 == target && 下一个元素 > target

	var seachLeft func(start, end int) int  // 🟡 拼写错误：seach → search（下同）
	var seachRight func(start, end int) int // 🟡 拼写错误
	seachLeft = func(start, end int) int {
		if start > end {
			return -1
		}
		mid := start + (end-start)/2
		if mid-1 < 0 { // 🔴 BUG 1：mid==0 且 nums[0]!=target 时直接返回 -1，
			if nums[mid] == target { //    丢掉了右半边的候选（此窗口可能有两个元素）
				return mid //    反例：nums=[1,3,5,6,7], target=3
			} //      seachLeft(0,4)→mid=2(5>3)→seachLeft(0,1)→mid=0, nums[0]=1≠3
			return -1 //      → 返回 -1 ❌，正确答案 1
		}

		if nums[mid] < target {
			return seachLeft(mid+1, end)
		} else if nums[mid] > target {
			return seachLeft(start, mid-1)
		} else {
			if nums[mid-1] < target {
				return mid // ✅ mid 是第一个 target
			} else {
				return seachLeft(start, mid-1) // ✅ 左边还有 target，继续向左
			}
		}
		return -1 // 🟡 死代码：所有分支都已 return，永远执行不到
	}

	seachRight = func(start, end int) int {
		if start > end {
			return -1
		}
		mid := start + (end-start)/2
		if mid+1 > len(nums)-1 { // ✅ mid==len-1 只在 start==end 时出现，单元素窗口返回 -1 正确
			if nums[mid] == target {
				return mid
			}
			return -1
		}

		if nums[mid] < target {
			return seachRight(mid+1, end)
		} else if nums[mid] > target {
			return seachRight(start, mid-1)
		} else {
			if nums[mid+1] > target {
				return mid // ✅ mid 是最后一个 target
			} else {
				return seachRight(start, mid-1) // 🔴 BUG 2：nums[mid+1]==target 时，
				//    最后一个 target 在右边，应该 seachRight(mid+1, end)，
				//    却递归到了左边，永远找不到
				//    反例：nums=[5,7,7,7,10], target=7 → 返回 -1 ❌，正确答案 3
				//    反例：nums=[8,8], target=8 → 返回 -1 ❌，正确答案 1
			}
		}
		return -1 // 🟡 死代码
	}

	return []int{seachLeft(0, len(nums)-1), seachRight(0, len(nums)-1)}
}

// 【V0 点评】
// 编译通过，main 里的 3 个用例恰好全过，但存在两个致命 bug：
//   🔴 BUG 1（seachLeft 特判分支）：mid==0 且 nums[0]!=target 时直接返回 -1，
//      丢掉了同一窗口右边的候选
//      反例：nums=[1,3,5,6,7], target=3 → 返回 [-1,-1]，正确 [1,1]
//           nums=[5,7], target=7 → 返回 [-1,...]，正确 [1,1]
//   🔴 BUG 2（seachRight else 分支）：nums[mid+1]==target 时向左递归，
//      但最后一个 target 一定在右边，永远找不到
//      反例：nums=[5,7,7,7,10], target=7 → 返回 [.., -1]，正确 [1,3]
//           nums=[8,8], target=8 → 返回 [0,-1]，正确 [0,1]
//   🟡 拼写错误：seachLeft/seachRight 应为 searchLeft/searchRight
//   🟡 两处死代码：闭包末尾的 return -1 永远执行不到
//   🟡 特判 mid-1<0 / mid+1>len-1 的写法脆弱：BUG 1 正是特判引入的
//      标准做法用 lower_bound 模板，不需要任何邻居特判
//   🟡 测试覆盖不足：没测"重复元素跨越 mid"和"target 在窗口第二位"的用例
//   ✅ 终止条件 start > end 正确
//   ✅ mid = start + (end-start)/2 防溢出写法正确
//   ✅ seachLeft 主体逻辑（找第一个）正确；seachRight 的 nums[mid+1]>target 判断本身正确

func searchRange_Fix(nums []int, target int) []int {
	// 起始位置：当前元素 == target && 上一个元素 < target
	// 终止位置：当前元素 == target && 下一个元素 > target

	var seachLeft func(start, end int) int
	var seachRight func(start, end int) int
	seachLeft = func(start, end int) int {
		if start > end {
			return -1
		}
		mid := start + (end-start)/2
		if mid-1 < 0 {
			if nums[mid] == target {
				return mid
			} else {
				return seachLeft(mid+1, end)
			}
		}

		if nums[mid] < target {
			return seachLeft(mid+1, end)
		} else if nums[mid] > target {
			return seachLeft(start, mid-1)
		} else {
			if nums[mid-1] < target {
				return mid
			} else {
				return seachLeft(start, mid-1)
			}
		}
	}

	seachRight = func(start, end int) int {
		if start > end {
			return -1
		}
		mid := start + (end-start)/2
		if mid+1 > len(nums)-1 {
			if nums[mid] == target {
				return mid
			} else {
				return seachRight(start, mid-1)
			}
		}

		if nums[mid] < target {
			return seachRight(mid+1, end)
		} else if nums[mid] > target {
			return seachRight(start, mid-1)
		} else {
			if nums[mid+1] > target {
				return mid
			} else {
				return seachRight(mid+1, end)
			}
		}
	}

	return []int{seachLeft(0, len(nums)-1), seachRight(0, len(nums)-1)}
}

// searchRangeV1 修复版：lower_bound 模板，时间 O(log n)，空间 O(1)
func searchRangeV1(nums []int, target int) []int {
	// lowerBound 返回第一个 >= t 的下标（不存在则返回 len(nums)）
	var lowerBound func(t int) int
	lowerBound = func(t int) int {
		left, right := 0, len(nums) // right 是开区间边界
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < t {
				left = mid + 1 // [mid+1, right)
			} else {
				right = mid // [left, mid)
			}
		}
		return left
	}

	start := lowerBound(target) // 第一个 >= target 的位置
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1} // target 不存在
	}
	end := lowerBound(target+1) - 1 // 第一个 > target 的位置再减 1
	return []int{start, end}
}

// 【V1 点评 - 修复版】
// ✅ 修复 BUG 1/BUG 2：lower_bound 模板不需要任何邻居特判，边界天然正确
// ✅ 只写一个辅助函数，两次调用：first = lower_bound(target)，last = lower_bound(target+1)-1
// ✅ 先校验 start 越界/不等，不存在时统一返回 [-1,-1]，避免两次无效二分
// ✅ 迭代写法，时间 O(log n)，空间 O(1)
// 面试推荐写法：lower_bound/upper_bound 两个模板可以通吃所有"找边界"的二分题

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                     | 时间     | 空间   | 面试期望       |
// |--------------------------|----------|--------|----------------|
// | 邻居判断法（V0 有 bug）   | O(log n) | O(log n) | ❌ 两个方向 bug |
// | 递归邻居判断修复          | O(log n) | O(log n) | ⚠️ 可以但特判多 |
// | lower_bound 模板（V1）    | O(log n) | O(1)   | ✅ 面试首选     |
//
// 【核心易错点】
// 1. 找"最后一个 target"时，nums[mid+1]==target 必须向右递归（V0 BUG 2 的根源）
// 2. 邻居判断（nums[mid-1]<target）需要特判 mid 在数组边界，特判极易出错
// 3. 更稳的思路：不判断邻居，判断"收缩方向"——
//    lower_bound(x)：nums[mid] < x → 收右；否则收左。循环结束时 left 就是答案
// 4. last = lower_bound(target+1) - 1 是常用技巧，注意先确认 target 存在再减
//
// 【面试讲解建议】
// 1. 先说思路："两次二分，first = lower_bound(target)，last = lower_bound(target+1)-1"
// 2. 强调 lower_bound 的定义：第一个 >= target 的位置（左闭右开区间）
// 3. 提到为什么不判断邻居：边界特判容易出错，模板更稳
// 4. 分析复杂度：时间 O(log n)，空间 O(1)
// ──────────────────────────────────────────────────────────

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3,4] ✅ V0 碰巧对
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1,-1] ✅
	fmt.Println(searchRange([]int{}, 0))                  // [-1,-1] ✅

	// V0 的反例：V0 全错，V1 全对
	fmt.Println(searchRange_Fix([]int{1, 3, 5, 6, 7}, 3), searchRangeV1([]int{1, 3, 5, 6, 7}, 3))   // [-1,-1] [1,1]
	fmt.Println(searchRange_Fix([]int{5, 7, 7, 7, 10}, 7), searchRangeV1([]int{5, 7, 7, 7, 10}, 7)) // [1,-1] [1,3]
	fmt.Println(searchRange_Fix([]int{8, 8}, 8), searchRangeV1([]int{8, 8}, 8))                     // [0,-1] [0,1]
}

// 标准二分查找写法
func searchRange_V3(nums []int, target int) []int {
	left := findFirst(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLast(nums, target)
	return []int{left, right}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}
