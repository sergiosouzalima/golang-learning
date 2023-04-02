package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	records, err := readCSV("clientes.csv")
	if err != nil {
		log.Fatal(err)
	}

	validRecords := filterValidRecords(records)

	err = writeCSV("clientes_validos.csv", validRecords)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Arquivo clientes_validos.csv criado com dados válidos.")
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func writeCSV(filename string, records [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

func filterValidRecords(records [][]string) [][]string {
	idCounts := make(map[string]int)
	validRecords := [][]string{records[0]} // include the header row

	for i, record := range records {
		if i == 0 {
			continue // skip header row
		}
		id := strings.TrimSpace(record[0]) // assuming ID is the first column
		if id != "" {
			idCounts[id]++
			if idCounts[id] == 1 {
				validRecords = append(validRecords, record) // include valid row
			}
		}
	}

	return validRecords
}
