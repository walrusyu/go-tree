package main

import "github.com/walrusyu/go-tree/tree"

func main() {
	//tree := tree.CreateRBTree(10)
	//tree.Add(5)
	//tree.Add(20)
	//tree.Add(80)
	//tree.Add(90)
	//tree.Add(15)
	//tree.Add(8)
	//tree.Delete(10)
	//tree.Delete(15)
	//fmt.Printf("haha")
	t := tree.CreateAVLTree()
	t.Add(3)
	t.Add(2)
	t.Add(1)
	t.Add(4)
	t.Add(5)
	t.Add(6)
	t.Add(7)
	t.Add(10)
	t.Add(9)
	t.Add(8)
}
