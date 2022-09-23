package main

import "github.com/kidfrom/learn-golang/embedding"

func main() {
	parent := embedding.Parent{Name: "goparent"}
	// child := embedding.NewChild(parent)
	child := embedding.Child{Parent: &parent}
	child.CallParentMethod()
}
