package main

import (
	"fmt"
)

const (
	// Token/character we don't know about
	Illegal Type = "ILLEGAL"

	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "STRING"
	Number Type = "NUMBER"

	// The six structural tokens
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"

	// Values
	True  Type = "TRUE"
	False Type = "FALSE"
	Null  Type = "NULL"
)

func main() {
	fmt.Println("Welcome to JSON Parser!!!")

}
