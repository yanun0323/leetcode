package main

import (
	"strconv"
	"strings"
)

// FIXME: [0297] Timeout Passed
type Codec struct {
	null string
	sep  string
}

func Constructor() Codec {
	return Codec{
		null: "null",
		sep:  ",",
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	b := strings.Builder{}
	todoQueue := []*TreeNode{root}
	ifNextLayerExist := true
	for ifNextLayerExist {
		ifNextLayerExist = false
		for count := len(todoQueue); count > 0; count-- {
			if b.Len() > 0 {
				b.WriteString(c.sep)
			}
			if todoQueue[0] == nil {
				b.WriteString(c.null)
				todoQueue = todoQueue[1:]
				continue
			}
			b.WriteString(strconv.Itoa(todoQueue[0].Val))
			todoQueue = append(todoQueue, todoQueue[0].Left, todoQueue[0].Right)
			ifNextLayerExist = ifNextLayerExist || todoQueue[0].Left != nil || todoQueue[0].Right != nil
			todoQueue = todoQueue[1:]
		}
	}
	return b.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	todoArr := strings.Split(data, c.sep)
	num, _ := strconv.Atoi(todoArr[0])
	root := &TreeNode{
		Val: num,
	}
	todoArr = todoArr[1:]
	queue := []*TreeNode{root}
	for len(todoArr) > 0 {
		if todoArr[0] != c.null {
			ln, _ := strconv.Atoi(todoArr[0])
			nodeL := &TreeNode{
				Val: ln,
			}
			queue[0].Left = nodeL
			queue = append(queue, nodeL)
		}
		if todoArr[1] != c.null {
			rn, _ := strconv.Atoi(todoArr[1])
			nodeR := &TreeNode{
				Val: rn,
			}
			queue[0].Right = nodeR
			queue = append(queue, nodeR)
		}
		queue = queue[1:]
		todoArr = todoArr[2:]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
