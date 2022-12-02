package main

// MinStack 155 最小栈
type MinStack struct {
	sta    []int
	minSta []int
}

func Constructor02() MinStack {
	return MinStack{
		sta:    make([]int, 0),
		minSta: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) { //将元素val推入堆栈
	ms.sta = append(ms.sta, val)
	//min元素压入辅助栈
	if len(ms.minSta) == 0 || ms.GetMin() >= val {
		ms.minSta = append(ms.minSta, val)
	}
}

func (ms *MinStack) Pop() { //删除堆栈顶部的元素
	if ms.GetMin() == ms.sta[len(ms.sta)-1] {
		ms.minSta = ms.minSta[:len(ms.minSta)-1]
	}
	ms.sta = ms.sta[:len(ms.sta)-1]
}

func (ms *MinStack) Top() int { //获取堆栈顶部的元素
	return ms.sta[len(ms.sta)-1]
}

func (ms *MinStack) GetMin() int { //获取堆栈中的最小元素
	return ms.minSta[len(ms.minSta)-1]
}
