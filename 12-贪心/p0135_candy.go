//go:build p0041

// LeetCode 第135题 - 分发糖果 (Candy)
//
// 【难度】困难 | 【面试频率】★★★★ | 【标签】贪心/数组
//
// 【题目描述】
// n 个孩子排成一排，每个孩子有评分 ratings，你需要按照要求分发糖果：每个孩子至少 1 颗；评分更高的孩子比相邻两侧评分更低的孩子获得更多糖果，求最少糖果总数。
//
// 【示例 1】
// 输入：ratings = [1,0,2]
// 输出：5（分发 [2,1,2]）
//
// 【示例 2】
// 输入：ratings = [1,2,2]
// 输出：4（分发 [1,2,1]）
// 【数据范围】
// n == ratings.length；1 <= n <= 2*10^4；0 <= ratings[i] <= 2*10^4
//
// 【思路提示】两遍扫描：左→右保证右坡递增，右→左保证左坡递增，每位取两遍结果的 max 后求和。
//
// 题目链接：https://leetcode.cn/problems/candy/
package main

// candy 返回满足相邻条件的最少糖果总数。
func candy(ratings []int) int {
	// TODO: 实现
	return 0
}

func main() {
	println(candy([]int{1, 0, 2})) // 5
	println(candy([]int{1, 2, 2})) // 4
}
