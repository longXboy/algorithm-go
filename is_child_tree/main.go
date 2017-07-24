package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func main() {

	f0 := TreeNode{4, nil, nil}
	f1 := TreeNode{7, nil, nil}
	f2 := TreeNode{2, &f0, &f1}
	f3 := TreeNode{9, nil, nil}
	f4 := TreeNode{8, &f3, &f2}
	f5 := TreeNode{7, nil, nil}
	f6 := TreeNode{8, &f4, &f5}

	c0 := TreeNode{9, nil, nil}
	c1 := TreeNode{2, nil, nil}
	c2 := TreeNode{8, &c0, &c1}

	fmt.Println(find(&f6, &c2))
}

func find(father *TreeNode, child *TreeNode) bool {
	if father == nil || child == nil {
		return false
	}
	fmt.Println(father.value, child.value)
	if father.value == child.value {
		equal := compare(father.left, child.left)
		if equal {
			equal = compare(father.right, child.right)
			if equal {
				return true
			}
		}
	}

	if find(father.left, child) {
		return true
	}

	if find(father.right, child) {
		return true
	}
	return false
}

func compare(father *TreeNode, child *TreeNode) bool {
	if child == nil {
		return true
	} else if father == nil {
		return false
	}
	if father.value != child.value {
		return false
	}

	if !compare(father.left, child.left) {
		return false
	}
	if !compare(father.right, child.right) {
		return false
	}
	return true
}
