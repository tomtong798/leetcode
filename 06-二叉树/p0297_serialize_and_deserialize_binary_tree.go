//go:build p0041

// LeetCode 第297题 - 二叉树的序列化与反序列化 (Serialize and Deserialize Binary Tree)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】树、深度优先搜索、广度优先搜索、字符串、设计、二叉树
//
// 【题目描述】
// 序列化是将一个数据结构或对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
// 同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定序列 / 反序列化算法执行逻辑，
// 你只需要保证一个二叉树可以被序列化为一个字符串并且这个字符串可以被反序列化为原来的树结构。
//
// 【示例 1】
// 输入：root = [1,2,3,null,null,4,5]
// 输出：[1,2,3,null,null,4,5]
//
// 【示例 2】
// 输入：root = []
// 输出：[]
//
// 【数据范围】
// 树中结点数在范围 [0, 10^4] 内
// -1000 <= Node.val <= 1000
//
// 【思路提示】前序递归 + "null" 占位最简：序列化时空节点输出 "null"，
// 反序列化时按逗号分割，递归消费 tokens 遇 "null" 返回 nil。
//
// 题目链接：https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
package main

import "strings"

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec 序列化/反序列化器。
type Codec struct {
	// TODO: 定义字段（可不需要）
}

// ConstructorCodec 构造 Codec。
func ConstructorCodec() Codec {
	// TODO: 实现
	return Codec{}
}

// Serialize 序列化二叉树为字符串。
func (c *Codec) Serialize(root *TreeNode) string {
	// TODO: 实现
	return ""
}

// Deserialize 从字符串反序列化二叉树。
func (c *Codec) Deserialize(data string) *TreeNode {
	// TODO: 实现
	return nil
}

func main() {
	c := ConstructorCodec()
	c.Deserialize(c.Serialize(nil))
	_ = strings.Split
}
