package trie

import (
	"fmt"
	"github.com/karlseguin/scratch"
)

type Leaf struct {
	id     int
	suffix string
}

type Node struct {
	leafs map[byte]*Leaf
	nodes map[byte]*Node
}

func newNode() *Node {
	return &Node{
		nodes: make(map[byte]*Node),
	}
}

type Trie struct {
	root *Node
	results *scratch.IntsPool
}

func New(c *Configuration) *Trie {
	return &Trie{
		root: newNode(),
		results: scratch.NewInts(c.maxResults, c.resultPoolCount),
	}
}

func (t *Trie) Insert(value string, id int) {
	if len(value) == 0 {
		return
	}
	node, exists := t.root, false
	for i := 0; i < len(value); i++ {
		c := value[i]
		parent := node
		node, exists = node.nodes[c]
		if exists {
			continue
		}
		if leaf, exists := parent.leafs[c]; exists {
			if leaf.suffix == value[i+1:] {
				leaf.id = id
				break
			}
			node = newNode()
			parent.nodes[c] = node

			ls := len(leaf.suffix)
			if ls == 0 {
				node.addLeaf(value, 0, id)
				break
			}

			suffix, value := leaf.suffix, value[i+1:]
			lv := len(value)
			if lv == 0 {
				parent.leafs[c] = &Leaf{id, ""}
				node.addLeaf(suffix, 0, leaf.id)
				break
			}

			delete(parent.leafs, c)
			j := 0
			for ; j < ls && j < lv && suffix[j] == value[j]; j++ {
				parent = node
				node = newNode()
				parent.nodes[suffix[j]] = node
			}
			if j != lv {
				node.addLeaf(value, j, id)
			} else {
				parent.addLeaf(value, j-1, id)
			}
			if j != ls {
				node.addLeaf(suffix, j, leaf.id)
			} else if ls != 0 {
				parent.addLeaf(suffix, j-1, leaf.id)
			}
		} else {
			if parent.leafs == nil {
				parent.leafs = make(map[byte]*Leaf)
			}
			parent.leafs[c] = &Leaf{id, value[i+1:]}
		}
		break
	}
}

func (t *Trie) Dump() {
	Dump(t.root, "")
}

func Dump(n *Node, prefix string) {
	for k, value := range n.leafs {
		fmt.Println(prefix, string(k), "=>", value.suffix, value.id)
	}
	for k, node := range n.nodes {
		fmt.Println(prefix, string(k), "->")
		Dump(node, prefix+"   ")
	}
}

func (t *Trie) Find(prefix string) Result {
	if len(prefix) == 0 {
		return EmptyResult
	}

	node, exists := t.root, false
	parent := node
	i, l := 0, len(prefix)
	for ; i < l; i++ {
		if node, exists = node.nodes[prefix[i]]; exists == false {
			break
		}
		parent = node
	}
	if exists == false {
		if i == l-1 {
			if leaf, exists := parent.leafs[prefix[i]]; exists {
				result := t.results.Checkout()
				result.Add(leaf.id)
				return result
			}
		}
		return EmptyResult
	}

	result := t.results.Checkout()
	populate(node, result)
	return result
}

func populate(node *Node, result *scratch.Ints) bool {
	for _, leaf := range node.leafs {
		if result.Add(leaf.id) == false {
			return false
		}
	}
	for _, node := range node.nodes {
		if populate(node, result) == false {
			return false
		}
	}
	return true
}

func (n *Node) addLeaf(value string, index int, id int) {
	if n.leafs == nil {
		n.leafs = make(map[byte]*Leaf)
	}
	n.leafs[value[index]] = &Leaf{id, value[index+1:]}
}
