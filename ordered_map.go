package orderedmap

import (
	"cmp"
)

type Key interface {
	comparable
	cmp.Ordered
}

type OrderedMap[K Key, V any] struct {
	keys   []K
	values map[K]V
}

func New[K Key, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, ok := m.values[key]
	return value, ok
}

func (m *OrderedMap[K, V]) Has(key K) bool {
	_, ok := m.values[key]
	return ok
}

func (m *OrderedMap[K, V]) Del(key K) {
	delete(m.values, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *OrderedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *OrderedMap[K, V]) Values() []V {
	values := make([]V, len(m.keys))
	for i, key := range m.keys {
		values[i] = m.values[key]
	}
	return values
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.keys)
}

func (m *OrderedMap[K, V]) Clear() {
	m.keys = make([]K, 0)
	m.values = make(map[K]V)
}

func (m *OrderedMap[K, V]) Clone() *OrderedMap[K, V] {
	clone := New[K, V]()
	for _, key := range m.keys {
		clone.Set(key, m.values[key])
	}
	return clone
}

func (m *OrderedMap[K, V]) Merge(other *OrderedMap[K, V]) {
	for _, key := range other.keys {
		m.Set(key, other.values[key])
	}
}

func (m *OrderedMap[K, V]) MergeMap(other map[K]V) {
	for key, value := range other {
		m.Set(key, value)
	}
}

func (m *OrderedMap[K, V]) Range(f func(key K, value V) bool) {
	for _, key := range m.keys {
		if !f(key, m.values[key]) {
			break
		}
	}
}
