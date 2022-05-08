package store

import (
	"fmt"
	"github.com/zzzz465/portal/sd/internal/errors"
	"github.com/zzzz465/portal/sd/internal/types"
)

type InMemoryStore struct {
	dict map[string]types.Record
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		dict: map[string]types.Record{},
	}
}

func (s *InMemoryStore) WriteRecord(key string, record types.Record) error {
	s.dict[key] = record

	return nil
}

func (s *InMemoryStore) GetRecord(key string) (*types.Record, error) {
	record, exists := s.dict[key]

	if exists {
		return &record, nil
	}

	return nil, errors.NewNotExistError(fmt.Sprintf("key %s not exists.", key))
}

func (s *InMemoryStore) DeleteRecord(key string) error {
	delete(s.dict, key)

	return nil
}
