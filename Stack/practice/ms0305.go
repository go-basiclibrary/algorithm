package main

// SortedStack 面试题03.05. 栈排序
//type SortedStack struct {
//	sort.IntSlice
//}
//
//func (s SortedStack) Less(i, j int) bool {
//	return s.IntSlice[i] < s.IntSlice[j]
//}
//
//func Constructor02() SortedStack {
//	return SortedStack{
//		make([]int, 0),
//	}
//}
//
//func (s *SortedStack) Push(val int) {
//	s.IntSlice = append(s.IntSlice, val)
//	sort.Sort(s.IntSlice)
//}
//
//func (s *SortedStack) Pop() {
//	if s.IsEmpty() {
//		return
//	}
//	s.IntSlice = s.IntSlice[:s.IntSlice.Len()-1]
//}
//
//func (s *SortedStack) Peek() int {
//	if s.IsEmpty() {
//		return -1
//	}
//	return s.IntSlice[s.IntSlice.Len()-1]
//}
//
//func (s *SortedStack) IsEmpty() bool {
//	if s.Len() == 0 {
//		return true
//	}
//
//	return false
//}

//type Stack []int
//
//func (s *Stack) Len() int {
//	return len(*s)
//}
//
//func (s *Stack) Push(val int) {
//	*s = append(*s, val)
//}
//
//func (s *Stack) Pop() int {
//	if len(*s) == 0 {
//		return -1
//	}
//	val := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return val
//}
//
//func (s *Stack) Peek() int {
//	if len(*s) == 0 {
//		return -1
//	}
//	return (*s)[len(*s)-1]
//}
//
//func (s *Stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
//type SortedStack struct {
//	s1, s2 *Stack
//}
//
//func Constructor02() SortedStack {
//	return SortedStack{&Stack{}, &Stack{}}
//}
//
//func (s *SortedStack) Push(val int) {
//	s1, s2 := s.s1, s.s2
//	for s2.Len() > 0 && s2.Peek() > val {
//		s1.Push(s2.Pop())
//	}
//	for s1.Len() > 0 && s1.Peek() < val {
//		s2.Push(s1.Pop())
//	}
//	(*s1).Push(val)
//}
//
//func (s *SortedStack) Pop() {
//	s.drainToS1()
//	s.s1.Pop()
//}
//
//func (s *SortedStack) Peek() int {
//	s.drainToS1()
//	return s.s1.Peek()
//}
//
//func (s *SortedStack) IsEmpty() bool {
//	return s.s1.IsEmpty() && s.s2.IsEmpty()
//}
//
//func (s *SortedStack) drainToS1() {
//	for s.s2.Len() > 0 {
//		s.s1.Push(s.s2.Pop())
//	}
//}

// SortedStack 方法三,降序栈加辅助栈
type SortedStack struct {
	DescendingStack []int //主降序栈
	AuxiliaryStack  []int //辅助栈
}

func Constructor02() SortedStack {
	return SortedStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (s *SortedStack) Push(val int) {
	for len(s.AuxiliaryStack) > 0 && s.AuxiliaryStack[len(s.AuxiliaryStack)-1] > val {
		s.DescendingStack = append(s.DescendingStack, s.AuxiliaryStack[len(s.AuxiliaryStack)-1])
		s.AuxiliaryStack = s.AuxiliaryStack[:len(s.AuxiliaryStack)-1]
	}
	for len(s.DescendingStack) > 0 && s.DescendingStack[len(s.DescendingStack)-1] < val {
		//入辅助栈,直到不小于val
		s.AuxiliaryStack = append(s.AuxiliaryStack, s.DescendingStack[len(s.DescendingStack)-1])
		s.DescendingStack = s.DescendingStack[:len(s.DescendingStack)-1]
	}
	s.DescendingStack = append(s.DescendingStack, val)
}

func (s *SortedStack) Pop() {
	for len(s.AuxiliaryStack) > 0 {
		s.DescendingStack = append(s.DescendingStack, s.AuxiliaryStack[len(s.AuxiliaryStack)-1])
		s.AuxiliaryStack = s.AuxiliaryStack[:len(s.AuxiliaryStack)-1]
	}
	if len(s.DescendingStack) > 0 {
		s.DescendingStack = s.DescendingStack[:len(s.DescendingStack)-1]
	}
}

func (s *SortedStack) Peek() int {
	for len(s.AuxiliaryStack) > 0 {
		s.DescendingStack = append(s.DescendingStack, s.AuxiliaryStack[len(s.AuxiliaryStack)-1])
		s.AuxiliaryStack = s.AuxiliaryStack[:len(s.AuxiliaryStack)-1]
	}

	if len(s.DescendingStack) > 0 {
		return s.DescendingStack[len(s.DescendingStack)-1]
	}
	return -1
}

func (s *SortedStack) IsEmpty() bool {
	return len(s.DescendingStack) == 0 && len(s.AuxiliaryStack) == 0
}
