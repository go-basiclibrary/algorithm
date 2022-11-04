package main

// MyCircularDeque 641 双端队列,数组实现
type MyCircularDeque struct {
	dq       []int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		make([]int, 0),
		k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if len(d.dq) == d.capacity {
		return false
	}
	newdq := make([]int, 0)
	newdq = append(newdq, value)
	newdq = append(newdq, d.dq...)
	d.dq = newdq

	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if len(d.dq) == d.capacity {
		return false
	}
	d.dq = append(d.dq, value)

	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if len(d.dq) == 0 {
		return false
	}

	d.dq = d.dq[1:]
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if len(d.dq) == 0 {
		return false
	}

	d.dq = d.dq[:len(d.dq)-1]
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if len(d.dq) == 0 {
		return -1
	}

	return d.dq[0]
}

func (d *MyCircularDeque) GetRear() int {
	if len(d.dq) == 0 {
		return -1
	}

	return d.dq[len(d.dq)-1]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return len(d.dq) == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return len(d.dq) == d.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
