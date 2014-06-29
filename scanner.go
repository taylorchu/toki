package toki

import (
	"fmt"
	"regexp"
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

func (this *Position) moveByString(s string) {
	this.Line += strings.Count(s, "\n")
	last := strings.LastIndex(s, "\n")
	if last != -1 {
		this.Column = 1
	}
	this.Column += utf8.RuneCountInString(s[last+1:])
}

type TokenDef struct {
	Type   TokenType
	Pattern string
	regexpCompiled  *regexp.Regexp
}

func (this *TokenDef) compile() {
	this.regexpCompiled = regexp.MustCompile("^" + this.Pattern)
}

type Scanner struct {
	Space string
	pos   Position
	input string
	def   []TokenDef
}

type Token struct {
	Type  TokenType
	Value string
	Pos   Position
}

func (this *Token) String() string {
	return fmt.Sprintf("Line: %v, Column: %v, %v", this.Pos.Line, this.Pos.Column, this.Value)
}

func New(defs []TokenDef, s string) *Scanner {
	this := new(Scanner)
	this.input = s
	this.pos.Line = 1
	this.pos.Column = 1
	this.Space = `\s`
	for i := range defs {
		defs[i].compile()
	}
	this.def = defs
	return this
}

func (this *Scanner) skip() {
	result := regexp.MustCompile("^[" + this.Space + "]+").FindString(this.input)
	if result == "" {
		return
	}
	this.input = strings.TrimPrefix(this.input, result)
	this.pos.moveByString(result)
}

func (this *Scanner) find() *Token {
	this.skip()
	if len(this.input) == 0 {
		return &Token{Type: TokenEOF, Pos: this.pos}
	}
	for _, r := range this.def {
		result := r.regexpCompiled.FindString(this.input)
		if result == "" {
			continue
		}
		return &Token{Type: r.Type, Value: result, Pos: this.pos}
	}
	return &Token{Type: TokenError, Pos: this.pos, Value: this.input}
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
	this.pos.moveByString(t.Value)
	return t
}
