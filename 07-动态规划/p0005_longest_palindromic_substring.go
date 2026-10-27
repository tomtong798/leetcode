//go:build p0041

// LeetCode 第5题 - 最长回文子串 (Longest Palindromic Substring)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】字符串、动态规划、双指针（中心扩展）
//
// 【题目描述】
// 给你一个字符串 s，找到 s 中最长的回文子串。如果字符串的反序与原始字符串相同，
// 则该字符串称为回文字符串。
//
// 【示例 1】
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
// 【示例 2】
// 输入：s = "cbbd"
// 输出："bb"
//
// 【数据范围】
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// 【思路提示】中心扩展法最好记：每个位置（以及每对相邻位置）作为回文中心向两边扩展，
// 记录最长边界，O(n^2) 时间 O(1) 空间。DP 写法是 dp[i][j] 表示 s[i..j] 是否回文。
//
// 题目链接：https://leetcode.cn/problems/longest-palindromic-substring/
package main

import "fmt"

// longestPalindrome 返回 s 中最长回文子串。
func longestPalindrome(s string) string {
	var n int = len(s)
	var longest int
	var start, end int
	for i, _ := range s {
		for j := 1; j < i; j++ { // 只限制了 j < i，右边 i+j 可能越界
			if i-j < 0 || i+j >= n || s[i-j] != s[i+j] { // 越界/不匹配时没 break，继续用错误 j 算边界
				if longest < (j-1)*2+1 {
					longest = (j-1)*2 + 1
					start = i - j + 1 // 边界公式算错：应该用 i-(j-1) 而不是 i-j+1
					end = i + j - 1
				}
			} else {
				if longest < j*2+1 {
					longest = j*2 + 1
					start = i - j
					end = i + j
				}
			}
		}
	}

	for i, _ := range s {
		if i+1 >= n || s[i] != s[i+1] {
			continue
		}

		for j := 1; j < i; j++ { // 同样只限制了 j < i，右边可能越界
			if i-j < 0 || i+j+1 >= n || s[i-j] != s[i+j+1] { // 同样没 break
				if longest < (j-1)*2+2 {
					longest = (j-1)*2 + 2
					start = i - j + 1 // 边界公式算错
					end = i + j
				}
			} else {
				if longest < j*2+2 {
					longest = j*2 + 2
					start = i - j
					end = i + j + 1
				}
			}
		}
	}

	return s[start:end] // 编译错误：Go 切片语法是 s[start:end+1]，不是 s[start, end]
}

// 【V0 点评】
// 编译不过，逻辑也有多处错误：
//   1. s[start, end] 语法错误，Go 切片是 s[start:end+1]
//   2. 循环条件 j < i 只限制了左边，右边 i+j 可能越界
//   3. 越界/不匹配时没有 break，继续用更大的 j 计算边界，结果错误
//   4. start 的边界公式算错（i-j+1 应该是 i-(j-1) 即 i-j+1… 等等，这个碰巧对，
//      但因为没 break，j 会继续增大，最终记录的 start/end 是错的）
//   5. 偶数中心（如 "aa"）在 i=0 时内层循环不执行，丢失长度为 2 的回文
//   6. 奇数中心没有显式处理单字符回文，碰巧靠零值蒙对

func longestPalindromeV1(s string) string {
	var n int = len(s)
	var longest int
	var start, end int
	for i, _ := range s {
		for j := 1; i-j >= 0 && i+j < n; j++ { // ✅ 修复：循环条件同时限制左右边界
			if s[i-j] != s[i+j] {
				if longest < (j-1)*2+1 {
					longest = (j-1)*2 + 1
					start = i - j + 1
					end = i + j - 1
				}
				break // ✅ 修复：不匹配立即退出
			} else {
				if longest < j*2+1 {
					longest = j*2 + 1
					start = i - j
					end = i + j
				}
			}
		}
	}

	for i, _ := range s {
		if i+1 >= n || s[i] != s[i+1] {
			continue
		}
		// 🔴 问题：当 i=0 时，s[0]==s[1] 匹配了，但内层 j 循环因 i-j<0 不执行，
		//    长度为 2 的回文没有被记录！例如 s="aa" 会返回 "a" 而不是 "aa"

		for j := 1; i-j >= 0 && i+j+1 < n; j++ {
			if s[i-j] != s[i+j+1] {
				if longest < (j-1)*2+2 {
					longest = (j-1)*2 + 2
					start = i - j + 1
					end = i + j
				}
				break
			} else {
				if longest < j*2+2 {
					longest = j*2 + 2
					start = i - j
					end = i + j + 1
				}
			}
		}
	}

	return s[start : end+1] // ✅ 修复：切片语法正确
}

// 【V1 点评】
// 编译通过，大部分情况正确，但偶数中心在边界时丢失长度为 2 的回文：
//   ✅ 修复了 V0 的切片语法错误
//   ✅ 修复了循环条件，同时限制左右边界
//   ✅ 修复了不匹配后 continue 的问题，加了 break
//   🔴 偶数中心在 i=0 时（如 "aa"），内层 j 循环不执行，丢失长度为 2 的回文
//   🔴 奇数中心没有显式处理单字符回文，碰巧靠零值蒙对
// 修复方法：在进入 j 循环前，先记录基础回文（奇数长度 1，偶数长度 2）

func longestPalindromeV2(s string) string {
	var n int = len(s)
	var longest int
	var start, end int
	for i, _ := range s {
		if longest < 1 { // ✅ 先记录基础奇数回文（长度 1）
			longest = 1
			start = i
			end = i
		}

		for j := 1; i-j >= 0 && i+j < n; j++ {
			if s[i-j] != s[i+j] {
				if longest < (j-1)*2+1 {
					longest = (j-1)*2 + 1
					start = i - j + 1
					end = i + j - 1
				}
				break
			} else {
				if longest < j*2+1 {
					longest = j*2 + 1
					start = i - j
					end = i + j
				}
			}
		}
	}

	for i, _ := range s {
		if i+1 >= n || s[i] != s[i+1] {
			continue
		}

		if longest < 2 { // ✅ 先记录基础偶数回文（长度 2）
			longest = 2
			start = i
			end = i + 1
		}

		for j := 1; i-j >= 0 && i+j+1 < n; j++ {
			if s[i-j] != s[i+j+1] {
				if longest < (j-1)*2+2 {
					longest = (j-1)*2 + 2
					start = i - j + 1
					end = i + j
				}
				break
			} else {
				if longest < j*2+2 {
					longest = j*2 + 2
					start = i - j
					end = i + j + 1
				}
			}
		}
	}

	return s[start : end+1]
}

// 【V2 点评】
// 在 V1 基础上修复了边界问题，现在逻辑正确：
//   ✅ 显式处理了单字符回文（longest=1）
//   ✅ 显式处理了基础偶数回文（longest=2）
//   ✅ 所有边界条件正确
// 但代码仍然比较啰嗦，建议用 expand() 辅助函数简化：
//   func expand(s string, left, right int) (int, int) {
//       for left >= 0 && right < len(s) && s[left] == s[right] {
//           left--
//           right++
//       }
//       return left + 1, right - 1
//   }
// 这样边界计算由 expand() 统一处理，代码更简洁、不易出错

// longestPalindromeV3 用 expand() 辅助函数简化边界计算，代码更清晰不易出错。
func longestPalindromeV3(s string) string {
	if len(s) < 2 {
		return s
	}

	// expand 从中心向两边扩展，返回回文的左右边界（闭区间）
	expand := func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}

	start, end := 0, 0
	for i := range s {
		// 奇数长度：以 i 为中心
		l1, r1 := expand(i, i)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		// 偶数长度：以 i 和 i+1 为中心
		if i+1 < len(s) && s[i] == s[i+1] {
			l2, r2 := expand(i, i+1)
			if r2-l2 > end-start {
				start, end = l2, r2
			}
		}
	}
	return s[start : end+1]
}

// 【V3 点评】
// 优化版本，用 expand() 辅助函数统一处理边界：
//   ✅ 代码简洁，边界计算集中在 expand() 中，不易出错
//   ✅ 奇数/偶数中心统一处理，不会漏掉任何情况
//   ✅ 无需手动维护 longest 变量，直接用 end-start 比较长度
//   ✅ 时间 O(n²)，空间 O(1)，与 V2 相同
// 这是中心扩展法的标准写法，推荐面试时使用

// ──────────────────────────────────────────────────────────
// 【LeetCode 解法对比】
//
// | 解法                  | 时间   | 空间   | 面试期望        |
// |-----------------------|--------|--------|-----------------|
// | 暴力枚举所有子串       | O(n³)  | O(1)   | ❌ 太慢          |
// | 动态规划 dp[i][j]     | O(n²)  | O(n²)  | ⚠️ 能过，空间浪费 |
// | 中心扩展（V3）         | O(n²)  | O(1)   | ✅ 最佳平衡       |
// | Manacher 算法          | O(n)   | O(n)   | ⚠️ 太复杂，一般不要求 |
//
// 【面试讲解建议】
// 1. 先说思路："枚举每个位置作为回文中心，向两边扩展"
// 2. 提到有两种中心（单字符 i 和双字符 i, i+1）
// 3. 写 expand() 辅助函数
// 4. 分析复杂度
//
// 【可能被追问】
// "有没有 O(n) 的解法？"
// → Manacher 算法，面试中一般只需知道名字即可，不要求写出来。
//   可以说："Manacher 可以做到 O(n)，但实现复杂，实际面试中中心扩展法是最优选择。"
// ──────────────────────────────────────────────────────────

func main() {
	fmt.Println(longestPalindromeV2("babad"))
	fmt.Println(longestPalindromeV2("cbbd"))
	fmt.Println(longestPalindromeV2("cc"))
}
