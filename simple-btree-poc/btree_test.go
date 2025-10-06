package btree

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// Helper: compares byte slices
func eq(a, b []byte) bool {
	return bytes.Equal(a, b)
}

func TestEmptyTree(t *testing.T) {
	bt := New()
	if v := bt.Search([]byte("A")); v != nil {
		t.Fatalf("expected nil on empty tree, got %v", v)
	}
}

func TestSingleInsert(t *testing.T) {
	bt := New()
	bt.Insert([]byte("A"), []byte("1"))
	got := bt.Search([]byte("A"))
	if !eq(got, []byte("1")) {
		t.Fatalf("expected 1, got %v", got)
	}
}

func TestBTreeBasicOperations(t *testing.T) {
	bt := New()

	bt.Insert([]byte("A"), []byte("1"))
	bt.Insert([]byte("B"), []byte("2"))
	bt.Insert([]byte("C"), []byte("3"))
	bt.Insert([]byte("D"), []byte("4"))
	bt.Insert([]byte("E"), []byte("5"))

	tests := []struct {
		key      []byte
		expected []byte
	}{
		{[]byte("A"), []byte("1")},
		{[]byte("B"), []byte("2")},
		{[]byte("C"), []byte("3")},
		{[]byte("D"), []byte("4")},
		{[]byte("E"), []byte("5")},
		{[]byte("Z"), nil}, // not found
	}

	for _, tt := range tests {
		val := bt.Search(tt.key)
		if !eq(val, tt.expected) {
			t.Errorf("Search(%s) = %s, want %s", tt.key, val, tt.expected)
		}
	}
}

func TestBTreeUpdate(t *testing.T) {
	bt := New()
	bt.Insert([]byte("A"), []byte("1"))
	bt.Insert([]byte("A"), []byte("updated"))

	val := bt.Search([]byte("A"))
	if !eq(val, []byte("updated")) {
		t.Errorf("Expected updated value, got %s", val)
	}
}

func TestBTreeNonExistingKey(t *testing.T) {
	bt := New()
	bt.Insert([]byte("A"), []byte("1"))
	bt.Insert([]byte("B"), []byte("2"))
	if v := bt.Search([]byte("X")); v != nil {
		t.Errorf("Expected nil for missing key, got %v", v)
	}
}

func TestBTreeBulkInsert(t *testing.T) {
	bt := New()
	keys := []string{"M", "B", "Q", "A", "Z", "L", "N", "T", "F", "E", "R"}
	for i, k := range keys {
		bt.Insert([]byte(k), []byte(fmt.Sprintf("val-%d", i)))
	}

	for i, k := range keys {
		v := bt.Search([]byte(k))
		expected := []byte(fmt.Sprintf("val-%d", i))
		if !eq(v, expected) {
			t.Errorf("Search(%s) = %s, want %s", k, v, expected)
		}
	}
}

func TestBTreeInOrderTraversal(t *testing.T) {
	bt := New()
	keys := []string{"C", "A", "E", "B", "D"}
	for _, k := range keys {
		bt.Insert([]byte(k), []byte(k))
	}

	var result [][]byte
	inOrder(bt.root, &result)

	expected := [][]byte{[]byte("A"), []byte("B"), []byte("C"), []byte("D"), []byte("E")}
	if len(result) != len(expected) {
		t.Fatalf("in-order traversal length mismatch: got %d, want %d", len(result), len(expected))
	}

	for i := range result {
		if !eq(result[i], expected[i]) {
			t.Errorf("in-order[%d] = %s, want %s", i, result[i], expected[i])
		}
	}
}

// Helper for in-order traversal
func inOrder(n *Node, out *[][]byte) {
	if n == nil {
		return
	}
	for i := 0; i < len(n.keys); i++ {
		if !n.leaf {
			inOrder(n.children[i], out)
		}
		*out = append(*out, n.keys[i])
	}
	if !n.leaf {
		inOrder(n.children[len(n.keys)], out)
	}
}

func TestBTreeRandomInsertions(t *testing.T) {
	bt := New()
	rand.Seed(time.Now().UnixNano())

	const N = 1000
	keys := make([]string, N)
	for i := range keys {
		keys[i] = fmt.Sprintf("k%04d", rand.Intn(10000))
		bt.Insert([]byte(keys[i]), []byte(keys[i]))
	}

	// Sort unique keys
	unique := make(map[string]struct{})
	for _, k := range keys {
		unique[k] = struct{}{}
	}

	var sorted []string
	for k := range unique {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)

	// Ensure all keys exist
	for _, k := range sorted {
		v := bt.Search([]byte(k))
		if !eq(v, []byte(k)) {
			t.Fatalf("Search(%s) = %s, want %s", k, v, k)
		}
	}
}
