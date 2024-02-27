package main

import "fmt"

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

// Given the root of a binary tree, return the length of the diameter of the tree.
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := diameterOfBinaryTree(root.Left)
	rightHeight := diameterOfBinaryTree(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func main() {
	tree := binaryTree{}

	tree.root = newNode(1)
	tree.root.Left = newNode(2)
	tree.root.Right = newNode(3)
	tree.root.Left.Left = newNode(4)
	tree.root.Left.Right = newNode(5)
	tree.root.Right.Left = newNode(6)

	fmt.Println(diameterOfBinaryTree(tree.root))
}
