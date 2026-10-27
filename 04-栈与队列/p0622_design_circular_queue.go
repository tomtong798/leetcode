//go:build p0041

// LeetCode 第622题 - 设计循环队列 (Design Circular Queue)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】设计、队列、数组
//
// 【题目描述】
// 设计你的循环队列实现。循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则
// 并且队尾被连接在队首之后以形成一个循环。它也被称为"环形缓冲器"。
// 实现 MyCircularQueue 类：
//   - MyCircularQueue(k) 构造器，设置队列长度为 k
//   - boolean enQueue(int value) 向循环队列插入一个元素
//   - boolean deQueue() 从循环队列中删除一个元素
//   - int Front() 从队首获取元素
//   - int Rear() 获取队尾元素
//   - boolean isEmpty() / isFull()
//
// 【示例】
// 输入：
// ["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// 输出：[null, true, true, true, false, 3, true, true, true, 4]
//
// 【数据范围】
// 1 <= k <= 1000
// 0 <= value <= 1000
// 最多调用 3000 次
//
// 【思路提示】固定容量数组 + front 索引 + size 计数，最简洁不易错。
//
// 题目链接：https://leetcode.cn/problems/design-circular-queue/
package main

// MyCircularQueue 循环队列。
type MyCircularQueue struct {
	// TODO: 定义字段
}

// ConstructorCircularQueue 构造循环队列。
func ConstructorCircularQueue(k int) MyCircularQueue {
	// TODO: 实现
	return MyCircularQueue{}
}

// EnQueue 入队。
func (q *MyCircularQueue) EnQueue(value int) bool {
	// TODO: 实现
	return false
}

// DeQueue 出队。
func (q *MyCircularQueue) DeQueue() bool {
	// TODO: 实现
	return false
}

// Front 返回队首元素。
func (q *MyCircularQueue) Front() int {
	// TODO: 实现
	return -1
}

// Rear 返回队尾元素。
func (q *MyCircularQueue) Rear() int {
	// TODO: 实现
	return -1
}

// IsEmpty 判空。
func (q *MyCircularQueue) IsEmpty() bool {
	// TODO: 实现
	return false
}

// IsFull 判满。
func (q *MyCircularQueue) IsFull() bool {
	// TODO: 实现
	return false
}

func main() {
	q := ConstructorCircularQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4) // false
	q.Rear()     // 3
	q.IsFull()   // true
	q.DeQueue()  // true
	q.EnQueue(4) // true
	q.Rear()     // 4
}
