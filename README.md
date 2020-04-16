# leetcode-helper

## Tree

leetcode use a json array as input to represent a tree,

it's hard for us to read the tree.

with the helper, you can turn the input to a readable one, for example:

```go
func TestSum(t *testing.T) {
	input = `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := getTreeFromInput(input)
	PrintTree(``, tree)
}
```

the function `getTreeFromInput` turn the input to a tree,

and the func `PrintTree` will print the tree in text

the following is the printed tree:

```
         |-- 2
                 |-- -6
             |-- 2
                 |-- -6
                     |-- -6
     |-- -3
         |-- -6
 |-- 9
     |-- 6
```

the root is at the left most place, the left child locate blow

the parent, the right child locate above the parent.

so the root is `9`, it's left child is `6`, it's right child is `-3`

for the node `-3`, it's right child is `2`


## LinkedList

the example code:

```go
func TestLinkedList (t *testing.T){
	input := `1,2,3,4,5,6`
	head := createLinkedListFromStr(input)
	PrintLinkedListNode(head)
}
```

the output is 
```
1->2->3->4->5->6->
```


