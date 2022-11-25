package main

import "fmt"

// 144 前序遍历 迭代
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	res := make([]int, 0)

	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}

// 中序遍历迭代
func inorderTraversal05(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	res := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		root = cur.Right
	}
	return res
}

// 中序遍历迭代方式
func inorderTraversal6(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			//先将左节点都加入栈
			stack = append(stack, root)
			root = root.Left
		}
		// 这里就是如果到最左边,没有左节点了,就要开始出栈处理进栈的节点
		// 遵循左根右原则
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		//查看当前节点是否存在右节点
		root = cur.Right
	}
	return res
}

// 前序遍历迭代
func preorderTraversal02(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val) //根
			stack = append(stack, root)
			root = root.Left //左
		}
		root = stack[len(stack)-1].Right //右
		stack = stack[:len(stack)-1]
	}
	return res
}

// 后序遍历迭代
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	res := make([]int, 0)
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)
	res = append(res, root.Val)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == nil {
			continue
		}
		res = append(res, cur.Val)
		stack = append(stack, cur.Left)
		stack = append(stack, cur.Right)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func main() {
	res := []int{1, 2, 3, 4}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	fmt.Println(res)
}
