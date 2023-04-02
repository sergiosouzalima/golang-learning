package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EncodingCharacters struct {
	FieldSeparator        string
	ComponentSeparator    string
	RepetitionSeparator   string
	EscapeCharacter       string
	SubcomponentSeparator string
}

type MSHSegment struct {
	EncodingChars        EncodingCharacters
	SendingApplication   string
	SendingFacility      string
	ReceivingApplication string
	ReceivingFacility    string
	MessageDateTime      string
	Security             string
	MessageType          string
	EventType            string
	MessageControlID     string
	ProcessingID         string
	VersionID            string
}

func main() {
	mshLine := readFirstLine("hl7.txt")
	mshSegment := parseMSHSegment(mshLine)
	printMSHSegmentBreakdown(mshSegment)
}

func readFirstLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func parseMSHSegment(msh string) MSHSegment {
	components := strings.Split(msh, "|")

	encodingChars := components[1]
	enc := EncodingCharacters{
		FieldSeparator:        "|",
		ComponentSeparator:    string(encodingChars[0]),
		RepetitionSeparator:   string(encodingChars[1]),
		EscapeCharacter:       string(encodingChars[2]),
		SubcomponentSeparator: string(encodingChars[3]),
	}

	msgTypeEvent := strings.Split(components[8], enc.ComponentSeparator)

	return MSHSegment{
		EncodingChars:        enc,
		SendingApplication:   components[2],
		SendingFacility:      components[3],
		ReceivingApplication: components[4],
		ReceivingFacility:    components[5],
		MessageDateTime:      components[6],
		Security:             components[7],
		MessageType:          msgTypeEvent[0],
		EventType:            msgTypeEvent[1],
		MessageControlID:     components[9],
		ProcessingID:         components[10],
		VersionID:            components[11],
	}
}

func printMSHSegmentBreakdown(msh MSHSegment) {
	fmt.Printf("MSH Segment Breakdown:\n")
	fmt.Printf("Encoding Characters:\n")
	fmt.Printf("Field Separator: %q, Component Separator: %q, Repetition Separator: %q, Escape Character: %q, Subcomponent Separator: %q\n",
		msh.EncodingChars.FieldSeparator, msh.EncodingChars.ComponentSeparator, msh.EncodingChars.RepetitionSeparator, msh.EncodingChars.EscapeCharacter, msh.EncodingChars.SubcomponentSeparator)
	fmt.Printf("Sending Application: %s, Sending Facility: %s\n", msh.SendingApplication, msh.SendingFacility)
	fmt.Printf("Receiving Application: %s, Receiving Facility: %s\n", msh.ReceivingApplication, msh.ReceivingFacility)
	fmt.Printf("Message Date/Time: %s, Security: %s\n", msh.MessageDateTime, msh.Security)
	fmt.Printf("Message Type: %s, Event Type: %s\n", msh.MessageType, msh.EventType)
	fmt.Printf("Message Control ID: %s, Processing ID: %s, Version ID: %s\n", msh.MessageControlID, msh.ProcessingID, msh.VersionID)
	fmt.Printf("\n")
}
