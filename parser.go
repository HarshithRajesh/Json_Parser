package main

import (
// "fmt"
// "os"
)

//	func main() {
//		data, _ := os.ReadFile(os.Args[1])
//		src := string(data)
//
//		tokens := tokenize(src)
//		valid := parse(tokens)
//		if valid {
//			os.Exit(0)
//		} else {
//			os.Exit(1)
//		}
//	}
func parse(tokens []Token) bool {
	i := 0

	if i >= len(tokens) || tokens[i].Kind != "LBRACE" {
		return false
	}
	i++

	if i < len(tokens) && tokens[i].Kind == "RBRACE" {
		i++
		return tokens[i].Kind == "EOF"
	}
	if i >= len(tokens) || tokens[i].Kind != "STRING" {
		return false
	}
	i++
	if i >= len(tokens) || tokens[i].Kind != "COLON" {
		return false
	}
	i++
	if i >= len(tokens) || tokens[i].Kind != "STRING" {
		return false
	}
	i++
	if i >= len(tokens) || tokens[i].Kind != "RBRACE" {
		return false
	}
	i++
	return i < len(tokens) && tokens[i].Kind == "EOF"

}
