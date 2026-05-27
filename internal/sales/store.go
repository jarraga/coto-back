package sales

import "sync"

type Store struct {
	mu     sync.RWMutex
	nextID int
	sales  []Sale
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		sales:  make([]Sale, 0),
	}
}

func (s *Store) Save(sale Sale) Sale {

	s.mu.Lock()
	defer s.mu.Unlock()

	sale.ID = s.nextID
	s.nextID++
	s.sales = append(s.sales, sale)

	return sale
}

func (s *Store) FindAll() []Sale {

	s.mu.RLock()
	defer s.mu.RUnlock()

	copiedSales := make([]Sale, len(s.sales))
	copy(copiedSales, s.sales)

	return copiedSales
}
