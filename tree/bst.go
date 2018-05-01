package tree

import (
	"fmt"
	"bytes"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
Inorder iterative tree traversal
https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion-and-without-stack/
 */
func (root *Node) ToString() string {
	node := root
	var buffer bytes.Buffer
	for {
		if node.left != nil {
			buffer.WriteString(fmt.Sprintf("%d ", node.left.value))
		}
		break;
	}
	return buffer.String()
}

func (root *Node) Max() int {
	node := root
	for {
		if node.right != nil {
			node = node.right
		} else {
			return node.value
		}
	}
}

func (root *Node) SecondMax() int {
	node := root
	for {
		//There is left subtree only.
		if node.left != nil && node.right == nil {
			return node.left.Max()
		}

		if node.right != nil && node.right.left == nil && node.right.right == nil {
			return node.value
		}
		node = node.right
	}
}

/*
A method that check that a binary tree is a valid binary search tree.
 */
func IsBST(root *Node) bool {
	return isBSTCheck(root, -10000000, 10000000)
}

func isBSTCheck(node *Node, lowerBound int, upperBound int) bool {
	if node == nil {
		return true
	}
	if node.value < lowerBound || node.value > upperBound {
		return false
	}
	return isBSTCheck(node.left, lowerBound, node.value) &&
		isBSTCheck(node.right, node.value, upperBound)
}
