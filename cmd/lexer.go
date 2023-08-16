package main

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
