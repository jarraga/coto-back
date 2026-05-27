package sales

type Service struct {
	store    *Store
	seedFunc func() []Sale
}

func NewService(store *Store, seedFunc func() []Sale) *Service {
	return &Service{
		store:    store,
		seedFunc: seedFunc,
	}
}

func (s *Service) Create(sale Sale) Sale {
	return s.store.Save(sale)
}

func (s *Service) Clear() {

	s.store.Clear()
}

func (s *Service) Seed() {

	s.store.Clear()

	for _, sale := range s.seedFunc() {

		s.store.Save(sale)
	}
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

func (s *Service) ModelPercentagesByCenter() []CenterModelPercentage {
	storedSales := s.store.FindAll()
	totalUnits := 0
	unitsByCenterAndModel := make(map[DistributionCenter]map[CarModel]int)

	for _, center := range DistributionCenters() {
		unitsByCenterAndModel[center] = make(map[CarModel]int)
	}

	for _, sale := range storedSales {
		totalUnits += sale.Units
		unitsByCenterAndModel[sale.DistributionCenter][sale.Model] += sale.Units
	}

	percentages := make([]CenterModelPercentage, 0, len(DistributionCenters())*len(Models()))
	for _, center := range DistributionCenters() {
		for _, model := range Models() {
			units := unitsByCenterAndModel[center][model]
			percentages = append(percentages, CenterModelPercentage{
				DistributionCenter: center,
				Model:              model,
				Units:              units,
				Percentage:         percentage(units, totalUnits),
			})
		}
	}

	return percentages
}

func percentage(units int, total int) float64 {
	if total == 0 {
		return 0
	}

	return (float64(units) / float64(total)) * 100
}
