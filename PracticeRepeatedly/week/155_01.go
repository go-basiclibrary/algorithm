package main

type MinStack struct {
	Stack    []int
	MinStore []int
}

func Constructor155() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (m *MinStack) Push(val int) { //void push(int val) 将元素val推入堆栈。
	if len(m.MinStore) == 0 || m.MinStore[len(m.MinStore)-1] >= val {
		m.MinStore = append(m.MinStore, val)
	}
	m.Stack = append(m.Stack, val)
}

func (m *MinStack) Pop() { //void pop() 删除堆栈顶部的元素。
	if m.Stack[len(m.Stack)-1] == m.MinStore[len(m.MinStore)-1] {
		m.MinStore = m.MinStore[:len(m.MinStore)-1]
	}
	m.Stack = m.Stack[:len(m.Stack)-1]
}

func (m *MinStack) Top() int { //int top() 获取堆栈顶部的元素。
	return m.Stack[len(m.Stack)-1]
}

func (m *MinStack) GetMin() int { //int getMin() 获取堆栈中的最小元素。
	if len(m.MinStore) > 0 {
		return m.MinStore[len(m.MinStore)-1]
	}
	return m.Stack[len(m.Stack)-1]
}
