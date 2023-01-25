package tree

type TrieTree[T any] struct {
	sep *string
}

type TrieTreeStepSearcher[T any] struct {
	segments []string
	pos      int
	node     *trieNode[T]
}

type trieNode[T any] struct {
	key   *string
	value *T
}

func NewTrieTree[T any]() *TrieTree[T] {
	return &TrieTree[T]{}
}

func (t *TrieTree[T]) Add(item *string, value *T) {
}

func (t *TrieTree[T]) Del(item *string) {

}

func (t *TrieTree[T]) Searcher(item *string) {

}
