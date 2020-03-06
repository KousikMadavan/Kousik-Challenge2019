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
	var index, sumCostPerGB int
	_ = sumCostPerGB
	for summaryKey, sum := range inputSummary {
		for _, val := range result {
			// sum is the total amount for each theatre
			if sum >= val.MinSize && sum <= val.MaxSize {
				canAccomodate = true
				sumCostPerGB = val.CostPerGB
				break
			}
			continue
		}
		var finalAmount int
		var resultPartner string
		var output *datamodel.Output
		var isDeliveryPossible bool
		for _, input := range inputData {
			inputSize, _ := strconv.Atoi(input[1])
			deliveryID := strings.TrimSpace(input[0])
			if index < len(inputData) {
				if !canAccomodate {
					finalAmount, resultPartner, isDeliveryPossible = resultParse(summaryKey, result, finalAmount, inputSize, resultPartner, isDeliveryPossible)
					index++
					output = &datamodel.Output{
						DeliveryID:         deliveryID,
						IsDeliveryPossible: isDeliveryPossible,
						PartnerID:          resultPartner,
						ActualCost:         finalAmount,
					}
				} else {
					finalAmount, resultPartner, isDeliveryPossible = resultParse(input[2], result, finalAmount, inputSize, resultPartner, isDeliveryPossible)
					index++
					output = &datamodel.Output{
						DeliveryID:         deliveryID,
						IsDeliveryPossible: isDeliveryPossible,
						PartnerID:          resultPartner,
						ActualCost:         finalAmount,
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
}

func resultParse(summaryKey string, result []*datamodel.TheatrePartnerData, finalAmount int, inputSize int, resultPartner string, isDeliveryPossible bool) (int, string, bool) {
	for _, v := range result {
		if summaryKey == v.TheatreID && inputSize > v.MinSize && inputSize < v.MaxSize && (v.AssignedUnits == 0 || (v.AssignedUnits > 0 && (v.Capacity-v.AssignedUnits >= inputSize))) {
			finalAmount = inputSize * v.CostPerGB
			if finalAmount < v.MinCost {
				finalAmount = v.MinCost
			}
			resultPartner = v.PartnerID
			updateAssignedUnits(resultPartner, result, inputSize)
			isDeliveryPossible = true
			break
		}
	}
	return finalAmount, resultPartner, isDeliveryPossible
}

func updateAssignedUnits(resultPartner string, result []*datamodel.TheatrePartnerData, inputSize int) {
	for _, v := range result {
		if v.PartnerID == resultPartner {
			v.IsAssigned = true
			v.AssignedUnits = inputSize
		}
	}
}
