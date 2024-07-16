package ordered

type OrderedSet[K Key] struct {
	*OrderedMap[K, bool]
}

func NewSet[K Key]() *OrderedSet[K] {
	return &OrderedSet[K]{New[K, bool]()}
}

func (m *OrderedSet[K]) Add(key K) {
	m.Set(key, true)
}

func (m *OrderedSet[K]) Elements() []K {
	return m.Keys()
}

func (m *OrderedSet[K]) Union(other *OrderedSet[K]) *OrderedSet[K] {
	result := NewSet[K]()
	result.Merge(m.OrderedMap)
	result.Merge(other.OrderedMap)
	return result
}

func (m *OrderedSet[K]) Intersection(other *OrderedSet[K]) *OrderedSet[K] {
	result := NewSet[K]()
	for _, key := range m.Keys() {
		if other.Has(key) {
			result.Add(key)
		}
	}
	return result
}

func (m *OrderedSet[K]) Difference(other *OrderedSet[K]) *OrderedSet[K] {
	result := NewSet[K]()
	result.Merge(m.OrderedMap)
	for _, key := range other.Keys() {
		result.Del(key)
	}
	return result
}
