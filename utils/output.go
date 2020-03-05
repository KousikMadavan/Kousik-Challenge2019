package utils

import (
	"fmt"
	"os"

	"../datamodel"
)

// CreateOutputCSV ...
func CreateOutputCSV(output []*datamodel.Output, fileName string) {
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(fileName)
		if err != nil {
			return
		}
		defer file.Close()
	}
	openFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer openFile.Close()
	for _, out := range output {
		_, err = openFile.WriteString(fmt.Sprintf("%s,%t,%s,%d\n", out.DeliveryID, out.IsDeliveryPossible, out.PartnerID, out.ActualCost))
	}
}
