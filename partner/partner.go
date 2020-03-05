package partner

import (
	"strconv"
	"strings"

	"../constant"
	"../datamodel"
	"../utils"
)

// ReadPartnersCSV ...
func ReadPartnersCSV() []*datamodel.Partner {

	records, _ := utils.CommonCSVReader(constant.Partners)
	partnerRecords := make([]*datamodel.Partner, 0)
	// Iterate through the records and bind to the above slice of partnerRecords
	for i, record := range records {
		if i > 0 {
			//Data manipulation and type conversion
			partnerRecord := new(datamodel.Partner)
			//Required to trim space as we have trailing spaces in the CSV file
			sizeSlab := strings.TrimSpace(record[1])
			minCost, _ := strconv.Atoi(strings.TrimSpace(record[2]))
			costPerGB, _ := strconv.Atoi(strings.TrimSpace(record[3]))
			size := strings.Split(sizeSlab, "-")
			minSize, _ := strconv.Atoi(size[0])
			maxSize, _ := strconv.Atoi(size[1])

			partnerRecord = &datamodel.Partner{
				TheatreID: strings.TrimSpace(record[0]),
				SizeSlab:  strings.TrimSpace(record[1]),
				MinCost:   minCost,
				CostPerGB: costPerGB,
				PartnerID: strings.TrimSpace(record[4]),
				MinSize:   minSize,
				MaxSize:   maxSize,
			}
			partnerRecords = append(partnerRecords, partnerRecord)
		}
	}
	return partnerRecords
}

// GetPartnerCapacities ...
func GetPartnerCapacities(capacityData [][]string) map[string]int {
	partnerCapacities := make(map[string]int)
	for i, v := range capacityData {
		if i > 0 {
			capacityInGB, _ := strconv.Atoi(v[1])
			partnerID := strings.TrimSpace(v[0])
			partnerCapacities[partnerID] = capacityInGB
		}
	}
	return partnerCapacities
}

// GetInputSummary ...
func GetInputSummary(inputData [][]string) map[string]int {
	inputSummary := make(map[string]int)
	for _, v := range inputData {
		size, _ := strconv.Atoi(v[1])
		theatreID := strings.TrimSpace(v[2])
		inputSummary[theatreID] += size
	}
	return inputSummary
}
