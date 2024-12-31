package tree

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
)

type GTNode struct {
	Node
	Path  string
	Guide *core.Guide
}
type GTTree struct {
	Model
	nodes []GTNode
}
