package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Title string
	Owner Owner
}

type Owner struct {
	Name string
	Age  int
}

func main() {
	// Create a config struct with some data
	config := Config{
		Title: "Example",
		Owner: Owner{
			Name: "Bob Lee",
			Age:  30,
		},
	}

	// Check if the file already exists
	_, err := os.Stat("config.toml")
	if err == nil {
		fmt.Println("File already exists")
		return
	}

	// Create a file to write to
	file, err := os.Create("config.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Encode the config to TOML and write to the file
	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data written to config.toml")
}
