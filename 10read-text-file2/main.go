package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("Connections.csv")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read and ignore the first 3 lines
	for i := 0; i < 3; i++ {
		if !scanner.Scan() {
			fmt.Println("Error: Not enough lines in the file")
			return
		}
	}

	// Process the rest of the file as CSV
	reader := csv.NewReader(strings.NewReader(scanner.Text()))

	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading the header:", err)
		return
	}

	outputFile, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Error creating the output file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing the header:", err)
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading the record:", err)
			return
		}

		email := record[2]
		if validEmail(email) {
			err = writer.Write(record)
			if err != nil {
				fmt.Println("Error writing the record:", err)
				return
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		reader = csv.NewReader(strings.NewReader(line))

		record, err := reader.Read()
		if err != nil {
			fmt.Println("Error reading the record:", err)
			return
		}

		email := record[2]
		if validEmail(email) {
			err = writer.Write(record)
			if err != nil {
				fmt.Println("Error writing the record:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning the file:", err)
		return
	}
}

func validEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	valid := regexp.MustCompile(regex)
	return valid.MatchString(email)
}
