package btree

import (
	"bytes"
	"fmt"
	"sort"
)

const (
	MinDegree = 2
)

type BTree struct {
	root *Node
}

type Node struct {
	keys     [][]byte
	values   [][]byte
	children []*Node
	leaf     bool
}

func New() *BTree {
	return &BTree{root: &Node{leaf: true}}
}

func (t *BTree) Search(key []byte) []byte {
	return t.root.search(key)
}

func (n *Node) search(key []byte) []byte {
	// Binary search for position
	i := sort.Search(len(n.keys), func(i int) bool {
		return bytes.Compare(n.keys[i], key) >= 0
	})

	if i < len(n.keys) && bytes.Equal(n.keys[i], key) {
		return n.values[i]
	}
	if n.leaf {
		return nil
	}

	return n.children[i].search(key)
}

func (t *BTree) Insert(key, value []byte) {
	r := t.root

	if len(r.keys) == 2*MinDegree-1 {
		s := &Node{leaf: false}
		s.children = []*Node{r}
		s.splitChild(0)
		t.root = s
	}
	t.root.insertNonFull(key, value)
}

func (n *Node) insertNonFull(key, value []byte) {

	// Binary search for position
	pos := sort.Search(len(n.keys), func(i int) bool {
		return bytes.Compare(n.keys[i], key) >= 0
	})

	// Update if key exists
	if pos < len(n.keys) && bytes.Equal(n.keys[pos], key) {
		n.values[pos] = value
		return
	}

	if n.leaf {
		// Insert new key/value
		n.keys = append(n.keys, nil)
		n.values = append(n.values, nil)
		copy(n.keys[pos+1:], n.keys[pos:])
		copy(n.values[pos+1:], n.values[pos:])
		n.keys[pos] = key
		n.values[pos] = value
	} else {
		if len(n.children[pos].keys) == 2*MinDegree-1 {
			n.splitChild(pos)
			if bytes.Compare(key, n.keys[pos]) > 0 {
				pos++
			}
		}
		n.children[pos].insertNonFull(key, value)
	}
}

func (n *Node) splitChild(i int) {
	y := n.children[i]
	z := &Node{leaf: y.leaf}
	m := MinDegree - 1

	midKey := y.keys[m]
	midVal := y.values[m]

	z.keys = append([][]byte(nil), y.keys[m+1:]...)
	z.values = append([][]byte(nil), y.values[m+1:]...)

	y.keys = y.keys[:m]
	y.values = y.values[:m]

	if !y.leaf {
		z.children = append([]*Node(nil), y.children[m+1:]...)
		y.children = y.children[:m+1]
	}

	n.keys = append(n.keys, nil)
	n.values = append(n.values, nil)
	copy(n.keys[i+1:], n.keys[i:])
	copy(n.values[i+1:], n.values[i:])
	n.keys[i] = midKey
	n.values[i] = midVal

	n.children = append(n.children, nil)
	copy(n.children[i+2:], n.children[i+1:])
	n.children[i+1] = z
}

func (t *BTree) Print() {
	t.root.print(0)
}

func (n *Node) print(level int) {
	fmt.Printf("%sKeys: ", bytes.Repeat([]byte("  "), level))
	for _, k := range n.keys {
		fmt.Printf("%s ", string(k))
	}
	fmt.Println()
	for _, c := range n.children {
		c.print(level + 1)
	}
}
