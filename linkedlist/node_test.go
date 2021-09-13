package linkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	nodeOne := New(0)
	require.Equal(t, &Node{0, nil}, nodeOne)

	nodeTwo := New(1)
	require.Equal(t, &Node{1, nil}, nodeTwo)

	nodeOne.setNext(nodeTwo)
	require.Equal(t,&Node{0, nodeTwo}, nodeOne)

	next := nodeOne.next()
	require.Equal(t, nodeTwo, next)

	nodeOne.setItem(3)
	require.Equal(t,&Node{3, nodeTwo}, nodeOne)

	item := nodeOne.item()
	require.Equal(t, 3, item)
}
