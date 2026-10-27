//go:build p0041

// LeetCode 第232题 - 用栈实现队列 (Implement Queue using Stacks)
//
// 【难度】简单 | 【面试频率】★★★★ | 【标签】栈、设计、队列
//
// 【题目描述】
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作
// （push、pop、peek、empty）：
// 实现 MyQueue 类：
//   - void push(int x) 将元素 x 推到队列的末尾
//   - int pop() 从队列的开头移除并返回元素
//   - int peek() 返回队列开头的元素
//   - boolean empty() 如果队列为空，返回 true；否则，返回 false
//
// 说明：你只能使用标准的栈操作 —— 只有 push to top, peek/pop from top, size,
// 和 is empty 操作是合法的。
//
// 【示例 1】
// 输入：
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 1, 1, false]
//
// 【数据范围】
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 所有操作都是合法的
//
// 【思路提示】入栈 inStack + 出栈 outStack：push 压入 inStack；
// pop/peek 时若 outStack 为空，把 inStack 全部倒入 outStack（顺序刚好反转成队列序）。
//
// 题目链接：https://leetcode.cn/problems/implement-queue-using-stacks/
package main

// MyQueue 用栈实现的队列。
type MyQueue struct {
	// TODO: 定义字段
}

// ConstructorQueue 构造队列。
func ConstructorQueue() MyQueue {
	// TODO: 实现
	return MyQueue{}
}

// Push 入队。
func (q *MyQueue) Push(x int) {
	// TODO: 实现
}

// Pop 出队并返回队首元素。
func (q *MyQueue) Pop() int {
	// TODO: 实现
	return 0
}

// Peek 返回队首元素。
func (q *MyQueue) Peek() int {
	// TODO: 实现
	return 0
}

// Empty 判断队列是否为空。
func (q *MyQueue) Empty() bool {
	// TODO: 实现
	return false
}

func main() {
	q := ConstructorQueue()
	q.Push(1)
	q.Push(2)
	q.Peek() // 1
	q.Pop()  // 1
	q.Empty() // false
}
