package sales

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) Create(sale Sale) Sale {
	return s.store.Save(sale)
}

func (s *Service) TotalVolume() TotalVolume {
	storedSales := s.store.FindAll()

	volume := TotalVolume{}
	for _, sale := range storedSales {
		volume.Units += sale.Units
		volume.TotalCents += sale.TotalCents
	}

	return volume
}
