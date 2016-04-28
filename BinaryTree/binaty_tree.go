package main

import (
	"fmt"
)

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

// Adds new node to the three containing this value
func (bn *BinaryNode) add(value int) {
	if value <= bn.Value {
		if bn.Left != nil {
			bn.add(value)
		} else {
			bn.Left = &BinaryNode{Value: value}
		}
	} else {
		if bn.Right != nil {
			bn.add(value)
		} else {
			bn.Right = &BinaryNode{Value: value}
		}
	}
}

// Remove value from binary tree
func (bn *BinaryNode) delete() *BinaryNode {
	if bn.Left == nil && bn.Right == nil {
		return nil
	}
	if bn.Left == nil {
		return bn.Right
	}
	if bn.Right == nil {
		return bn.Left
	}

	child := bn.Left
	grandchild := bn.Right

	if grandchild != nil {
		for grandchild != nil {
			child = grandchild
			grandchild = child.Right
		}

		bn.Value = grandchild.Value
		child.Right = grandchild.Left
	} else {
		bn.Left = child.Left
		bn.Value = child.Value
	}

	return bn
}

type BinaryTree struct {
	Root *BinaryNode
}

//  Insert value into proper location in BR
func (bt *BinaryTree) add (value int) {
	if bt.Root == nil {
		bt.Root = &BinaryNode{Value:value}
	} else {
		bt.Root.add(value)
	}
}

// Check BT contains target value
func (bt *BinaryTree) contains (target int) bool {
	node := bt.Root
	for node != nil {
		if target == node.Value {
			return true
		}
		if target < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return false
}

func main() {
	a := BinaryTree{}
	a.add(5)
	a.add(10)
	a.add(1)
	fmt.Println(a.contains(11), a.contains(1))
}
