package main

// 105  从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 结束条件
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	// 查找中序遍历根位置
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	// 根据位置,进行下一层遍历
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])   // 左子树
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:]) // 右子树
	return root
}
