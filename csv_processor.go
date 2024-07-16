package main

import (
	"encoding/csv"
	"strings"
)

func readCsvData(csvData string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(csvData))
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}
