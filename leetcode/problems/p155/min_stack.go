package p155

type MinStack struct {
	data []int
	mins []int
	size int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0, 8),
		mins: make([]int, 0, 8),
	}
}

func (t *MinStack) Push(val int) {
	t.data = append(t.data, val)
	if t.size == 0 {
		t.size++
		t.mins = append(t.mins, val)
		return
	}
	min := t.GetMin()
	if val < min {
		min = val
	}
	t.mins = append(t.mins, min)
	t.size++
}

func (t *MinStack) Pop() {
	t.data = t.data[:t.size-1]
	t.mins = t.mins[:t.size-1]
	t.size--
}

func (t *MinStack) Top() int {
	return t.data[t.size-1]
}

func (t *MinStack) GetMin() int {
	return t.mins[t.size-1]
}
