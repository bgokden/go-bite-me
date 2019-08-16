package gobiteme

import (
	"bytes"
)

// GoBiteMe is a wrapper struct for all functions
type GoBiteMe struct {
	Root *Tree
}

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value *Bite
	Right *Tree
}

// Bite is Key Value Pair
type Bite struct {
	Key   []byte
	Value []byte
}

// insert triggers a recursive insert method
func insert(t *Tree, key, value []byte) *Tree {
	if t == nil {
		return &Tree{
			Left: nil,
			Value: &Bite{
				Key:   key,
				Value: value,
			},
			Right: nil,
		}
	}
	diff := bytes.Compare(key, t.Value.Key)
	if diff == 0 {
		return t
	}
	if diff < 0 {
		t.Left = insert(t.Left, key, value)
		return t
	}
	t.Right = insert(t.Right, key, value)
	return t
}

// New returns just an empty instance
func New() *GoBiteMe {
	return &GoBiteMe{}
}

// Bite allow insterting a data
func (b *GoBiteMe) Bite(key, value []byte) {
	if b.Root == nil {
		b.Root = &Tree{nil, &Bite{key, value}, nil}
	} else {
		insert(b.Root, key, value)
	}
}

// ThrowUp lets you iterate over whole data
func (b *GoBiteMe) ThrowUp(bucket func(*Bite)) {
	if b.Root == nil {
		return
	}
	throwUpMore(b.Root, bucket)
}

// throwUpMore is recursive function needed for ThrowUp
func throwUpMore(t *Tree, bucket func(bite *Bite)) {
	if t == nil {
		return
	}
	throwUpMore(t.Left, bucket)
	bucket(t.Value)
	throwUpMore(t.Right, bucket)
}
