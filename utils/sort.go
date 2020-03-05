package utils

import (
	"../datamodel"
)

// SortPartner ...
func SortPartner(partner []*datamodel.TheatrePartnerData) []*datamodel.TheatrePartnerData {
	sort(partner)
	return partner
}

func sort(items []*datamodel.TheatrePartnerData) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i].CostPerGB >= items[i+1].CostPerGB {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
