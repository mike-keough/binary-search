package bst

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	val   int
}

type Tree struct {
	Root *Node
}

func (t *Tree) BuildTree(arr []int) /*Node*/ {
	for _, v := range arr {
		fmt.Println(v)
	}
}
