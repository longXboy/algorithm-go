package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	a7 := &ListNode{value: 15, next: nil}

	a6 := &ListNode{value: 13, next: a7}

	a5 := &ListNode{value: 12, next: a6}

	a4 := &ListNode{value: 11, next: a5}

	a3 := &ListNode{value: 9, next: a4}
	a2 := &ListNode{value: 7, next: a3}
	a1 := &ListNode{value: 5, next: a2}
	a0 := &ListNode{value: 1, next: a1}

	b6 := &ListNode{value: 14, next: nil}
	b5 := &ListNode{value: 14, next: b6}
	b4 := &ListNode{value: 10, next: b5}
	b3 := &ListNode{value: 8, next: b4}
	b2 := &ListNode{value: 6, next: b3}
	b1 := &ListNode{value: 4, next: b2}
	b0 := &ListNode{value: 2, next: b1}
	h := merge(a0, b0)
	for h != nil {
		fmt.Println(h.value)
		h = h.next
	}
}

func merge(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	var little *ListNode
	var big *ListNode
	if a.value < b.value {
		little = a
		big = b
	} else {
		little = b
		big = a
	}
	head := little
	var littleTemp *ListNode
	var bigTemp *ListNode
	for {
		if big == nil {
			break
		}
		if little == nil {
			if big != nil {
				littleTemp.next = big
			}
			break
		}
		if little.value < big.value {
			littleTemp = little
			little = little.next
		} else {
			littleTemp.next = big
			for big != nil && little.value >= big.value {
				bigTemp = big
				big = big.next
			}
			bigTemp.next = little
			littleTemp = bigTemp
		}
	}
	return head
}
