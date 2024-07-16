package main

import (
	"encoding/csv"
	"errors"
	"strings"
)

type Filter struct {
	Column   string
	Operator string
	Value    string
}

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

func parseFilters(rowFilterDefinitions string) ([]Filter, error) {
	filterStrings := strings.Split(rowFilterDefinitions, "\n")
	filters := make([]Filter, len(filterStrings))

	for i, filterString := range filterStrings {
		parts := strings.SplitN(filterString, ">", 2)

		if len(parts) == 2 {
			filters[i] = Filter{parts[0], ">", parts[1]}
			continue
		}

		parts = strings.SplitN(filterString, "<", 2)

		if len(parts) == 2 {
			filters[i] = Filter{parts[0], "<", parts[1]}
			continue
		}

		parts = strings.SplitN(filterString, "=", 2)

		if len(parts) == 2 {
			filters[i] = Filter{parts[0], "=", parts[1]}
		}
	}

	return filters, nil
}

func applyFilters(records [][]string, filters []Filter, headers []string) [][]string {
	var result [][]string
	result = append(result, headers)

	for _, row := range records[1:] {
		if matchFilters(row, filters, headers) {
			result = append(result, row)
		}
	}

	return result
}

func matchFilters(row []string, filters []Filter, headers []string) bool {
	for _, filter := range filters {
		colIndex := -1

		for i, header := range headers {
			if filter.Column == header {
				colIndex = i
				break
			}
		}

		if colIndex == -1 {
			return false
		}

		cellValue := row[colIndex]

		switch filter.Operator {
		case ">":
			if cellValue <= filter.Value {
				return false
			}
		case "<":
			if cellValue >= filter.Value {
				return false
			}
		case "=":
			if cellValue != filter.Value {
				return false
			}
		}
	}

	return true
}

func selectColumns(records [][]string, colIndexes []int) [][]string {
	var result [][]string

	for _, row := range records {
		selectedRow := make([]string, len(colIndexes))

		for i, index := range colIndexes {
			selectedRow[i] = row[index]
		}

		result = append(result, selectedRow)
	}

	return result
}
