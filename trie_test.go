package trie

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type TrieTests struct{}

func Test_Trie(t *testing.T) {
	Expectify(new(TrieTests), t)
}

func (_ *TrieTests) ItemsWithDifferentPrefixes() {
	t := New(Configure())
	t.Insert("paul", 1)
	t.Insert("leto", 2)
	assertResult(t.Find("p"), 1)
	assertResult(t.Find("l"), 2)
}

func assertResult(result Result, ids ...int) {
	Expect(result.Len()).To.Equal(len(ids))
	for index, id := range result.Ids() {
		Expect(id).To.Equal(ids[index])
	}
}
