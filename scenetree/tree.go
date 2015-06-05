package scenetree

import (
	glm "github.com/go-gl/mathgl/mgl32"
)

type Renderable interface {
	Render() //bunch of arguments to differentiate shadows, cubemap and whatnot
}

type Node struct {
	children  []Renderable
	transform glm.Mat4
}

func (n *Node) Render() {
	for _, child := range n.children {
		child.Render()
	}
}
