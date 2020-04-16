package leetcode

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  string
	Next *ListNode
}

func createLinkedListFromStr(input string) *ListNode {
	arr := strings.Split(input, `,`)
	if len(arr) == 0 {
		return nil
	}
	node := &ListNode{arr[0], nil}
	p := node
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{arr[i], nil}
		p = p.Next
	}
	return node
}

func PrintLinkedListNode(node *ListNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Printf("%s->", p.Val)
	}
	fmt.Println()
}
