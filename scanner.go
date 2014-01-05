package toki

import (
	//"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type TokenType uint32

const (
	TokenEOF TokenType = 1<<32 - 1 - iota
	TokenError
)

type Position struct {
	Line   int
	Column int
}

func (this *Position) MoveByString(s string) {
	this.Line += strings.Count(s, "\n")
	last := strings.LastIndex(s, "\n")
	if last != -1 {
		this.Column = 0
	}
	this.Column += utf8.RuneCountInString(s[last+1:])
}

type TokenDef struct {
	Type   TokenType
	Regexp string
}

func (this *TokenDef) Compile() TokenDefRuntime {
	return TokenDefRuntime{this.Type, regexp.MustCompile("^" + this.Regexp)}
}

type TokenDefRuntime struct {
	Type   TokenType
	Regexp *regexp.Regexp
}

type Scanner struct {
	Space string
	pos   Position
	input string
	def   []TokenDefRuntime
}

type Token struct {
	Type  TokenType
	Value string
	Pos   Position
}

func (this *Token) String() string {
	return "Line: " + strconv.Itoa(this.Pos.Line) + ", Column: " + strconv.Itoa(this.Pos.Column) + ", " + this.Value
}

func (this *Scanner) Init(defs []TokenDef, s string) *Scanner {
	this.input = s
	this.pos.Line = 1
	this.Space = "\t\r\n "
	for _, def := range defs {
		this.def = append(this.def, def.Compile())
	}
	return this
}

func (this *Scanner) skip() {
	result := regexp.MustCompile("^[" + this.Space + "]+").FindString(this.input)
	if result == "" {
		return
	}
	this.input = strings.TrimPrefix(this.input, result)
	this.pos.MoveByString(result)
}

func (this *Scanner) find() *Token {
	this.skip()
	if len(this.input) == 0 {
		return &Token{Type: TokenEOF, Pos: this.pos}
	}
	for _, r := range this.def {
		result := r.Regexp.FindString(this.input)
		if result == "" {
			continue
		}
		return &Token{Type: r.Type, Value: result, Pos: this.pos}
	}
	return &Token{Type: TokenError, Pos: this.pos}
}

func (this *Scanner) Peek() *Token {
	return this.find()
}

func (this *Scanner) Next() *Token {
	t := this.find()
	if t.Type == TokenError || t.Type == TokenEOF {
		return t
	}
	this.input = strings.TrimPrefix(this.input, t.Value)
	this.pos.MoveByString(t.Value)
	return t
}
