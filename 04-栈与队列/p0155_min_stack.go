//go:build p0041

// LeetCode 第155题 - 最小栈 (Min Stack)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】栈、设计
//
// 【题目描述】
// 设计一个支持 push、pop、top 操作，并能在常数时间内检索到最小元素的栈。
// 实现 MinStack 类：
//   - MinStack() 初始化堆栈对象。
//   - void push(int val) 将元素 val 推入堆栈。
//   - void pop() 删除堆栈顶部的元素。
//   - int top() 获取堆栈顶部的元素。
//   - int getMin() 获取堆栈中的最小元素。
//
// 【示例 1】
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 【数据范围】
// -2^31 <= val <= 2^31 - 1
// pop、top 和 getMin 操作总是在非空栈上调用
// push、pop、top 和 getMin 最多被调用 3 * 10^4 次
//
// 【思路提示】辅助栈与主栈同步：辅助栈顶始终保存"当前栈内最小值"；
// push 时辅助栈压入 min(val, 辅助栈顶)；pop 时两栈同时弹出。
//
// 题目链接：https://leetcode.cn/problems/min-stack/
package main

// MinStack 最小栈。
type MinStack struct {
	// TODO: 定义字段
}

// ConstructorMinStack 构造最小栈（避免与 LRU 的 Constructor 重名）。
func ConstructorMinStack() MinStack {
	// TODO: 实现
	return MinStack{}
}

// Push 入栈。
func (s *MinStack) Push(val int) {
	// TODO: 实现
}

// Pop 出栈。
func (s *MinStack) Pop() {
	// TODO: 实现
}

// Top 返回栈顶元素。
func (s *MinStack) Top() int {
	// TODO: 实现
	return 0
}

// GetMin 返回栈内最小元素。
func (s *MinStack) GetMin() int {
	// TODO: 实现
	return 0
}

func main() {
	s := ConstructorMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.GetMin() // -3
	s.Pop()
	s.Top()    // 0
	s.GetMin() // -2
}
