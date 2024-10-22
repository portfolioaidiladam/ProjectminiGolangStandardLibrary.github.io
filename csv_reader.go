package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	// csv string
	csvString := "aidil,adam,baik\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	// create reader
	reader := csv.NewReader(strings.NewReader(csvString))

	// read csv
	for {
		record, err := reader.Read()
		// if error is end of file then break
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
