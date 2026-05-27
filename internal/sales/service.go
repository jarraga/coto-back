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

func (s *Service) VolumeByCenter() []CenterVolume {
	storedSales := s.store.FindAll()
	volumesByCenter := make(map[DistributionCenter]CenterVolume)

	for _, center := range DistributionCenters() {
		volumesByCenter[center] = CenterVolume{
			DistributionCenter: center,
		}
	}

	for _, sale := range storedSales {
		volume := volumesByCenter[sale.DistributionCenter]
		volume.Units += sale.Units
		volume.TotalCents += sale.TotalCents
		volumesByCenter[sale.DistributionCenter] = volume
	}

	volumes := make([]CenterVolume, 0, len(volumesByCenter))
	for _, center := range DistributionCenters() {
		volumes = append(volumes, volumesByCenter[center])
	}

	return volumes
}
