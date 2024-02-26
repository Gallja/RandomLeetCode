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

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// valuto i casi base
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Left == nil && q.Left != nil {
		return false
	}

	if p.Right == nil && q.Right != nil {
		return false
	}

	// Confronta i valori dei nodi
	if p.Val != q.Val {
		return false
	}

	// Visita il sottoalbero sinistro
	if !isSameTree(p.Left, q.Left) {
		return false
	}

	// Visita il sottoalbero destro
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

func main() {
	tree1 := binaryTree{}
	tree2 := binaryTree{}

	tree1.root = newNode(1)
	tree1.root.Left = newNode(2)
	tree1.root.Right = newNode(3)

	tree2.root = newNode(1)
	tree2.root.Left = newNode(2)
	tree2.root.Right = newNode(3)

	fmt.Println(isSameTree(tree1.root, tree2.root))
}
