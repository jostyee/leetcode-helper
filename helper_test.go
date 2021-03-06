package leetcode

import (
	"fmt"
	"testing"
)

func TestPrintTree(t *testing.T) {
	input := `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := ParseTreeFromInput(input)
	PrintTree(``, tree)
}

func TestRenderTree(t *testing.T) {
	input := `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := ParseTreeFromInput(input)
	RenderTree(tree)
}

func TestGraph(t *testing.T) {
	input := `[[4,3,1],[3,2,4],[3],[4],[]]`
	graph := ParseGraphInput(input)
	RenderGraph(graph)
}

func TestGraphEdges(t *testing.T) {
	input := `[[0,1],[0,2],[2,5],[3,4],[4,2]]`
	edges := ParseEdgesInput(input)
	RenderGraphByEdges(edges)
}

func TestLinkedList(t *testing.T) {
	input := `1,2,3,4,5,6`
	head := ParseLinkedListFromStr(input)
	// get int linked list instead of string
	//head := GetIntLinkedListFromStr(input)
	PrintLinkedListNode(head)
}

func allPathsSourceTarget(graph [][]int) [][]int {
	if len(graph) == 0 {
		return nil
	}
	path := []int{0}
	out := [][]int{}
	helper(graph, 0, &path, &out)
	return out
}

func helper(graph [][]int, n int, path *[]int, out *[][]int) {
	if n == len(graph) {
		return
	}
	t := len(graph) - 1
	for _, d := range graph[n] {
		*path = append(*path, d)
		if d == t {
			copied := make([]int, len(*path))
			copy(copied, *path)
			*out = append(*out, copied)
		}
		helper(graph, d, path, out)
		*path = (*path)[:len(*path)-1]
	}
}

func TestTmp(t *testing.T) {
	//input := `[[0,1],[0,2],[2,5],[3,4],[4,2]]`
	input := `[[4,3,1],[3,2,4],[3],[4],[]]`
	in := ParseGraphInput(input)
	RenderGraph(in)
	out := allPathsSourceTarget(in)
	fmt.Println(out)
}
