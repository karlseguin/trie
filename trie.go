package trie

import (
	"fmt"
	"github.com/karlseguin/scratch"
	"strings"
)

type Leaf struct {
	id     int
	key    byte
	suffix string
}

type Node struct {
	leafs []Leaf
	nodes map[byte]*Node
}

type Trie struct {
	root    *Node
	results *scratch.IntsPool
}

func New(c *Configuration) *Trie {
	return &Trie{
		root:    &Node{},
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
		if leaf, exists := parent.findLeaf(c); exists {
			if leaf.suffix == value[i+1:] {
				leaf.id = id
				break
			}
			node = parent.addNode(c)

			ls := len(leaf.suffix)
			if ls == 0 {
				node.addLeaf(value, 0, id)
				break
			}

			suffix, value := leaf.suffix, value[i+1:]
			lv := len(value)
			if lv == 0 {
				parent.leafs = append(parent.leafs, Leaf{id, c, ""})
				node.addLeaf(suffix, 0, leaf.id)
				break
			}
			parent.deleteLeaf(c)
			j := 0
			for ; j < ls && j < lv && suffix[j] == value[j]; j++ {
				parent = node
				node = parent.addNode(suffix[j])
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
				parent.leafs = make([]Leaf, 0, 1)
			}
			parent.leafs = append(parent.leafs, Leaf{id, c, value[i+1:]})
		}
		break
	}
}

func (t *Trie) Dump() {
	Dump(t.root, "")
}

func Dump(n *Node, prefix string) {
	for _, value := range n.leafs {
		fmt.Println(prefix, string(value.key), "=>", value.suffix, value.id)
	}
	for k, node := range n.nodes {
		fmt.Println(prefix, string(k), "->")
		Dump(node, prefix+"   ")
	}
}
func (t *Trie) Stats() {
	stats := make(map[int]int)
	Stats(t.root, stats)
	for k, v := range stats {
		fmt.Println(k, v)
	}
}

func Stats(n *Node, stats map[int]int) {
	leafs := len(n.nodes)
	if _, exists := stats[leafs]; exists == false {
		stats[leafs] = 0
	}
	stats[leafs] += 1
	for _, node := range n.nodes {
		Stats(node, stats)
	}
}

func (t *Trie) Find(prefix string) Result {
	if len(prefix) == 0 {
		return EmptyResult
	}

	node, exists := t.root, false
	grand, parent := node, node
	i, l := 0, len(prefix)
	for ; i < l; i++ {
		if node, exists = node.nodes[prefix[i]]; exists == false {
			break
		}
		grand, parent = parent, node
		parent = node
	}
	if exists == false {
		leaf, exists := parent.findLeaf(prefix[i])
		if exists == false {
			return EmptyResult
		}
		if strings.HasPrefix(leaf.suffix, prefix[i+1:]) {
			result := t.results.Checkout()
			result.Add(leaf.id)
			return result
		}
		return EmptyResult
	}

	result := t.results.Checkout()
	if i == l {
		if leaf, exists := grand.findLeaf(prefix[l-1]); exists && len(leaf.suffix) == 0 {
			result.Add(leaf.id)
		}
	}
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
		n.leafs = make([]Leaf, 0, 1)
	}
	n.leafs = append(n.leafs, Leaf{id, value[index], value[index+1:]})
}


func (n *Node) findLeaf(c byte) (Leaf, bool) {
	for i, l := 0, len(n.leafs); i < l; i++ {
		leaf := n.leafs[i]
		if leaf.key == c {
			return leaf, true
		}
	}
	return Leaf{}, false
}

func (n *Node) deleteLeaf(c byte) {
	for i, l := 0, len(n.leafs); i < l; i++ {
		leaf := n.leafs[i]
		if leaf.key == c {
			n.leafs[i] = n.leafs[l-1]
			n.leafs = n.leafs[:l-1]
			return
		}
	}
}

func (n *Node) addNode(b byte) *Node {
	if n.nodes == nil {
		n.nodes = make(map[byte]*Node)
	}
	node := &Node{}
	n.nodes[b] = node
	return node
}
