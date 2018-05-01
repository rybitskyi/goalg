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
Write a method to check that a binary tree is a valid binary search tree.
 */

//func IsBST(node *Node) bool {
//	if node.left != nil {
//		if node.left.value > node.value {
//			return false
//		}
//		return IsBST(node.left)
//	}
//	if node.right != nil {
//		if node.right.value < node.value {
//			return false
//		}
//		return IsBST(node.right)
//	}
//	return true
//}

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

func IsBST(root *Node) bool {

	node := root
	//for {
	//	//if node.left != nil {
	//	//	//node.
	//	//}
	//	return false
	//}

	if node.left != nil {
		if node.left.value > node.value {
			return false
		}
		return IsBST(node.left)
	}
	if node.right != nil {
		if node.right.value < node.value {
			return false
		}
		return IsBST(node.right)
	}
	return true
}
