package store

import (
	"fmt"
	errors2 "github.com/cockroachdb/errors"
	"github.com/samber/lo"
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

func (s *InMemoryStore) GetRecords() ([]types.Record, error) {
	return lo.Values(s.dict), nil
}

func (s *InMemoryStore) GetRecord(key string) (*types.Record, error) {
	record, exists := s.dict[key]

	if exists {
		return &record, nil
	}

	return nil, errors2.Wrap(errors.ErrNotExist, fmt.Sprintf("key %s not exists.", key))
}

func (s *InMemoryStore) WriteRecord(record types.Record) error {
	key := fmt.Sprintf("%s-%s", record.Metadata.DataSource, record.Name)

	s.dict[key] = record

	return nil
}

func (s *InMemoryStore) DeleteRecord(key string) error {
	delete(s.dict, key)

	return nil
}
