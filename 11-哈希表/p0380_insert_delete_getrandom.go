//go:build p0041

// LeetCode 第380题 - O(1) 时间插入、删除和获取随机元素 (Insert Delete GetRandom O(1))
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】设计、数组、哈希表、数学、随机化
//
// 【题目描述】
// 实现 RandomizedSet 类：
//   - RandomizedSet() 初始化对象
//   - bool insert(int val) 当元素 val 不存在时插入并返回 true，否则返回 false
//   - bool remove(int val) 元素存在时删除并返回 true，否则返回 false
//   - int getRandom() 随机返回现有集合中某个元素（等概率）
// 所有操作平均时间复杂度 O(1)。
//
// 【示例 1】
// 输入：
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// 输出：[null, true, false, true, 2, true, false, 2]
//
// 【数据范围】
// -2^31 <= val <= 2^31 - 1
// 最多调用 2 * 10^5 次
//
// 【思路提示】数组 + 哈希表（值→下标）：
// 删除时把最后一个元素换到被删位置再 pop，保证 O(1)。
//
// 题目链接：https://leetcode.cn/problems/insert-delete-getrandom-o1/
package main

// RandomizedSet O(1) 随机集合。
type RandomizedSet struct {
	// TODO: 定义字段
}

// ConstructorRandomizedSet 构造集合。
func ConstructorRandomizedSet() RandomizedSet {
	// TODO: 实现
	return RandomizedSet{}
}

// Insert 插入元素。
func (r *RandomizedSet) Insert(val int) bool {
	// TODO: 实现
	return false
}

// Remove 删除元素。
func (r *RandomizedSet) Remove(val int) bool {
	// TODO: 实现
	return false
}

// GetRandom 等概率返回集合中一个元素。
func (r *RandomizedSet) GetRandom() int {
	// TODO: 实现
	return 0
}

func main() {
	rs := ConstructorRandomizedSet()
	rs.Insert(1)
	rs.Remove(2)
	rs.Insert(2)
	rs.GetRandom()
	rs.Remove(1)
	rs.Insert(2)
	rs.GetRandom()
}
