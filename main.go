package main

import "fmt"

func main() {
	tree := CreateRBTree(10)
	tree.Add(5)
	tree.Add(20)
	tree.Add(80)
	tree.Add(90)
	tree.Add(15)
	tree.Add(8)
	//tree.Add(6)
	fmt.Printf("haha")
}
