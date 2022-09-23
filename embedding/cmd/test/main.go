package main

import "github.com/kidfrom/learn-golang/embedding"

func main() {
	parent := &embedding.Parent{}
	child := embedding.NewChild(parent)
	child.CallParentMethod()
}
