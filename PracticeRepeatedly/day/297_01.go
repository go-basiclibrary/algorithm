package main

import (
	"strconv"
	"strings"
)

// Codec 297 二叉树的序列化与反序列化
// 中序遍历
type Codec struct {
}

func Constructor03() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var sorder func(node *TreeNode)
	sb := strings.Builder{}
	sorder = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val) + ",")
		sorder(node.Left)
		sorder(node.Right)
	}
	sorder(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
