package p706

const DefCap = 8
const LoadFactor = 2.5

type MyHashMap struct {
	size  int
	lock  bool
	table [][]*store
}

type store struct {
	key   int
	value int
}

func Constructor() MyHashMap {
	return MyHashMap{
		size:  0,
		lock:  false,
		table: make([][]*store, DefCap),
	}
}

func (t *MyHashMap) Put(key int, value int) {
	i := hash(key, t.cap())
	tab, ok := insert(t.table[i], key, value)
	if ok {
		t.size++
		t.table[i] = tab
		resize(t)
	}
}

func (t *MyHashMap) Get(key int) int {
	i := hash(key, t.cap())
	idx, ok := binarySearch(t.table[i], key)
	if ok {
		return t.table[i][idx].value
	}
	return -1
}

func (t *MyHashMap) Remove(key int) {
	i := hash(key, t.cap())
	tab, ok := remove(t.table[i], key)
	if ok {
		t.table[i] = tab
		t.size--
		resize(t)
	}
}

func hash(num, cap int) int {
	return num % cap
}

func (t *MyHashMap) cap() int {
	return len(t.table)
}

func insert(tab []*store, k, v int) ([]*store, bool) {
	pos := 0
	for ; pos < len(tab); pos++ {
		if tab[pos].key == k {
			tab[pos].value = v
			return nil, false
		}
		if tab[pos].key > k {
			break
		}
	}

	newTab := make([]*store, len(tab)+1)
	newMember := &store{
		key:   k,
		value: v,
	}
	for i := 0; i < len(newTab); i++ {
		if i == pos {
			newTab[i] = newMember
		} else if i < pos {
			newTab[i] = tab[i]
		} else {
			newTab[i] = tab[i-1]
		}
	}
	return newTab, true
}

func binarySearch(tab []*store, k int) (int, bool) {
	l, r, mid := 0, len(tab)-1, 0
	for l <= r {
		mid = l + (r-l)/2
		if tab[mid].key == k {
			return mid, true
		} else if tab[mid].key > k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return 0, false
}

func remove(tab []*store, k int) ([]*store, bool) {
	idx, ok := binarySearch(tab, k)
	if !ok {
		return tab, false
	}
	newTab := make([]*store, len(tab)-1)
	for i := 0; i < len(tab); i++ {
		if i == idx {
			continue
		}
		if i < idx {
			newTab[i] = tab[i]
		} else {
			newTab[i-1] = tab[i]
		}
	}
	return newTab, true
}

func resize(m *MyHashMap) {
	factor := float32(m.size) / float32(m.cap())
	if m.lock || factor < LoadFactor {
		return
	}

	m.lock = true

	newTabCap := m.cap() << 1
	tmpTab := make([]*store, m.size)
	pos := 0
	for r := 0; r < len(m.table); r++ {
		for c := 0; c < len(m.table[r]); c++ {
			tmpTab[pos] = m.table[r][c]
			pos++
		}
	}

	m.table = make([][]*store, newTabCap)
	m.size = 0
	for _, item := range tmpTab {
		m.Put(item.key, item.value)
	}

	m.lock = false
}
