package embedding

import "fmt"

type Parent struct {
	Name string
}

func (p *Parent) parentMethod() {
	fmt.Println(p.Name)
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
