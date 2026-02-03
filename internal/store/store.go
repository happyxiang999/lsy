package store

import (
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"lsy/internal/types"
)

type Store struct {
	mu    sync.RWMutex
	items map[string]types.Item
	seq   uint64
}

func NewStore() *Store {
	return &Store{
		items: make(map[string]types.Item),
	}
}

func (s *Store) List() []types.Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]types.Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}

	return items
}

func (s *Store) Get(id string) (types.Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[id]
	return item, ok
}

func (s *Store) Create(name, description string) types.Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID()
	now := time.Now().UTC()
	item := types.Item{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	s.items[id] = item
	return item
}

func (s *Store) Update(id string, name *string, description *string) (types.Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[id]
	if !ok {
		return types.Item{}, false
	}

	if name != nil {
		item.Name = *name
	}
	if description != nil {
		item.Description = *description
	}
	item.UpdatedAt = time.Now().UTC()

	s.items[id] = item
	return item, true
}

func (s *Store) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return false
	}

	delete(s.items, id)
	return true
}

func (s *Store) nextID() string {
	seq := atomic.AddUint64(&s.seq, 1)
	stamp := time.Now().UnixNano()
	return strconv.FormatInt(stamp, 36) + "-" + strconv.FormatUint(seq, 36)
}
