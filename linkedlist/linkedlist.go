package linkedlist

type linkedlist struct {
	Head *node
	Current *node
	Size int
}

// New returns a new linked list
func New() *linkedlist {
	return &linkedlist{
		nil,
		nil,
		0,
	}
}

// head gets the linkedlist head
func (l *linkedlist) head() *node {
	return l.Head
}

// AppendToHead appends an item to a head
func (l *linkedlist) appendToHead(item *node) {

}

// LastNode gets the lastNode of a linked list
func (l *linkedlist) lastNode() *node {

}

// append an item to a linkedlist
func (l *linkedlist) append(item *node) {

}

// remove a node from a linked list
func (l *linkedlist) remove(item *node) {

}
// findNode by its item.
func (l *linkedlist) findNode(item int) {

}

// reverse the linked list order from last to first
func (l linkedlist) reverse(head *node) {

}

// prints the items in a linked list
func (l *linkedlist) print() {

}