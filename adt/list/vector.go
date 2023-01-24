package list

type Vector[T any] struct {
	increment int
	container []T
}

func CreateVector[T any](increment int) *Vector[T] {
	if increment <= 0 {
		increment = 4
	}
	return &Vector[T]{increment: increment, container: make([]T, 0, increment)}
}

func (t *Vector[T]) Add(element T) {
	t.grow()
	t.container = append(t.container, element)
}

func (t *Vector[T]) Get(index int) (*T, error) {
	if err := checkIndex(index, t.Size()); err != nil {
		return nil, err
	}
	return &t.container[index], nil
}

func (t *Vector[T]) Size() int {
	return len(t.container)
}

func (t *Vector[T]) grow() {
	if cap(t.container) == len(t.container) {
		newContainer := make([]T, len(t.container), cap(t.container)+t.increment)
		copy(newContainer, t.container)
		t.container = newContainer
	}
}
