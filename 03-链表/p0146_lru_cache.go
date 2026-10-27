//go:build p0041

// LeetCode 第146题 - LRU 缓存 (LRU Cache)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】设计、链表、哈希表
//
// 【题目描述】
// 请你设计并实现一个满足 LRU (最近最少使用) 缓存约束的数据结构。
// 实现 LRUCache 类：
//   - LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
//   - int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1。
//   - void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value；
//     如果不存在，则向缓存中插入该组 key-value。如果插入操作导致关键字数目超过 capacity，
//     则应该逐出最久未使用的关键字。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 【示例】
// 输入：
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1,1], [2,2], [1], [3,3], [2], [4,4], [1], [3], [4]]
// 输出：
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 【数据范围】
// 1 <= capacity <= 3000
// 0 <= key <= 10^4
// 0 <= value <= 10^5
// 每次调用 get 或 put 的次数最多 2 * 10^5
//
// 【思路提示】哈希表 + 双向链表：链表头存最新、尾存最旧；
// get/put 都把节点移到头部；put 超容量时删尾节点。
// Go 里也可以直接用 container/list。
//
// 题目链接：https://leetcode.cn/problems/lru-cache/
package main

// ListNode 单链表节点（与 LeetCode Go 环境定义一致）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// LRUCache LRU 缓存。
type LRUCache struct {
	// TODO: 定义字段（容量、哈希表、双向链表）
}

// Constructor 构造 LRU 缓存。
func Constructor(capacity int) LRUCache {
	// TODO: 实现
	return LRUCache{}
}

// Get 返回 key 对应值，不存在返回 -1。
func (c *LRUCache) Get(key int) int {
	// TODO: 实现
	return -1
}

// Put 写入 key-value，超容量时淘汰最久未使用的键。
func (c *LRUCache) Put(key int, value int) {
	// TODO: 实现
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
	c.Put(3, 3) // 淘汰 key=2
	c.Get(2)    // -1
	c.Put(4, 4) // 淘汰 key=1
	c.Get(1)    // -1
	c.Get(3)    // 3
	c.Get(4)    // 4
}
