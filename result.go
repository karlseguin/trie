package trie

// The result of a Find.
// Must be released once you're done with it.
type Result interface {
	// The number of found ids
	Len() int

	// The found ids
	Ids() []int

	// Releases the result so that the item can be re-used
	Release()
}

// An empty result
var EmptyResult empty

type empty struct{}

func (_ empty) Len() int {
	return 0
}

func (_ empty) Ids() []int {
	return nil
}

func (_ empty) Release() {
}
