package main

type MinStack struct {
	stackStore []int
	stackMin   []int
}

func Constructor02() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	if len(m.stackMin) == 0 || m.stackMin[len(m.stackMin)-1] >= val {
		m.stackMin = append(m.stackMin, val)
	}
	m.stackStore = append(m.stackStore, val)
}

func (m *MinStack) Pop() {
	if m.stackStore[len(m.stackStore)-1] == m.stackMin[len(m.stackMin)-1] {
		m.stackMin = m.stackMin[:len(m.stackMin)-1]
	}
	m.stackStore = m.stackStore[:len(m.stackStore)-1]
}

func (m *MinStack) Top() int {
	return m.stackStore[len(m.stackStore)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.stackMin) > 0 {
		return m.stackMin[len(m.stackMin)-1]
	}
	return m.stackStore[len(m.stackStore)-1]
}
