package main

import (
	"fmt"
)

func main() {
	h := insertKey(nil, 1)
	h = insertKey(h, 2)
	h = insertKey(h, 3)
	h = insertKey(h, 4)
	h = insertKey(h, 5)
	h = insertKey(h, 6)
	h = insertKey(h, 13)
	h = insertKey(h, 12)
	h = insertKey(h, 11)
	h = insertKey(h, 10)

	h = insertKey(h, 9)
	h = insertKey(h, 8)
	h = insertKey(h, 7)
	h.breadthPrint()
	fmt.Println("")
	fmt.Println("===========")
	h = deleteKey(h, 12)
	h.breadthPrint()

}

type AVLTreeNode struct {
	key   int
	depth int
	left  *AVLTreeNode
	right *AVLTreeNode
}

func (head *AVLTreeNode) breadthPrint() {
	if head == nil {
		return
	}
	queue := make([]*AVLTreeNode, 0)
	lines := make([]int, 0)

	queue = append(queue, head)
	lines = append(lines, 0)
	lastLine := 0
	for {
		if lastLine != lines[0] {
			fmt.Println(" ")
		}
		first := queue[0]
		fmt.Printf("%d(%d) ", first.key, first.depth)

		lastLine = lines[0]
		if first.left != nil {
			queue = append(queue, first.left)
			lines = append(lines, lastLine+1)
		}
		if first.right != nil {
			queue = append(queue, first.right)
			lines = append(lines, lastLine+1)
		}
		if len(queue) == 1 {
			break
		}
		queue = queue[1:]
		lines = lines[1:]
	}
}

func printMiddle(head *AVLTreeNode) {
	if head == nil {
		return
	}
	printMiddle(head.left)
	fmt.Println(head.key)
	printMiddle(head.right)
}

func depth(head *AVLTreeNode) int {
	if head == nil {
		return 0
	} else {
		return head.depth
	}
}

func updateDepth(head *AVLTreeNode) {
	if depth(head.left) > depth(head.right) {
		head.depth = depth(head.left) + 1
	} else {
		head.depth = depth(head.right) + 1
	}
}

func leftRightRotate(head *AVLTreeNode) *AVLTreeNode {
	head.left = leftRotate(head.left)
	updateDepth(head.left)
	return rightRotate(head)
}

func rightLeftRotate(head *AVLTreeNode) *AVLTreeNode {
	head.right = rightRotate(head.right)
	updateDepth(head.right)
	return leftRotate(head)
}

func rightRotate(head *AVLTreeNode) *AVLTreeNode {
	left := head.left
	head.left = head.left.right
	left.right = head
	updateDepth(head)
	return left
}

func leftRotate(head *AVLTreeNode) *AVLTreeNode {
	right := head.right
	head.right = head.right.left
	right.left = head
	updateDepth(head)
	return right
}

func findMin(head *AVLTreeNode) *AVLTreeNode {
	if head == nil {
		return nil
	} else if head.left == nil {
		return head
	}
	return findMin(head.left)
}

func findMax(head *AVLTreeNode) *AVLTreeNode {
	if head == nil {
		return nil
	} else if head.right == nil {
		return head
	}
	return findMax(head.right)
}

func deleteKey(head *AVLTreeNode, key int) *AVLTreeNode {
	if head == nil {
		return nil
	} else if head.key == key {
		if depth(head.right) > depth(head.left) {
			min := findMin(head.right)
			if min == nil {
				return nil
			}
			head.key = min.key
			head.right = deleteKey(head.right, min.key)
		} else {
			max := findMax(head.left)
			if max == nil {
				return nil
			}
			head.key = max.key
			head.left = deleteKey(head.left, max.key)
		}
	} else if head.key > key {
		head.left = deleteKey(head.left, key)
		if depth(head.left)-depth(head.right) == 2 {
			if depth(head.left.left) > depth(head.left.right) {
				head = rightRotate(head)
			} else if depth(head.left.left) == depth(head.left.right) {
				panic("left.left == left.right")
			} else {
				head = leftRightRotate(head)
			}
		} else if depth(head.left)-depth(head.right) > 2 {
			panic("left - right > 2")
		}
	} else {
		head.right = deleteKey(head.right, key)
		if depth(head.right)-depth(head.left) == 2 {
			if depth(head.right.right) > depth(head.right.left) {
				head = leftRotate(head)
			} else if depth(head.right.left) == depth(head.right.right) {
				panic("left.left == left.right")
			} else {
				head = rightLeftRotate(head)
			}
		} else if depth(head.right)-depth(head.left) > 2 {
			panic("right - left > 2")
		}
	}

	updateDepth(head)

	return head
}

func insertKey(head *AVLTreeNode, key int) *AVLTreeNode {
	if head == nil {
		head = new(AVLTreeNode)
		head.key = key
	} else if head.key == key {
		return head
	} else if head.key > key {
		head.left = insertKey(head.left, key)
		if depth(head.left)-depth(head.right) == 2 {
			if depth(head.left.left) > depth(head.left.right) {
				head = rightRotate(head)
			} else if depth(head.left.left) == depth(head.left.right) {
				panic("left.left == left.right")
			} else {
				head = leftRightRotate(head)
			}
		} else if depth(head.left)-depth(head.right) > 2 {
			panic("left - right > 2")
		}
	} else {
		head.right = insertKey(head.right, key)
		if depth(head.right)-depth(head.left) == 2 {
			if depth(head.right.right) > depth(head.right.left) {
				head = leftRotate(head)
			} else if depth(head.right.left) == depth(head.right.right) {
				panic("left.left == left.right")
			} else {
				head = rightLeftRotate(head)
			}
		} else if depth(head.right)-depth(head.left) > 2 {
			panic("right - left > 2")
		}
	}

	updateDepth(head)

	return head
}
