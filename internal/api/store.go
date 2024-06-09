package api

import (
	"fmt"
	"github.com/nomionz/currency-converter/pkg/ecb"
	"github.com/shopspring/decimal"
	"sync"
)

var _ ecb.Store = (*MemStore)(nil)

type MemStore struct {
	data map[string]decimal.Decimal
	mu   sync.RWMutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		data: make(map[string]decimal.Decimal),
	}
}

func (m *MemStore) Get(key string) (decimal.Decimal, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res, ok := m.data[key]
	if !ok {
		return decimal.Decimal{}, fmt.Errorf("key %s not found", key)
	}

	return res, nil
}

func (m *MemStore) GetKeys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]string, 0)
	for k := range m.data {
		keys = append(keys, k)
	}

	return keys
}

func (m *MemStore) Put(key string, value decimal.Decimal) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}
