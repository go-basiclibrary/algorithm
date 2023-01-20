package main

import (
	"strconv"
	"strings"
)

// Codec 297 二叉树序列化与反序列化
type Codec struct {
}

func Constructor02() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	// 前序遍历
	var generate func(node *TreeNode)
	var sb strings.Builder
	generate = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val) + ",")
		generate(node.Left)
		generate(node.Right)
	}
	generate(root)
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
