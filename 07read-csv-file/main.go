package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Connection represents a row in the CSV file
type Connection struct {
	FirstName    string
	LastName     string
	EmailAddress string
	Company      string
	Position     string
	ConnectedOn  string
}

func main() {
	// Open the CSV file
	file, err := os.Open("Connections.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fmt.Println("passo 001")
	counter := 0

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// ignore the first 4 lines
		counter++
		if counter <= 4 {
			continue
		}

		s := scanner.Text()
		fields := strings.Split(s, ",")
		var arrFields [6]string    // declare a fixed size array
		copy(arrFields[:], fields) // copy the slice elements to the array

		// do something with a line
		fmt.Printf("line %d: %s\n", counter, arrFields[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Create a CSV reader from the file
	//reader := csv.NewReader(file)
	/*reader := csv.NewReader(bufio.NewReader(file))


		if err != nil {
			//log.Fatal(err)
		}

		fmt.Println("passo 001.030")

	}

	fmt.Println("passo 002")

	// Read the first line as header and ignore it
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Header:", header)

	// Read the rest of the lines as records and store them in a slice of Connection structs
	var connections []Connection
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error occurred
		}

		// Create a Connection struct from the record and append it to the slice
		connection := Connection{
			FirstName:    record[0],
			LastName:     record[1],
			EmailAddress: record[2],
			Company:      record[3],
			Position:     record[4],
			ConnectedOn:  record[5],
		}
		connections = append(connections, connection)
	}

	// Print the slice of connections for testing purposes
	fmt.Println("Connections:")
	for _, connection := range connections {
		fmt.Println(connection)
	}
	*/
}
