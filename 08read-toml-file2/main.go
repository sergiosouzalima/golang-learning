package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	SourceFile         string   `toml:"source_file"`
	TargetFile         string   `toml:"target_file"`
	SourceFileSkipRows int      `toml:"source_file_skip_rows"`
	EmailValidation    bool     `toml:"email_validation"`
	PositionFilterFor  []string `toml:"position_filter_for"`
}

type AppConfig struct {
	Config Config `toml:"config"`
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	var appConfig AppConfig
	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&appConfig)
	if err != nil {
		log.Fatalf("Error decoding file: %s", err)
	}

	fmt.Printf("Source File: %s\n", appConfig.Config.SourceFile)
	fmt.Printf("Target File: %s\n", appConfig.Config.TargetFile)
	fmt.Printf("Source File Skip Rows: %d\n", appConfig.Config.SourceFileSkipRows)
	fmt.Printf("Email Validation: %t\n", appConfig.Config.EmailValidation)
	fmt.Printf("Position Filter For: %v\n", appConfig.Config.PositionFilterFor)

	fmt.Println("Position Filter For:")
	for index, position := range appConfig.Config.PositionFilterFor {
		fmt.Printf("%d: %s\n", index+1, position)
	}
}
