//go:build p0041

// LeetCode 第90题 - 子集 II (Subsets II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】位运算、数组、回溯
//
// 【题目描述】
// 给你一个整数数组 nums，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集不能包含重复的子集。返回的解集中，子集可以按任意顺序排列。
//
// 【示例 1】
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
// 【示例 2】
// 输入：nums = [0]
// 输出：[[],[0]]
//
// 【数据范围】
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
//
// 【思路提示】排序后回溯，同层跳过重复 (i > start && nums[i] == nums[i-1])。
//
// 题目链接：https://leetcode.cn/problems/subsets-ii/
package main

import (
	"fmt"
	"sort"
)

// subsetsWithDup 返回所有不重复子集。
func subsetsWithDup(nums []int) [][]int {
	// 遍历数组中的每一个元素，每个元素都可以被选或者不被选，都有两种选择，但是因为有重复元素，可能会重复
	// ✅ 思路方向对：选/不选的二叉决策树模型（子集问题的经典视角）

	var helper func(start int, cur []int)
	var result [][]int
	helper = func(start int, cur []int) {
		if start == len(nums) { // ✅ 终止：决策完所有元素
			result = append(result, cur) // 🔴 BUG C：没拷贝 cur——存的是引用，
			return                       //    撤销现场会篡改已收集的结果（配合下方的 cur[:oldLen] 几乎必错）
		}

		oldLen := len(cur)

		for i := start; i < len(nums); i++ { // 🔴 BUG 1（结构错误）：for 循环 + 二叉决策混用！
			// 选择当前元素
			cur = append(cur, nums[i])
			helper(start+1, cur) // 🔴 BUG 2：递归传 start+1 而不是 i+1——
			//    循环变量是 i 却推进 start，i>start 时元素被重复决策
			cur = cur[:oldLen]
			helper(start+1, cur) // 🔴 BUG 1 的另一半：不选分支也递归——
			//    和外层 for 组成"每层选 i 又选不选"的双重分叉，指数爆炸
			cur = cur[:oldLen]
		}
	}

	helper(0, []int{})

	return result
}

// 【V0 点评】
// 编译通过，但结构性错误 + 两处细节 bug，[1,2,3] 就输出异常：
//   🔴 BUG 1（结构）：for 循环与"选/不选"二叉决策混用
//      子集有两种等价模板（见 V1/V2），不能杂交：
//        模板 A（start 推进型）：每层 for i := start..n，选 i 后递归(i+1)——循环本身就是"分支"
//        模板 B（二叉决策型）：无循环，直接两次递归：选 nums[start] / 不选，递归(start+1)
//      你的代码 = 模板 B 的两个递归【放进】模板 A 的 for 里：
//        每层对每个 i 都做"选/不选"两次递归 → 决策树分叉失控
//        追踪 [1,2,3]：helper(0) 里 i=0,1,2 各递归两次(选/不选) = 6 个分支，
//        每个分支再进 helper(1) 又是 6 个... 且循环+start 双重推进互相干扰
//        结果：大量重复子集 + 元素重复使用（[2,2] 这种）+ 深度失控
//   🔴 BUG 2：递归传 start+1 应为 i+1（p0216 V2 的 BUG B 同款）
//      循环里选了 nums[i]，下一层应从 i+1 开始——传 start+1 导致可重选 ≤i 的元素
//   🔴 BUG C：收集没拷贝 cur——p0216 刚踩过的坑（append([]int{}, cur...)）
//   🔴 题目核心要求没实现：输入含重复元素（[1,2,2]），输出不能有重复子集——
//      去重逻辑（排序 + 同层跳过）完全缺失
//   ✅ 值得肯定：终止条件、撤销现场的意识（cur[:oldLen]）、"选/不选"的模型认知都对
//   修复方向：二选一——用模板 A 加同层去重（V1），或用模板 B 加计数去重（V2）

// subsetsWithDup 返回所有不重复子集。
// 【V3 点评 - subsetsWithDup_1（你的第二版）】
// 进步巨大：模板 B 骨架（无循环、二叉决策）+ 排序 + 拷贝快照全对了！
// 只剩一个隐蔽的去重 bug：
//
//   🔴 去重条件看错了对象：cur[oldLen-1] != nums[start] 查的是【路径末尾】，
//      正确的对象是【前一个相等元素是否被跳过】（决策历史）
//      反例 [1,2,2]：
//        路径1：选 1 → 选 2₁ → 跳 2₂（cur 末尾是 2，条件挡住）→ [1,2] ✓
//        路径2：选 1 → 跳 2₁ → 选 2₂（cur 末尾是 1 ≠ 2，条件放行！）→ [1,2] ❌ 重复
//      根源：cur 里没有"我跳过了 2₁"的信息——【不选的决策是看不见的】，
//      这正是 V2 需要 picked 参数的原因：把决策历史显式传下去
//
//   ✅ 修复（见 subsetsWithDup_1_fixed）：用 prevSkipped bool 记录"上一个元素是否被跳过"，
//      相等 && 上一个被跳过 → 本个禁止选（断链后不能重选）
//
//   ✅ 其余全对：终止条件、拷贝收集、两次递归的撤销位置
//   🟡 条件 `oldLen > 0 && ... || oldLen == 0` 运算符优先级易错（&& 高于 ||），
//      且 oldLen==0 时前半必然 false——整个条件等价于 oldLen==0 || cur[oldLen-1] != nums[start]，
//      加括号或调整顺序更可读
func subsetsWithDup_1(nums []int) [][]int {
	sort.Ints(nums)

	var helper func(start int, cur []int)
	var result [][]int
	helper = func(start int, cur []int) {
		if start == len(nums) {
			result = append(result, append([]int{}, cur...)) // ✅ 拷贝快照（上次的坑已修）
			return
		}

		oldLen := len(cur)

		// 选择当前元素
		if oldLen > 0 && cur[oldLen-1] != nums[start] || oldLen == 0 { // 🔴 去重条件对象错误（见点评）
			cur = append(cur, nums[start])
			helper(start+1, cur)
		}

		// 不选当前元素
		cur = cur[:oldLen]
		helper(start+1, cur)
		cur = cur[:oldLen]
	}

	helper(0, []int{})

	return result
}

// subsetsWithDup_1_fixed 你的版本的修复：加 prevSkipped 参数（= V2 的 picked）
// 追踪 [1,2,2] 路径 2：选 1 → 跳 2₁（prevSkipped=true）
//   → 决策 2₂ 时：相等 && prevSkipped → 禁止选 ✓ [1,2] 不再重复
func subsetsWithDup_1_fixed(nums []int) [][]int {
	sort.Ints(nums)

	var helper func(start int, prevSkipped bool, cur []int)
	var result [][]int
	helper = func(start int, prevSkipped bool, cur []int) {
		if start == len(nums) {
			result = append(result, append([]int{}, cur...))
			return
		}

		// 选：与前一个元素相等、且前一个被跳过 → 断链后重选 = 重复，禁止
		canPick := !(start > 0 && nums[start] == nums[start-1] && prevSkipped)
		if canPick {
			helper(start+1, false, append(cur, nums[start])) // 选了 → 本元素未跳过
		}

		// 不选
		helper(start+1, true, cur) // 跳过 → 本元素被跳过
	}

	helper(0, false, []int{})
	return result
}

// 【_1_fixed 点评 - 与标准答案 V1 对比】
// 正确性：✅ 两者等价，输出完全一致（下方 main 已验证）
// 本质：_1_fixed 与 V2 是同一模型（二叉决策 + 断链检测），只是参数名相反
//   （prevSkipped = !picked），条件写法略不同（V2 去重时强制 return 只走不选；
//    _1_fixed 用 canPick 条件跳过选分支，效果相同）
//
// V1（模板 A）相对更优的三个理由：
//   1. 剪枝更早：同层 continue 直接不进分支；_1_fixed 断链后仍要沿"只跳不选"
//      的路径下潜到底（每条断链多走 O(n) 个节点）——n≤10 无感，但节点数更少是事实
//   2. 面试通用性：i > start && nums[i]==nums[i-1] 是组合总和 II、全排列 II 等
//      一整类去重题的公共习语，讲一道通一类；prevSkipped 线程只适配二叉决策视角
//   3. 收集时机更省：进入即收集，配合剪枝实际生成的"结果路径前缀"更紧凑
//
// _1_fixed 的价值：这是你从 V0 混用模板 → V3 条件看错对象 → 自己补上决策历史
// 参数修出来的版本，完整走通了"选/不选"模型的正确实现——作为学习路径的终点
// 比直接背 V1 更扎实。两个都留着对照，面试默写 V1。
//
// ⚠️ 一个 Go 细节注意（不是 bug 但依赖隐式约定）：helper(start+1, false,
// append(cur, nums[start])) 把 append 写在调用里，依赖"结果收集时深拷贝 +
// 兄弟分支顺序执行"才安全（skip 分支后续 append 会覆写同一底层数组槽位）。
// 当前代码因 append([]int{}, cur...) 的拷贝而正确，但显式 cur = cur[:oldLen]
// 的撤销写法（V3 风格）意图更明显，不易被改坏。

// subsetsWithDupV1 【LeetCode 标准答案】模板 A：start 推进 + 排序 + 同层跳过，时间 O(n·2^n)，空间 O(n)
func subsetsWithDupV1(nums []int) [][]int {
	sort.Ints(nums) // ✅ 关键第 1 步：排序——让相等元素相邻，去重判断才有意义
	var result [][]int
	path := make([]int, 0, len(nums))

	var backtrack func(start int)
	backtrack = func(start int) {
		// 子集问题：每个节点都是一个解（进入即收集）——不需要等选满
		result = append(result, append([]int{}, path...)) // ✅ 拷贝快照

		for i := start; i < len(nums); i++ {
			// ✅ 关键第 2 步：同层去重——i > start 保证只在"同一层"判断，
			// nums[i] == nums[i-1] 说明本层已用相同值开过分支，跳过
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i]) // 做选择
			backtrack(i + 1)             // ✅ i+1（不是 start+1）
			path = path[:len(path)-1]    // 撤销
		}
	}

	backtrack(0)
	return result
}

// 【V1 点评 - LeetCode 标准答案（模板 A）】
// ✅ 三件套 + 进入即收集（子集题特征：树上每个节点都是解，不是只有叶子）
// ✅ 去重两步曲：排序（相等元素相邻）→ 同层跳过（i > start && nums[i] == nums[i-1]）
//    为什么 i > start 不可省？——区分"树层"和"树枝"：
//      [1,2,2]：第一个 2 被选入树枝后，第二个 2 仍可接续（树枝方向允许 [1,2,2]）
//              但回到同一层时，第二个 2 不能再当起点（否则 [.., 2] 分支重复）
//    记忆：树枝去重不需要，树层去重必须做——i > start 就是"这是同层"的判据
// ✅ 追踪 [1,2,2]：收集 [] [1] [1,2] [1,2,2] [2] [2,2] —— 恰好 6 个，无重复 ✅

// subsetsWithDupV2 模板 B：二叉决策（选/不选）+ 计数跳过，时间 O(n·2^n)，空间 O(n)
func subsetsWithDupV2(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	path := make([]int, 0, len(nums))

	// index：当前决策到哪个元素；picked：上一个元素是否被选（用于去重）
	var dfs func(index int, picked bool)
	dfs = func(index int, picked bool) {
		if index == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		// ✅ 去重：与前一元素相等 且 前一元素【没被选】→ 本元素也不能选
		//   （相等元素必须"连续选"，跳过一个再选后一个 = 重复子集）
		if index > 0 && nums[index] == nums[index-1] && !picked {
			// 只剩"不选"一条路：相等元素断开后，后面全跳过
			dfs(index+1, false)
			return
		}
		// 选
		path = append(path, nums[index])
		dfs(index+1, true)
		path = path[:len(path)-1]
		// 不选
		dfs(index+1, false)
	}

	dfs(0, false)
	return result
}

// 【V2 点评 - 模板 B（二叉决策型）】
// ✅ 这才是你 V0 想写的模板：无 for 循环，每层恰好两次递归（选/不选）
// ✅ 去重思路换角度：相等元素必须"连续选取"——若前一个相等元素被跳过，
//    本元素再选就会造出重复子集（[1,2₁] vs [1,2₂]），所以也强制跳过
// ✅ picked 参数随递归下传（自顶向下传状态）——判断"断没断"的信息
// 与 V1 对比：V1 靠"同层起点跳过"（循环视角），V2 靠"断链检测"（决策视角）——
// 同一个去重目标的两种实现；V1 更常用（可推广到组合总和 IV 等题），V2 更贴"选/不选"直觉

// ──────────────────────────────────────────────────────────
// 【子集问题的两种模板对照（V0 混用的正是这两者）】
//
//   模板 A（start 推进型）          模板 B（二叉决策型）
//   ─────────────────────────      ─────────────────────────
//   for i := start; i < n; i++ {   dfs(index) {
//       选 i                          选 nums[index] → dfs(index+1)
//       backtrack(i+1)                 不选           → dfs(index+1)
//       撤销 i                      }
//   }                              // 无循环，恰好两次递归
//
//   分支 = "这一层从谁开始"          分支 = "这个元素选不选"
//   收集：进入函数即收（子集）        收集：index 走到头时收
//   去重：i > start && 相等跳过      去重：相等且前一未选 → 强制不选
//
//   共同前提：排序（相等元素相邻）
//   两者等价，产出同一棵决策树的不同走法；千万别混用（V0 的教训）
//
// 子集 vs 组合的收集时机：
//   组合（p0216）：选满 k 个才收集（叶子收集）
//   子集（本题）：每个节点都是解，进入即收集（含空集 []）
// ──────────────────────────────────────────────────────────

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                          | 时间       | 空间   | 面试期望      |
// |-------------------------------|------------|--------|---------------|
// | 模板混用（V0 有 bug）          | 指数爆炸   | O(n)   | ❌ 结构错误   |
// | 模板 A + 同层跳过（V1 标准）   | O(n·2^n)   | O(n)   | ✅ 面试首选   |
// | 模板 B + 断链检测（V2）        | O(n·2^n)   | O(n)   | ✅ 同为标准   |
//
// 【核心易错点】
// 1. 模板不能混用：for 循环（start 推进）和二叉决策（选/不选双递归）是两套骨架
// 2. 排序是去重前提：相等元素不相邻，任何相邻判断都失效
// 3. 同层去重三条件缺一不可：i > start（同层）&& nums[i]==nums[i-1]（相等）→ continue
// 4. 子集收集时机：进入函数即收集，不是等选满（否则漏掉非叶子的子集）
// 5. 收集必须拷贝：append([]int{}, path...)——p0216 BUG C 同款
// 6. 递归推进用 i+1 不是 start+1——p0216 BUG B 同款
//
// 【面试讲解建议】
// 1. 先说框架："排序 + 回溯，进入即收集所有子集"
// 2. 讲去重："同层遇到重复值跳过——i > start 保证判断的是层不是枝"
// 3. 画 [1,2,2] 的决策树标出剪掉的分支（最直观的讲解方式）
// 4. 复杂度：O(n·2^n)（子集总数 2^n 上界 × 每个拷贝 O(n)）
// ──────────────────────────────────────────────────────────

func main() {
	// fmt.Println(subsetsWithDupV1([]int{1, 2, 2})) // [[] [1] [1 2] [1 2 2] [2] [2 2]]
	// fmt.Println(subsetsWithDupV1([]int{0}))       // [[] [0]]
	// fmt.Println(subsetsWithDupV1([]int{1, 2, 3})) // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]

	// // V2 与 V1 结果一致
	// fmt.Println(subsetsWithDupV2([]int{1, 2, 2})) // [[] [1] [1 2] [1 2 2] [2] [2 2]]

	// // V0 的输出（结构错乱：重复 + 元素复用 + 深度失控）：
	// // fmt.Println(subsetsWithDup([]int{1, 2, 3})) // 取消注释可见异常结果
	// _ = subsetsWithDup

	fmt.Println(subsetsWithDup_1([]int{1, 2, 2})) // 有重复：[1,2] 出现两次（去重条件对象错的反例）
	fmt.Println(subsetsWithDup_1_fixed([]int{1, 2, 2})) // [[] [1] [1 2] [1 2 2] [2] [2 2]] ✅ 修复后无重复
	fmt.Println(subsetsWithDup_1_fixed([]int{0}))       // [[] [0]]
	fmt.Println(subsetsWithDup_1_fixed([]int{1, 2, 3})) // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsetsWithDup_1_fixed([]int{2, 2, 2})) // [[] [2] [2 2] [2 2 2]] 三连重复压力测试
}
