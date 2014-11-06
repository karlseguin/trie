package trie

// Configuration for the trie. Initiate using Configure()
type Configuration struct {
	resultPoolCount int
	maxResults      int
}

// Initiate a configuration with sensible defaults
func Configure() *Configuration {
	return &Configuration{
		maxResults:      20,
		resultPoolCount: 16,
	}
}

// When using the Find(prefix string) method, the number of
// found ids will be limited to max. Other methods are provided
// to retrieve more than this limit, but Find() is optimized
// for performance / memory and should be used when possible
// default: 20
func (c *Configuration) MaxResults(max int) *Configuration {
	c.maxResults = max
	return c
}

// Find(prefix string) use a result pool to minimize the amount of
// allocations needed. This is the number of available results in the
// pool. (The pool won't block when drained)
// default: 16
func (c *Configuration) ResultPoolCount(count int) *Configuration {
	c.resultPoolCount = count
	return c
}
