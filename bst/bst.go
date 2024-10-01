package bst

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	val   int
}

type Tree struct {
	Root []*Node
}

func Sort() {
	arr := []int{1, 7, 4, 23, 8, 9, 4, 3, 5, 7, 9, 67, 6345, 324}
	for value := range arr {
		fmt.Println(value)
	}
}

// func (t *Tree) BuildTree(arr []int) *Node {
// 	arr = []int{1, 7, 4, 23, 8, 9, 4, 3, 5, 7, 9, 67, 6345, 324}
// }
