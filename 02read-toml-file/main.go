package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	// Load the TOML file
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// Get the postgres section as a map
	postgres := config.Get("postgres").(*toml.Tree).ToMap()

	// Print out the user and password
	fmt.Println("User:", postgres["user"])
	fmt.Println("Password:", postgres["password"])
}
