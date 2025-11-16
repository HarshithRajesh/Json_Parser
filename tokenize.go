package main

import (
	"fmt"
	"os"
	// "unicode"
)

type Token struct {
	Kind  string
	Value string
}

//
// func main() {
// 	data, _ := os.ReadFile(os.Args[1])
// 	src := string(data)
//
// 	tokens := tokenize(src)
// 	for _, t := range tokens {
// 		fmt.Printf("TOKEN: %s %s\n", t.Kind, t.Value)
// 	}
// }

func tokenize(src string) []Token {
	var tokens []Token
	i := 0

	for i < len(src) {
		ch := src[i]

		if ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r' {
			i++
			continue
		}
		if ch == '{' {
			tokens = append(tokens, Token{Kind: "LBRACE"})
			i++
			continue
		}
		if ch == '}' {
			tokens = append(tokens, Token{Kind: "RBRACE"})
			i++
			continue
		}
		if ch == ':' {

			tokens = append(tokens, Token{Kind: "COLON"})
			i++
			continue
		}
		if ch == '"' {

			start := i + 1
			i++
			for i < len(src) && src[i] != '"' {
				i++
			}
			if i >= len(src) {
				fmt.Println("Error : Unclosed string")
				os.Exit(1)
			}
			value := src[start:i]
			tokens = append(tokens, Token{Kind: "STRING", Value: value})
			i++
			continue
		}
		fmt.Println("ERROR: bad char '&c'\n", ch)
		os.Exit(1)
	}
	tokens = append(tokens, Token{Kind: "EOF"})
	return tokens
}
