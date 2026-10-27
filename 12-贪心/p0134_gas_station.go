//go:build p0041

// LeetCode 第134题 - 加油站 (Gas Station)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】贪心/数组
//
// 【题目描述】
// 在一条环路上有 n 个加油站，gas[i] 是第 i 个加油站的油量，cost[i] 是从第 i 站开到下一站的耗油。油箱无限、初始为空，若能绕环路一圈则返回出发站点编号，否则返回 -1（解保证唯一）。
//
// 【示例 1】
// 输入：gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// 输出：3
// 【数据范围】
// n == gas.length == cost.length；1 <= n <= 10^5；0 <= gas[i], cost[i] <= 10^4
//
// 【思路提示】总油量 >= 总耗量则必有解；从 0 出发累计净油量 < 0 时，起点改为 i+1 并清零累计。
//
// 题目链接：https://leetcode.cn/problems/gas-station/
package main

// canCompleteCircuit 返回能绕环一圈的出发站点编号，无解返回 -1。
func canCompleteCircuit(gas []int, cost []int) int {
	// TODO: 实现
	return -1
}

func main() {
	println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3
}
