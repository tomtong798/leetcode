//go:build p0041

// LeetCode 第138题 - 随机链表的复制 (Copy List with Random Pointer)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】哈希表、链表
//
// 【题目描述】
// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random，
// 该指针可以指向链表中的任何节点或空节点。返回这个链表的深拷贝。
//
// 【示例 1】
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
//
// 【数据范围】
// 0 <= n <= 1000
// -10^4 <= Node.val <= 10^4
// Node.random 为 null 或指向链表中的一个节点
//
// 【思路提示】哈希表 原节点→新节点 两遍法；
// 进阶 O(1)：新节点插入原链表中间再拆分。
//
// 题目链接：https://leetcode.cn/problems/copy-list-with-random-pointer/
package main

// Node 带随机指针的链表节点。
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 深拷贝带随机指针的链表。
func copyRandomList(head *Node) *Node {
	// TODO: 实现
	return nil
}

func main() {
	copyRandomList(nil)
}
