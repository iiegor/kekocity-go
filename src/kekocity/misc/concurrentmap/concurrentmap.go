package concurrentmap

import (
	"sync"
)

type ConcurrentMap struct {
	values map[interface{}]interface{}
	mu     sync.RWMutex
}

func New() *ConcurrentMap {
	return &ConcurrentMap{values: make(map[interface{}]interface{})}
}

func (m *ConcurrentMap) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, _ := range m.values {
		delete(m.values, key)
	}
}

func (m *ConcurrentMap) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.values)
}

func (m *ConcurrentMap) Put(_key, _value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.values[_key] = _value
}

func (m *ConcurrentMap) Remove(_key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.values, _key)
}

func (m *ConcurrentMap) Contains(_key interface{}) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.values[_key]

	return ok
}

func (m *ConcurrentMap) Get(_key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.values[_key]
	return value, ok
}
