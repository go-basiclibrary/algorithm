package main

import "fmt"

//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，
//但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。
//当栈为空时，peek 返回 -1。

//type stack struct {
//	sort.IntSlice
//}
//
//func (x stack) Less(i, j int) bool {
//	return x.IntSlice[i] > x.IntSlice[j]
//}
//
//// SortedStack 面试 0305 栈排序
//type SortedStack struct {
//	stack
//}
//
//func Constructor03() SortedStack {
//	return SortedStack{
//		stack{make([]int, 0)},
//	}
//}
//
//func (s *SortedStack) Push(val int) {
//	s.IntSlice = append(s.IntSlice, val)
//	sort.Sort(s)
//}
//
//func (s *SortedStack) Pop() {
//	if s.IsEmpty() {
//		return
//	}
//	s.IntSlice = s.IntSlice[:s.IntSlice.Len()-1]
//}
//
//func (s *SortedStack) Peek() int { // 为空返回-1
//	if s.IsEmpty() {
//		return -1
//	}
//	return s.IntSlice[s.IntSlice.Len()-1]
//}
//
//func (s *SortedStack) IsEmpty() bool {
//	return s.IntSlice.Len() == 0
//}
//
func main() {
	s := Constructor04()
	s.Push(1)
	s.Pop()

	fmt.Println(s.IsEmpty())
}

// SortedStack 双栈求法,不用做排序
type SortedStack struct {
	s  []int // 数据栈
	sS []int // 存放小值
}

func Constructor04() SortedStack {
	return SortedStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (s *SortedStack) Push(val int) {
	for len(s.sS) > 0 && s.sS[len(s.sS)-1] > val {
		s.s = append(s.s, s.sS[len(s.sS)-1])
		s.sS = s.sS[:len(s.sS)-1]
	}
	for len(s.s) > 0 && s.s[len(s.s)-1] < val {
		s.sS = append(s.sS, s.s[len(s.s)-1])
		s.s = s.s[:len(s.s)-1]
	}
	s.s = append(s.s, val)
}

func (s *SortedStack) Pop() {
	if len(s.sS) > 0 {
		s.sS = s.sS[1:]
	} else if len(s.s) > 0 {
		s.s = s.s[:len(s.s)-1]
	}
}

func (s *SortedStack) Peek() int {
	if len(s.sS) > 0 {
		return s.sS[0]
	} else if len(s.s) > 0 {
		return s.s[len(s.s)-1]
	}
	return -1
}

func (s *SortedStack) IsEmpty() bool {
	return len(s.s) == 0 && len(s.sS) == 0
}
