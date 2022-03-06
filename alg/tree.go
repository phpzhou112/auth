package alg

import (
	"math"
)

// 为计算n叉树深度
type NcNode struct {
	Val      int
	Children []*NcNode
}

// 计算深度,广度优先遍历
func (n *NcNode) Deepth1(root *NcNode) int {
	var ans = 0
	if root == nil {
		return ans
	}
	var queue = []*NcNode{root}

	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
	}
	return ans
}

// 考虑深度优先遍历
func (n *NcNode) Deepth2(root *NcNode) int {
	var ans = 0
	if root == nil {
		return ans
	}

	var child = root.Children
	for _, node := range child {
		childDeep := n.Deepth2(node)
		ans = int(math.Max(float64(ans), float64(childDeep)))
	}
	return ans + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最近公共祖先
func (t *TreeNode) process(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := t.process(root.Left, p, q)
	right := t.process(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}
