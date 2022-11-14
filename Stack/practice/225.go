package main

// MyStack 225 用队列实现栈
type MyStack struct {
	stack []int
}

func Constructor225() MyStack {
	return MyStack{
		make([]int, 0),
	}
}

func (ms *MyStack) Push(x int) { //将元素x压入栈顶
	ms.stack = append(ms.stack, x)
}

func (ms *MyStack) Pop() int { // 移除并返回栈顶元素
	if len(ms.stack) > 0 {
		v := ms.stack[len(ms.stack)-1]
		ms.stack = ms.stack[:len(ms.stack)-1]
		return v
	}

	return -1
}

func (ms *MyStack) Top() int { // 返回栈顶元素
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1]
	}
	return -1
}

func (ms *MyStack) Empty() bool {
	if len(ms.stack) > 0 {
		return false
	}

	return true
}
