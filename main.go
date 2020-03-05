package main

import (
	"fmt"

	"./utils"

	"./constant"
	"./output"
	"./partner"
)

func main() {
	//Read Input from CSV Reader method
	inputData, _ := utils.CommonCSVReader(constant.Input)

	// Open the file
	partnerRecords := partner.ReadPartnersCSV()

	//Read Input from CSV Reader method
	capacityData, _ := utils.CommonCSVReader(constant.Capacities)

	//This has the capacities for each partner : Eg: P1: 350
	partnerCapacities := partner.GetPartnerCapacities(capacityData)

	//Fetch the input Summation of T1,T2 as 660,1025 respectively in keyvalue pair.
	inputSummary := partner.GetInputSummary(inputData)

	//Output1 logic and fetch back theatre partner mapping for output 2.
	theatrePartnerData := output.Output1(inputData, partnerRecords, partnerCapacities)

	//Sort functionality to order partner data with capacity starting from minCost.
	result := utils.SortPartner(theatrePartnerData)

	// Output2 logic and export the data to output2.csv
	output.Output2(inputSummary, inputData, result)

	//Complete
	fmt.Println("Complete")
}
