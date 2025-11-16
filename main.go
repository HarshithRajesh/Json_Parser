package main

//
// import (
// 	"fmt"
// 	"os"
// )
// func main() {
// 	fmt.Println("Welcome to JSON Parser!!!")
//
// 	if len(os.Args) != 2 {
// 		fmt.Fprintf(os.Stderr, "Usage: :%s <filename>\n", os.Args[0])
// 		os.Exit(1)
// 	}
//
// 	filename := os.Args[1]
//
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
// 		os.Exit(1)
// 	}
//
// 	lexer := NewLexer(string(data))
//   parser := NewParser(lexer)
//   if parser.ParseObject() && lexer.AtEnd(){
//     os.Exit(0)
//   }
//   os.Exit(1)
// }
//
// //Writing Token Defination
// type TokenType int
//
// const (
//   TokenEOF    TokenType = iota
//   TokenLBrace
//   TokenRBrace
//   TokenString
//   TokenColon
// )
