package main

import (
	"fmt"
)
const (
  BRACE_OPEN    =   "BRACE_OPEN"
  BRACE_CLOSE   =   "BRACE_CLOSE"
  STRING        =   "STRING"
  NUMBER        =   "NUMBER"
  COLON         =   "COLON"
  COMMA         =   "COMMA"
  TRUE          =   "TRUE"
  FALSE         =   "FALSE"
  NULL          =   "NULL"
  BRACKET_OPEN  =   "BRACKET_OPEN"
  BRACKET_CLOSE =   "BRACKET_CLOSE"

)

type Token struc{
  Type    string
  Value   string
}

func Tokenize(jsonString string) ([]Token,error){
  current := 0
  stringLength := len(jsonString)
  fmt.Println(stringLength)
  tokens := []Token{}
  
  for current < stringLength{

  }


}


func main() {
	fmt.Println("Welcome to JSON Parser!!!")
  
}
