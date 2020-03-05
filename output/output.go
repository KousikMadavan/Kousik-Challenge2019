package output

import (
	"fmt"
	"strconv"
	"strings"

	"../constant"
	"../datamodel"
	"../utils"
)

// Output1 ...
func Output1(inputData [][]string, partnerRecords []*datamodel.Partner, partnerCapacities map[string]int) []*datamodel.TheatrePartnerData {
	output1 := make([]*datamodel.Output, 0)
	theatrePartnerData := make([]*datamodel.TheatrePartnerData, 0)
	for i, input := range inputData {
		inputSize, _ := strconv.Atoi(input[1])
		var isDeliveryPossible bool
		var finalCost int
		var partnerID string
		output := new(datamodel.Output)
		for _, records := range partnerRecords {
			if inputSize > records.MinSize && inputSize < records.MaxSize && input[2] == records.TheatreID {
				isDeliveryPossible = true
				actualCost := records.CostPerGB * inputSize
				if finalCost > actualCost || finalCost == 0 {
					if actualCost < records.MinCost {
						finalCost = records.MinCost
						partnerID = records.PartnerID
					} else if actualCost > records.MinCost {
						finalCost = actualCost
						partnerID = records.PartnerID
					}
				}
			}
			if i == 0 {
				theatrePartner := &datamodel.TheatrePartnerData{Partner: datamodel.Partner{
					TheatreID: records.TheatreID,
					PartnerID: records.PartnerID,
					SizeSlab:  records.SizeSlab,
					CostPerGB: records.CostPerGB,
					MinSize:   records.MinSize,
					MaxSize:   records.MaxSize,
					MinCost:   records.MinCost,
				},
					Capacity: partnerCapacities[records.PartnerID],
				}
				theatrePartnerData = append(theatrePartnerData, theatrePartner)
			}
		}
		output = &datamodel.Output{
			DeliveryID:         input[0],
			IsDeliveryPossible: isDeliveryPossible,
			ActualCost:         finalCost,
			PartnerID:          partnerID,
		}
		output1 = append(output1, output)
	}
	utils.CreateOutputCSV(output1, constant.Output1)
	fmt.Println("Succesfully addressed problem 1")
	return theatrePartnerData
}

// Output2 ...
func Output2(inputSummary map[string]int, inputData [][]string, result []*datamodel.TheatrePartnerData) {
	var canAccomodate bool
	output2 := make([]*datamodel.Output, 0)
	var index int
	for summaryKey, sum := range inputSummary {
		for _, val := range result {
			// sum is the total amount for each theatre
			if sum >= val.MinSize && sum <= val.MaxSize {
				canAccomodate = true
			}
			continue
		}
		if !canAccomodate {
			for _, input := range inputData {
				inputSize, _ := strconv.Atoi(input[1])
				deliveryID := strings.TrimSpace(input[0])
				var finalAmount int
				var resultPartner string
				var output *datamodel.Output
				if index < len(inputData) {
					var isDeliveryPossible bool
					for _, v := range result {
						if summaryKey == v.TheatreID && inputSize > v.MinSize && inputSize < v.MaxSize && !v.IsAssigned {
							finalAmount = inputSize * v.CostPerGB
							if finalAmount < v.MinCost {
								finalAmount = v.MinCost
							}
							v.IsAssigned = true
							resultPartner = v.PartnerID
							isDeliveryPossible = true
							break
						}
					}
					index++
					output = &datamodel.Output{
						DeliveryID:         deliveryID,
						IsDeliveryPossible: isDeliveryPossible,
						PartnerID:          resultPartner,
						ActualCost:         finalAmount,
					}
				}
				if output != nil {
					output2 = append(output2, output)
				}
			}
		}
	}
	utils.CreateOutputCSV(output2, constant.Output2)
	fmt.Println("Succesfully addressed problem 2")
}
