package leetcode

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	input := `[5,4,8,11,null,13,4,7,2,null,null,null,1]`        // 48
	input = `[-10,9,20,null,null,15,7]`                         // 42
	input = `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]` // 16
	tree := getTreeFromInput(input)
	PrintTree(``, tree)
	fmt.Println(maxPathSumt(tree))
}

func TestLinkedList(t *testing.T) {
	input := `1,2,3,4,5,6`
	head := createLinkedListFromStr(input)
	PrintLinkedListNode(head)
}

func maxPathSumt(root *TreeNode) int {
	g := root.Val
	maxPathSumHelp(root, &g)
	return g
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSumHelp(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	ls := maxPathSumHelp(node.Left, max)
	rs := maxPathSumHelp(node.Right, max)
	sum3 := node.Val + ls + rs
	if sum3 > *max {
		*max = sum3
	}
	sum := node.Val
	if ls > rs {
		sum += ls
	} else {
		sum += rs
	}
	if sum < node.Val {
		sum = node.Val
	}
	if sum > *max {
		*max = sum
	}
	return sum
}
