package main

// MinStack 30 包含min函数的栈
type MinStack struct {
	stackStore []int
	stackMin   []int
}

func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (s *MinStack) Push(x int) {
	s.stackStore = append(s.stackStore, x)
	if len(s.stackMin) == 0 || s.stackMin[len(s.stackMin)-1] >= x {
		s.stackMin = append(s.stackMin, x)
	}
}

func (s *MinStack) Pop() {
	if s.stackStore[len(s.stackStore)-1] == s.stackMin[len(s.stackMin)-1] {
		s.stackMin = s.stackMin[:len(s.stackMin)-1]
	}
	s.stackStore = s.stackStore[:len(s.stackStore)-1]
}

func (s *MinStack) Top() int {
	return s.stackStore[len(s.stackStore)-1]
}

func (s *MinStack) Min() int {
	if len(s.stackMin) > 0 {
		return s.stackMin[len(s.stackMin)-1]
	}
	return s.stackStore[len(s.stackStore)-1]
}
