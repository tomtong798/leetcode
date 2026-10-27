//go:build p0041

// LeetCode 第71题 - 简化路径 (Simplify Path)
//
// 【难度】中等 | 【面试频率】★★★ | 【标签】栈、字符串
//
// 【题目描述】
// 给你一个字符串 path，表示指向某一文件或目录的 Unix 风格绝对路径（以 '/' 开头），
// 请你将其转化为更加简洁的规范路径。规范路径需遵循如下格式：
//   - 始终以斜杠 '/' 开头
//   - 两个目录名之间必须只有一个斜杠 '/'
//   - 不能以斜杠 '/' 结尾（根目录除外）
//   - 不能包含 '.' 和 '..'
//
// 【示例 1】
// 输入：path = "/home/"
// 输出："/home"
//
// 【示例 2】
// 输入：path = "/../"
// 输出："/"
//
// 【示例 3】
// 输入：path = "/home//foo/"
// 输出："/home/foo"
//
// 【示例 4】
// 输入：path = "/a/./b/../../c/"
// 输出："/c"
//
// 【数据范围】
// 1 <= path.length <= 3000
// path 由英文字母、数字、'.'、'/' 或 '_' 组成
//
// 【思路提示】按 '/' 分段入栈，".."弹栈，"."与空段跳过，最后用 '/' 拼接。
//
// 题目链接：https://leetcode.cn/problems/simplify-path/
package main

// simplifyPath 简化 Unix 路径。
func simplifyPath(path string) string {
	// TODO: 实现
	return ""
}

func main() {
	simplifyPath("/home/")
	simplifyPath("/../")
	simplifyPath("/home//foo/")
	simplifyPath("/a/./b/../../c/")
}
