package pscript

import (
	"unicode"
)

type TokenType int

const (
	TokenLet TokenType = iota
	TokenIdent
	TokenAssign
	TokenNumber
	TokenString
	TokenComma
	TokenLParen
	TokenRParen
	TokenEOF
	TokenFn
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token

	switch l.ch {
	case '=':
		tok = Token{Type: TokenAssign, Literal: "="}
	case ',':
		tok = Token{Type: TokenComma, Literal: ","}
	case '(':
		tok = Token{Type: TokenLParen, Literal: "("}
	case ')':
		tok = Token{Type: TokenRParen, Literal: ")"}
	case 0:
		tok = Token{Type: TokenEOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = TokenNumber
			tok.Literal = l.readNumber()
			return tok
		} else if l.ch == '"' {
			tok.Type = TokenString
			tok.Literal = l.readString()
			return tok
		} else {
			tok = Token{Type: TokenEOF, Literal: ""} // Illegal character
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	l.readChar() // skip "
	position := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	s := l.input[position:l.position]
	l.readChar() // skip "
	return s
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"let": TokenLet,
		"fn":  TokenFn,
	}
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TokenIdent
}
