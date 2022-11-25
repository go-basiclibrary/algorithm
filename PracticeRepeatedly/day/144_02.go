package main

// 144 前序遍历  迭代方法  根左右
func preorderTraversal03(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

// 中序遍历迭代 左根右
func inorderTraversal06(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 后续遍历 迭代 左右根  演化为根右左的倒序
func postorderTraversal02(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}

	reverseS(res)
	return res
}

func reverseS(res []int) {
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
}
