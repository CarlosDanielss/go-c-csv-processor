package main

import (
	"C"
	"fmt"
	"log"
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
