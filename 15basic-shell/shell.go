package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("basic > ")
		scanner.Scan()
		text := scanner.Text()

		tokens, err := Run("<stdin>", text)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(tokens)
		}
	}
}
