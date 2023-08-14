package main

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	IDENT
	INT
	ASSIGN
	PLUS
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
)

var tokenTypeNames = map[TokenType]string{
	ILLEGAL:   "ILLEGAL",
	EOF:       "EOF",
	IDENT:     "IDENT",
	INT:       "INT",
	ASSIGN:    "ASSIGN",
	PLUS:      "PLUS",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	LBRACE:    "LBRACE",
	RBRACE:    "RBRACE",
	FUNCTION:  "FUNCTION",
	LET:       "LET",
}

func TokStr(t TokenType) string {
	return tokenTypeNames[t]
}

type Token struct {
	Type    TokenType
	Literal string
}
