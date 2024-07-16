package main

import (
	"C"
	"fmt"
	"log"
	"os"
)

//export processCsv
func processCsv(csv *C.char, selectedColumns *C.char, rowFilterDefinitions *C.char) {
	csvData := C.GoString(csv)
	columns := C.GoString(selectedColumns)
	filters := C.GoString(rowFilterDefinitions)

	result, err := ProcessCsvData(csvData, columns, filters)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	fmt.Println(result)
}

//export processCsvFile
func processCsvFile(csvFilePath *C.char, selectedColumns *C.char, rowFilterDefinitions *C.char) {
	filePath := C.GoString(csvFilePath)
	csvData, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	selectedCols := C.GoString(selectedColumns)
	filterDefs := C.GoString(rowFilterDefinitions)

	result, err := ProcessCsvData(string(csvData), selectedCols, filterDefs)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	fmt.Println(result)
}

func main() {}
