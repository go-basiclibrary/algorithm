package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144 二叉树前序遍历 根左右
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	//递归
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		//终止条件
		if node == nil {
			return
		}
		//逻辑
		res = append(res, node.Val)
		//下层
		preorder(node.Left)
		preorder(node.Right)
		//清理当前层
	}
	preorder(root)

	return res
}

// 迭代
func preorderTraversal02(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// 前序遍历 morris 算法
func preorderTraversal03(root *TreeNode) []int {
	res := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left //将左根最右节点,指向头根
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil { // 找到左根最右根
				p2.Right = p1
				res = append(res, p1.Val)
				p1 = p1.Left
				continue
			}
			p2.Right = nil //清理指向
		} else {
			res = append(res, p1.Val)
		}
		p1 = p1.Right
	}
	return res
}
