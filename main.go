package main

import "fmt"

// 这种删除办法会修改原slice
func Delete[T any](src []T, index int) ([]T, error) {
	lenght := len(src)
	if index >= lenght || index < 0 {
		return nil, fmt.Errorf("index范围异常")
	}

	//逐步用后一个元素替换前一个元素
	for i := index; i+1 < lenght; i++ {
		src[i] = src[i+1]
	}

	return src[:lenght-1], nil
}

// Shrink 这是缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

//缩容纯抄袭学习 助教老师可以不批改
func calCapacity(c, l int) (int, bool) {
	// 容量 <=64 缩不缩都无所谓，因为浪费内存也浪费不了多少
	// 你可以考虑调大这个阈值，或者调小这个阈值
	if c <= 64 {
		return c, false
	}
	// 如果容量大于 2048，但是元素不足一半，
	// 降低为 0.625，也就是 5/8
	// 也就是比一半多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	// 如果在 2048 以内，并且元素不足 1/4，那么直接缩减为一半
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	// 整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}

func main() {
	src := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(src)

	src2, _ := Delete[int](src, 2)
	// src 打印结果是 0 1 3 4 5 5
	fmt.Print(src)
	// src2 打印结果是 0 1 3 4 5
	fmt.Print(src2)
}
