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
	t.Insert("aptitude", 5)
	t.Insert("apple", 2)
	t.Insert("applicable", 3)
	t.Insert("bass", 4)
	assertResult(t, "ap", "pl", 1, 2, 3)
	assertResult(t, "appl", "y", 1)
	assertResult(t, "appl", "e", 2)
	assertResult(t, "appl", "icable", 3)
	assertResult(t, "", "ap", 1, 2, 3, 5)
	assertResult(t, "ap", "titude", 5)
	assertResult(t, "", "bass", 4)
}

func (_ *TrieTests) Unicode() {
	t := New(Configure())
	t.Insert("☄hello", 1)
	t.Insert("☺hello", 2)
	t.Insert("☺happy", 3)
	t.Insert("☺smile", 4)
	assertResult(t, "", "☄hello", 1)
	assertResult(t, "", "☺", 2, 3, 4)
	assertResult(t, "☺", "h", 2, 3)
	assertResult(t, "☺h", "ello", 2)
	assertResult(t, "☺", "smile", 4)
}

func Benchmark_Load(b *testing.B) {
	trie := New(Configure())
	for i := 0; i < b.N; i++ {
		trie.Insert(randomWord(), i)
	}
}

func Benchmark_FindSmall(b *testing.B) {
	benchmarkSize(b, 25000)
}

func Benchmark_FindMedium(b *testing.B) {
	benchmarkSize(b, 250000)
}

func Benchmark_FindLarge(b *testing.B) {
	benchmarkSize(b, 2500000)
}

func benchmarkSize(b *testing.B, size int) {
	trie := New(Configure())
	for i := 0; i < size; i++ {
		trie.Insert(randomWord(), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := trie.Find(randomPrefix())
		result.Release()
	}
}

func assertResult(t *Trie, fixed string, token string, ids ...int) {
	runes := []rune(token)
	for i, l := 1, len(runes); i <= l; i++ {
		prefix := fixed + string(runes[:i])
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

	for _, extra := range []string{"!", "_", "z"} {
		prefix := fixed + token + extra
		result := t.Find(prefix)
		Expect(result.Len()).To.Equal(0).Message("Expected 0 result for %q", prefix)
		result.Release()
	}
}
