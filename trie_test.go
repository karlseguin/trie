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
	assertResult(t, "", "paul", 1)
	assertResult(t, "", "leto", 2)
}

func (_ *TrieTests) ItemsWithSamePrefixes() {
	t := New(Configure())
	t.Insert("leto", 1)
	t.Insert("leto II", 2)
	assertResult(t, "", "leto", 1, 2)
	assertResult(t, "leto", " II", 2)
}

func (_ *TrieTests) ItemsWithDeepSamePrefixes() {
	t := New(Configure())
	t.Insert("apply", 1)
	t.Insert("apple", 2)
	t.Insert("applicable", 3)
	assertResult(t, "", "appl", 1, 2, 3)
	assertResult(t, "appl", "y", 1)
	assertResult(t, "appl", "e", 2)
	assertResult(t, "appl", "icable", 3)
}

func assertResult(t *Trie, fixed string, token string, ids ...int) {
	for i, l := 1, len(token); i <= l; i++ {
		prefix := fixed + token[:i]
		result := t.Find(prefix)
		Expect(result.Len()).To.Equal(len(ids)).Message("Expected %d results for %q, got %d", len(ids), prefix, result.Len())

		//results aren't ordered (what a bummer)
		seen := make(map[int]bool)
		for _, id := range ids {
			seen[id] = false
		}

		for _, id := range result.Ids() {
			seen[id] = true
		}

		for k, v := range seen {
			Expect(v).To.Equal(true).Message("Expected to see id %d for %q", k, prefix)
		}

		result.Release()
	}

	for _, extra := range []string{"!", "_", token[len(token)-1:]} {
		prefix := fixed + token + extra
		result := t.Find(prefix)
		Expect(result.Len()).To.Equal(0).Message("Expected 0 result for %q", prefix)
		result.Release()
	}
}
