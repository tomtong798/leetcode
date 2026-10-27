//go:build p0041

// LeetCode 第739题 - 每日温度 (Daily Temperatures)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】栈、单调栈、数组
//
// 【题目描述】
// 给定一个整数数组 temperatures，表示每天的温度，返回一个数组 answer，
// 其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 【示例 1】
// 输入：temperatures = [73,74,75,71,69,72,76,73]
// 输出：[1,1,4,2,1,1,0,0]
//
// 【示例 2】
// 输入：temperatures = [30,40,50,60]
// 输出：[1,1,1,0]
//
// 【示例 3】
// 输入：temperatures = [30,60,90]
// 输出：[1,1,0]
//
// 【数据范围】
// 1 <= temperatures.length <= 10^5
// 30 <= temperatures[i] <= 100
//
// 【思路提示】单调递减栈存下标：当前温度 > 栈顶温度时，栈顶下标的答案就是
// 当前下标 - 栈顶下标，弹出并重复；否则当前下标入栈。
//
// 题目链接：https://leetcode.cn/problems/daily-temperatures/
package main

// dailyTemperatures 返回每天等待更高温度的天数。
func dailyTemperatures(temperatures []int) []int {
	// TODO: 实现
	return nil
}

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	dailyTemperatures([]int{30, 40, 50, 60})
	dailyTemperatures([]int{30, 60, 90})
}
