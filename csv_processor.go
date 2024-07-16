package main

import (
	"encoding/csv"
	"errors"
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

func getColumnIndexes(headers []string, selectedColumns string) ([]int, error) {
	selectedCols := strings.Split(selectedColumns, ",")
	colIndexes := make([]int, len(selectedCols))

	for i, col := range selectedCols {
		found := false

		for j, header := range headers {
			if col == header {
				colIndexes[i] = j
				found = true
				break
			}
		}

		if !found {
			return nil, errors.New("A coluna não foi encontrada " + col)
		}
	}

	return colIndexes, nil
}
