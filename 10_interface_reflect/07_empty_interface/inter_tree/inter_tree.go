package main

import "fmt"

type INodeTree interface {
	SetData(interface{})
	Next() INodeTree
	Prev() INodeTree
	SetNext(INodeTree)
	SetPrev(INodeTree)
}

type Node struct {
	l    INodeTree
	data interface{}
	r    INodeTree
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) Next() INodeTree {
	return n.r
}
func (n *Node) Prev() INodeTree {
	return n.l
}

func (n *Node) SetNext(p INodeTree) {
	n.r = p
}

func (n *Node) SetPrev(l INodeTree) {
	n.l = l
}

func NewNode(l INodeTree, data interface{}) INodeTree {
	n := &Node{l: l, data: data}
	if l != nil {
		l.SetNext(n)
	}
	return n
}

func main() {
	n := NewNode(nil,"酒")
	n1 := NewNode(n,1)
	NewNode(n1,[]int{1,2,3})
	P(n)
}

func P(node INodeTree) {
	for e := node;e != nil; e = e.Next(){
		fmt.Printf("e: %v\n", e)
	}
}
