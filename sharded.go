package trie

type ShardedTrie map[byte]*Trie

func NewSharded(c *Configuration) ShardedTrie {
	s := make(ShardedTrie, 255)
	for i := byte(0); i < byte(255); i++ {
		s[i] = New(c)
	}
	return s
}

func (t ShardedTrie) Insert(value string, id int) {
	if len(value) == 0 {
		return
	}
	t[value[0]].Insert(value[1:], id)
}

func (t ShardedTrie) Find(prefix string) Result {
	if len(prefix) == 0 {
		return EmptyResult
	}
	return t[prefix[0]].Find(prefix[1:])
}
