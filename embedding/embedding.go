package embedding

import "fmt"

type Parent struct{}

func (p *Parent) parentMethod() {
	fmt.Println("parent method")
}

type Child struct {
	*Parent
}

func (c *Child) CallParentMethod() {
	c.parentMethod()
}

func NewChild(parent *Parent) *Child {
	return &Child{parent}
}
