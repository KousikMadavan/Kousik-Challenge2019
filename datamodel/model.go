package datamodel

// Partner : Struct to maintain cost details related to each partner
type Partner struct {
	TheatreID string
	SizeSlab  string
	MinCost   int
	CostPerGB int
	PartnerID string
	MinSize   int
	MaxSize   int
}

// InputSummary ...
type InputSummary struct {
	TheatreID string
	TotalSize int
}

// Output ...
type Output struct {
	DeliveryID         string
	IsDeliveryPossible bool
	PartnerID          string
	ActualCost         int
}

// TheatrePartnerData ...
type TheatrePartnerData struct {
	Partner
	Capacity   int
	IsHigher   bool
	IsAssigned bool
}
