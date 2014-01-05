package main

import (
	".."
	"fmt"
)

const (
	NUMBER toki.TokenType = iota
	PLUS
	STRING
)

func main() {
	input := "1  + 2+3 + happy birthday  "
	fmt.Println("input:", input)
	s := new(toki.Scanner).Init(
		[]toki.TokenDef{
			{NUMBER, "[0-9]+"},
			{PLUS, "\\+"},
			{STRING, "[a-z]+"},
		}, input)
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
		fmt.Println(t.Type, t)
	}

}
