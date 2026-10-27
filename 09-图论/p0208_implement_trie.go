//go:build p0041

// LeetCode 第208题 - 实现 Trie (前缀树) (Implement Trie (Prefix Tree))
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】设计、字典树、哈希表、字符串
//
// 【题目描述】
// Trie（发音类似 "try"）或者说前缀树是一种树形数据结构，
// 用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，
// 例如自动补全和拼写检查。请你实现 Trie 类：
//   - Trie() 初始化前缀树对象。
//   - void insert(String word) 向前缀树中插入字符串 word。
//   - boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即在检索之前已经插入）；
//     否则，返回 false。
//   - boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix，
//     返回 true；否则，返回 false。
//
// 【示例 1】
// 输入：
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// 输出：
// [null, null, true, false, true, null, true]
//
// 【数据范围】
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
// insert、search 和 startsWith 调用次数总计不超过 3 * 10^4 次
//
// 【思路提示】每个节点维护 26 个孩子指针 + isEnd 标记；
// insert/search/startsWith 都是沿字符逐层下探，前两个区别只在末尾 isEnd 判断。
//
// 题目链接：https://leetcode.cn/problems/implement-trie-prefix-tree/
package main

// Trie 前缀树。
type Trie struct {
	// TODO: 定义字段（children [26]*Trie，isEnd bool）
}

// ConstructorTrie 构造前缀树（避免与其他 Constructor 重名）。
func ConstructorTrie() Trie {
	// TODO: 实现
	return Trie{}
}

// Insert 插入单词。
func (t *Trie) Insert(word string) {
	// TODO: 实现
}

// Search 查找完整单词是否存在。
func (t *Trie) Search(word string) bool {
	// TODO: 实现
	return false
}

// StartsWith 查找前缀是否存在。
func (t *Trie) StartsWith(prefix string) bool {
	// TODO: 实现
	return false
}

func main() {
	t := ConstructorTrie()
	t.Insert("apple")
	t.Search("apple")     // true
	t.Search("app")       // false
	t.StartsWith("app")   // true
	t.Insert("app")
	t.Search("app")       // true
}
