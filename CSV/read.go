package CSV

import (
	"encoding/csv"
	"os"
)

//reads Test csv file from Test given path
// the csv must contain only one column with the authors names
func readCsvFile(filePath string, separator rune) ([]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []string
	for _, line := range csvData {
		newLine := line[0]
		data = append(data, newLine)
	}

	return data, nil
}
