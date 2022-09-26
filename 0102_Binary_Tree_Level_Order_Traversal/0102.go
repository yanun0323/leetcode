package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	DivAndCoq(root, 0, &result)
	return result
}

func DivAndCoq(root *TreeNode, layer int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) <= layer {
		*result = append(*result, []int{})
	}
	r := *result
	r[layer] = append(r[layer], root.Val)

	DivAndCoq(root.Left, layer+1, result)
	DivAndCoq(root.Right, layer+1, result)
}

/* Queue */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for layer := 0; len(queue) > 0; layer++ {
		result = append(result, []int{})
		for size := len(queue); size > 0; size-- {
			result[layer] = append(result[layer], queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}

	return result
}
