package sales

import "time"

type CarModel string

const (
	ModelSedan   CarModel = "Sedan"
	ModelSUV     CarModel = "SUV"
	ModelOffroad CarModel = "Offroad"
	ModelSport   CarModel = "Sport"
)

type DistributionCenter string

const (
	CenterNorth DistributionCenter = "north"
	CenterSouth DistributionCenter = "south"
	CenterEast  DistributionCenter = "east"
	CenterWest  DistributionCenter = "west"
)

type Sale struct {
	ID                 int
	Model              CarModel
	DistributionCenter DistributionCenter
	Units              int
	UnitPriceCents     int
	TotalCents         int
	CreatedAt          time.Time
}

type TotalVolume struct {
	Units      int
	TotalCents int
}

type CenterVolume struct {
	DistributionCenter DistributionCenter
	Units              int
	TotalCents         int
}

type CenterModelPercentage struct {
	DistributionCenter DistributionCenter
	Model              CarModel
	Units              int
	Percentage         float64
}

func Models() []CarModel {
	return []CarModel{ModelSedan, ModelSUV, ModelOffroad, ModelSport}
}

func DistributionCenters() []DistributionCenter {
	return []DistributionCenter{CenterNorth, CenterSouth, CenterEast, CenterWest}
}

func IsValidDistributionCenter(center DistributionCenter) bool {
	for _, validCenter := range DistributionCenters() {
		if center == validCenter {
			return true
		}
	}

	return false
}

func UnitPriceCents(model CarModel) (int, bool) {

	sportBasePrice := 1820000

	prices := map[CarModel]int{
		ModelSedan:   800000,
		ModelSUV:     950000,
		ModelOffroad: 1250000,
		// Sport includes the extra 7% tax required by the exercise.
		ModelSport: sportBasePrice + (sportBasePrice * 7 / 100),
	}

	price, ok := prices[model]
	return price, ok
}
