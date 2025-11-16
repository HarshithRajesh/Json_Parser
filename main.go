package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to JSON Parser!!!")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: :%s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	content := string(data)

	if content == "{}" {
		fmt.Println("Valid JSON!")
		os.Exit(0)
	} else {
		fmt.Println("Invalid JSON!")
		os.Exit(0)
	}
}
