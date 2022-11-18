package main

// TreeNode 94 二叉树中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var reslove func(node *TreeNode)
	reslove = func(node *TreeNode) {
		//递归终结条件
		if node == nil {
			return
		}
		//处理当前层逻辑
		// 中序遍历重要思想,左根右
		reslove(node.Left)
		res = append(res, node.Val)
		//下到下层
		reslove(node.Right)
		//清理当前层
		// 无需清理
	}
	reslove(root)
	return res
}

// 迭代,递归演变而来,隐藏栈显现出来
func inorderTraversal02(root *TreeNode) []int {
	//维护一个栈,存储左根
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			// 将左根入栈
			stack = append(stack, root)
			root = root.Left
		}
		//取根,开始左右拿值
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// Morris中序遍历
func inorderTraversal03(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前root节点向左走一步,然后一直向右走至无法走未知的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				//有右子树且没有设置过指向root,则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将predecessor的右指针指向root,这样后面遍历完左子树,root.Left后,就能通过这个指向回到root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了root,则表示左子树root.Left已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { //没有左子树
			res = append(res, root.Val)
			// 若有右子树,则遍历右子树
			// 若没有右子树,则整颗右子树已遍历完
			root = root.Right
		}
	}
	return res
}
