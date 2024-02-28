package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type binaryTree struct {
	root *TreeNode
}

func newNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

// Given the root of a binary tree, return the leftmost value in the last row of the tree.
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return root.Val
	}

	// elemento più in profondità del sottoalbero di sx:
	sx := findBottomLeftValue(root.Left)

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if rightDepth > leftDepth {
		return findBottomLeftValue(root.Right)
	}

	return sx
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	tree := binaryTree{}

	tree.root = newNode(1)
	tree.root.Left = newNode(2)
	tree.root.Left.Left = newNode(4)
	tree.root.Right = newNode(3)
	tree.root.Right.Left = newNode(5)
	tree.root.Right.Left.Left = newNode(7)
	tree.root.Right.Right = newNode(6)

	fmt.Println(findBottomLeftValue(tree.root))
}
