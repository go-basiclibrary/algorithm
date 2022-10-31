package main

import "fmt"

// MinStack155 155 最小栈
type MinStack155 struct {
	stackStore []int
	stackMin   []int //存放最小值
}

func main() {
	ms := Constructor155()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	ms.Pop()
	fmt.Println(ms.GetMin())
}

func Constructor155() MinStack155 {
	return MinStack155{
		make([]int, 0),
		make([]int, 0),
	}
}

func (s *MinStack155) Push(val int) {
	s.stackStore = append(s.stackStore, val)
	//接入辅助栈,代表最小值
	//判断才能入辅助栈
	if len(s.stackMin) == 0 || s.stackMin[len(s.stackMin)-1] >= val {
		s.stackMin = append(s.stackMin, val)
	}
}

func (s *MinStack155) Pop() {
	if s.stackStore[len(s.stackStore)-1] == s.stackMin[len(s.stackMin)-1] {
		s.stackMin = s.stackMin[:len(s.stackMin)-1]
	}
	// 出栈
	s.stackStore = s.stackStore[:len(s.stackStore)-1]
}

func (s *MinStack155) Top() int {
	return s.stackStore[len(s.stackStore)-1]
}

func (s *MinStack155) GetMin() int {
	if len(s.stackMin) > 0 {
		return s.stackMin[len(s.stackMin)-1]
	}
	return s.stackStore[len(s.stackStore)-1]
}
