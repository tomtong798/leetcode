//go:build p0041

// LeetCode 第295题 - 数据流的中位数 (Find Median from Data Stream)
//
// 【难度】困难 | 【面试频率】★★★★★ | 【标签】双指针、数据流、堆（优先队列）
//
// 【题目描述】
// 中位数是有序整数列表中间的数。如果列表长度为偶数，则中位数是中间两个数的平均值。
// 设计一个数据结构支持：
//   - void addNum(int num) 从数据流中添加一个整数到数据结构中
//   - double findMedian() 返回目前所有元素的中位数
//
// 【示例 1】
// 输入：
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// 输出：[null, null, null, 1.5, null, 2.0]
//
// 【数据范围】
// -10^5 <= num <= 10^5
// 最多调用 5 * 10^4 次
//
// 【思路提示】双堆：大顶堆存前半（较小一半），小顶堆存后半（较大一半）；
// 保持两堆大小差 <= 1，中位数即堆顶（或两堆顶平均）。
// Go 用 container/heap 实现。
//
// 题目链接：https://leetcode.cn/problems/find-median-from-data-stream/
package main

// MedianFinder 数据流中位数查找器。
type MedianFinder struct {
	// TODO: 定义字段（大顶堆+小顶堆）
}

// ConstructorMedianFinder 构造中位数查找器。
func ConstructorMedianFinder() MedianFinder {
	// TODO: 实现
	return MedianFinder{}
}

// AddNum 添加一个数。
func (m *MedianFinder) AddNum(num int) {
	// TODO: 实现
}

// FindMedian 返回当前中位数。
func (m *MedianFinder) FindMedian() float64 {
	// TODO: 实现
	return 0
}

func main() {
	mf := ConstructorMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.FindMedian() // 1.5
	mf.AddNum(3)
	mf.FindMedian() // 2.0
}
