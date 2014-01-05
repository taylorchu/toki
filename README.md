toki
====

Regexp tokenizer/scanner in Go.

The scanner takes in a list of token definitions and string input. 
it is general and easy to use.

token definition
================
a token is a pair of constant and regexp string.

```

const (
	NUMBER toki.TokenType = iota
	PLUS
	STRING
)

input := "1  + 2+3 + happy birthday  "
fmt.Println("input:", input)
s := new(toki.Scanner).Init(
	[]toki.TokenDef{
		{NUMBER, "[0-9]+"},
		{PLUS, "\\+"},
		{STRING, "[a-z]+"},
	}, input)
for {
	t := s.Next() // or s.Peek()
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

```
