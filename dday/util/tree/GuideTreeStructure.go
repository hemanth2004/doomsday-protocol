package tree

import "github.com/hemanth2004/doomsday-protocol/dday/core/guides"

type GTNode struct {
	Node
	Path  string
	Guide *guides.Guide
}
type GTTree struct {
	Model
	nodes []GTNode
}
