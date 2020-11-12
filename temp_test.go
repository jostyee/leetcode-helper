package leetcode

import (
	"testing"
)

func TestTree(t *testing.T) {
	input := `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := GetTreeFromInput(input)
	PrintTree(``, tree)
}

func TestGraph(t *testing.T) {
	input := `[[4,3,1],[3,2,4],[3],[4],[]]`
	if err := RenderGraph(input); err != nil {
		panic(err)
	}
}

func TestLinkedList(t *testing.T) {
	input := `1,2,3,4,5,6`
	head := GetLinkedListFromStr(input)
	// get int linked list instead of string
	//head := GetIntLinkedListFromStr(input)
	PrintLinkedListNode(head)
}
