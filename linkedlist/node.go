package linkedlist

// node a part of a data structure.
type node struct {
	Item int
	Next *node
}

// New returns a new node
func New(item int) *node {
	return &node{
		item,
		nil,
	}
}

// Item returns the Node Item
func (n *node) item() int {
	return n.Item
}

// Next returns the next Node
func (n *node) next() *node {
	return n.Next
}

// SetItem will set the Node item
func (n *node) setItem(item int) int {
	n.Item = item
	return n.Item
}

// SetNext will set the Next Node Item
func (n *node) setNext(node *node) *node {
	n.Next = node
	return n.Next
}
