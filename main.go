// csv2json project main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCsv2Dict(csvFile string) []map[string]string {
	csvObj, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(csvObj)
	csvReader.FieldsPerRecord = -1
	var csvHeaders []string
	var csvDict = []map[string]string{}
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("%s", err)
		} else if csvHeaders == nil {
			csvHeaders = line
		} else {
			fieldDict := map[string]string{}
			for i, v := range line {
				fieldDict[csvHeaders[i]] = v
			}
			csvDict = append(csvDict, fieldDict)
		}

	}
	return csvDict
}

func main() {
	csvPtr := flag.String("csv", "", "csv file to be covert to json")
	flag.Parse()
	if *csvPtr == "" {
		log.Fatal("please give csv file")
	}

	jsonData := ReadCsv2Dict(*csvPtr)

	json, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(string(json))
}
