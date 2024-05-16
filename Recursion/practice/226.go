package main

// 226 反转二叉树
func main() {
	//var iSlice = make([]int, 3, 100)
	//setSlice(iSlice)
	//fmt.Println(cap(iSlice), len(iSlice))
	//fmt.Println(iSlice)
	//fmt.Printf("%p\n", &iSlice)
}

//func setSlice(slice []int) {
//	slice = append(slice, 1, 2, 3)
//	slice[1] = 2
//	fmt.Println(cap(slice), len(slice))
//	fmt.Printf("%p\n", &slice)
//}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// func1 递归
//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	left := invertTree(root.Left)
//	right := invertTree(root.Right)
//	root.Left = right
//	root.Right = left
//	return root
//}

// func2 直接交换,后续遍历
//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	invertTree(root.Left)
//	invertTree(root.Right)
//	root.Left, root.Right = root.Right, root.Left
//	return root
//}
