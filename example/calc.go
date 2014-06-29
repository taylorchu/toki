package main

import (
	"fmt"
	"github.com/taylorchu/toki"
)

const (
	NUMBER toki.TokenType = iota
	PLUS
	STRING
)

func main() {
	input := "1  + 2+3 + happy birthday  "
	fmt.Println("input:", input)
	s := toki.New([]toki.TokenDef{
			{Type: NUMBER, Pattern: "[0-9]+"},
			{Type: PLUS, Pattern: "\\+"},
			{Type: STRING, Pattern: "[a-z]+"},
		})
	s.SetInput(input)
	for {
		t := s.Next()
		if t.Type == toki.TokenEOF {
			fmt.Println("eof", t.Type, t)
			break
		}
		if t.Type == toki.TokenError {
			fmt.Println("error", t.Type, t)
			return
		}
		fmt.Println(t)
	}

}
